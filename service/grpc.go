package service

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	grpc_dv "github.com/aditya37/geospatial-tracking/delivery/gRPC"
	"github.com/aditya37/geospatial-tracking/infra"
	"github.com/aditya37/geospatial-tracking/proto"

	mqtt_manager "github.com/aditya37/geospatial-tracking/repository/mqtt"

	device_manager "github.com/aditya37/geospatial-tracking/repository/mysql/device-manager"

	device_case "github.com/aditya37/geospatial-tracking/usecase/device"
	config "github.com/aditya37/get-env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc interface {
	Run()
}

type grpcSvc struct {
	grpcTrackingDlv *grpc_dv.Trackingdeliv
	close           func()
}

func NewGrpc() (Grpc, error) {
	// mqtt infra
	if err := infra.NewMqttClientInstance(
		infra.MQTTConf{
			Host:     config.GetString("MQTT_HOST", "127.0.0.1"),
			Port:     int64(config.GetInt("MQTT_PORT", 1883)),
			Username: config.GetString("MQTT_USERNAME", ""),
			Password: config.GetString("MQTT_PASSWORD", ""),
			ClientId: "tracking1",
		},
	); err != nil {
		return nil, err
	}
	mqttClientInfra := infra.GetMqttClientInstance()
	if mqttClientInfra == nil {
		if err := infra.NewMqttClientInstance(
			infra.MQTTConf{
				Host:     config.GetString("MQTT_HOST", "127.0.0.1"),
				Port:     int64(config.GetInt("MQTT_PORT", 1883)),
				Username: config.GetString("MQTT_USERNAME", ""),
				Password: config.GetString("MQTT_PASSWORD", ""),
				ClientId: "tracking1",
			},
		); err != nil {
			return nil, err
		}
		mqttClientInfra = infra.GetMqttClientInstance()
	}

	// mysql infra
	if err := infra.NewMysqlClient(
		infra.MysqlConfigParam{
			Host:     config.GetString("DB_HOST", "127.0.0.1"),
			Port:     config.GetInt("DB_PORT", 3306),
			Name:     config.GetString("DB_NAME", "db_geofencing"),
			User:     config.GetString("DB_USER", "root"),
			Password: config.GetString("DB_PASSWORD", "admin"),
		},
	); err != nil {
		return nil, err
	}
	mysqlInfra := infra.GetMysqlClientInstance()
	if mysqlInfra == nil {
		if err := infra.NewMysqlClient(
			infra.MysqlConfigParam{
				Host:     config.GetString("DB_HOST", "127.0.0.1"),
				Port:     config.GetInt("DB_PORT", 3306),
				Name:     config.GetString("DB_NAME", "db_geofencing"),
				User:     config.GetString("DB_USER", "root"),
				Password: config.GetString("DB_PASSWORD", "admin"),
			},
		); err != nil {
			return nil, err
		}
		mysqlInfra = infra.GetMysqlClientInstance()
	}

	// mqtt repo
	mqttManager, err := mqtt_manager.NewMqttManager(mqttClientInfra)
	if err != nil {
		return nil, err
	}
	deviceManagerRepo, err := device_manager.NewDeviceManager(mysqlInfra)
	if err != nil {
		return nil, err
	}

	deviceUsecase := device_case.NewDeviceUsecase(
		mqttManager,
		deviceManagerRepo,
	)

	// async
	// mqtt qos = https://www.emqx.com/id/blog/introduction-to-mqtt-qos
	go mqttManager.Subscribe(
		config.GetString("MQTT_TOPIC_REGISTER_DEVICE", "/device/req/register"),
		byte(
			config.GetInt("MQTT_TOPIC_REGISTER_DEVICE_QOS", 1),
		),
		deviceUsecase.SubscribeRegisterDevice,
	)
	go mqttManager.Subscribe(
		config.GetString("MQTT_TOPIC_GPS_TRACKING", "/device/req/tracking"),
		// TODO: Set to env
		byte(2),
		deviceUsecase.SubscribeGPSTracking,
	)
	grpcTrackingDeliv := grpc_dv.NewTrackingDelivery(deviceUsecase)

	return &grpcSvc{
		grpcTrackingDlv: grpcTrackingDeliv,
		close: func() {
			log.Println("Take rest broh!!!, all connection has been closed")
			deviceManagerRepo.Close()
		},
	}, nil
}

func (g *grpcSvc) Run() {
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
		defer g.close()
	}()

	server := grpc.NewServer()
	proto.RegisterGeotrackingServer(server, g.grpcTrackingDlv)
	reflection.Register(server)
	go func() {
		errs <- serve(
			server,
			healthcheck(),
		)
	}()
	log.Fatalf("Stop server with error detail: %v", <-errs)
}

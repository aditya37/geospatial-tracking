package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	grpc_dv "github.com/aditya37/geospatial-tracking/delivery/gRPC"
	grpc_deliv_mid "github.com/aditya37/geospatial-tracking/delivery/middleware"
	"github.com/aditya37/geospatial-tracking/infra"
	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/geospatial-tracking/repository"
	mqtt_manager "github.com/aditya37/geospatial-tracking/repository/mqtt"
	getenv "github.com/aditya37/get-env"

	chan_repo "github.com/aditya37/geospatial-tracking/repository/channel"
	firebase_manager "github.com/aditya37/geospatial-tracking/repository/firebase"
	gcp_manager "github.com/aditya37/geospatial-tracking/repository/gcppubsub"
	device_manager "github.com/aditya37/geospatial-tracking/repository/mysql/device-manager"
	cache_manager "github.com/aditya37/geospatial-tracking/repository/redis"
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

	close func()
}

func NewGrpc() (Grpc, error) {
	ctx := context.Background()
	// mqtt infra
	if err := infra.NewMqttClientInstance(
		infra.GetMQTTConfig(),
	); err != nil {
		return nil, err
	}
	mqttClientInfra := infra.GetMqttClientInstance()
	if mqttClientInfra == nil {
		if err := infra.NewMqttClientInstance(
			infra.GetMQTTConfig(),
		); err != nil {
			return nil, err
		}
		mqttClientInfra = infra.GetMqttClientInstance()
	}
	// mysql infra
	if err := infra.NewMysqlClient(
		infra.GetMysqlConfig(),
	); err != nil {
		return nil, err
	}
	mysqlInfra := infra.GetMysqlClientInstance()
	if mysqlInfra == nil {
		if err := infra.NewMysqlClient(
			infra.GetMysqlConfig(),
		); err != nil {
			return nil, err
		}
		mysqlInfra = infra.GetMysqlClientInstance()
	}

	// redis instance
	infra.NewRedisInstance(infra.GetRedisConfig())
	redisInfra := infra.GetRedisInstance()
	if redisInfra == nil {
		infra.NewRedisInstance(infra.GetRedisConfig())
		redisInfra = infra.GetRedisInstance()
	}
	// gcppubsubInstance...
	infra.NewGcpPubsubInstance(
		ctx,
		getenv.GetString("GCP_PROJECT_ID", ""),
	)
	gcpPubsubInstane := infra.GetGcpPubsubInstance()
	if gcpPubsubInstane == nil {
		infra.NewGcpPubsubInstance(
			ctx,
			getenv.GetString("GCP_PROJECT_ID", ""),
		)
		gcpPubsubInstane = infra.GetGcpPubsubInstance()
	}

	// firebase instance...
	if err := infra.NewFirebaseClient(
		ctx,
		infra.GetFirebaseConfig(),
	); err != nil {
		return nil, err
	}
	firebaseInstance := infra.GetFirebaseInstance()
	if firebaseInstance == nil {
		if err := infra.NewFirebaseClient(
			ctx,
			infra.GetFirebaseConfig(),
		); err != nil {
			return nil, err
		}
		firebaseInstance = infra.GetFirebaseInstance()
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

	// cache repo
	cacheManagerRepo := cache_manager.CacheManager(redisInfra)

	// streamer data
	gpsChannelStream := repository.NewChannelStreamGPS()
	gpsChanForward := repository.NewTrackingForward()
	go gpsChannelStream.Run()

	// channel for monitoring device by id
	chanDeviceMonitoringById := chan_repo.NewMonitoringDeviceById()
	go chanDeviceMonitoringById.Run()

	// gcppubsubManager...
	gcpPubsubManager := gcp_manager.NewGcpPubsubManager(gcpPubsubInstane)

	// firebase..
	fbsStorage, err := firebase_manager.NewStorageBucket(
		ctx,
		firebaseInstance,
	)
	if err != nil {
		return nil, err
	}
	deviceUsecase := device_case.NewDeviceUsecase(
		mqttManager,
		deviceManagerRepo,
		gpsChannelStream,
		cacheManagerRepo,
		gpsChanForward,
		gcpPubsubManager,
		chanDeviceMonitoringById,
		fbsStorage,
	)
	gpsChanForward.Subscribe(deviceUsecase.ForwardGPSTracking)
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
	// subscribe mqtt topic device logs..
	go mqttManager.Subscribe(
		"/device/logs",
		byte(0),
		deviceUsecase.SubscribeDeviceLog,
	)
	// SubscribeDeviceDetectdInGeofence
	go mqttManager.Subscribe(
		config.GetString("MQTT_TOPIC_SUBSCRIBE_DEVICE_DETECT", "/device/geofence/detect"),
		byte(1),
		deviceUsecase.SubscribeDeviceDetect,
	)
	grpcTrackingDeliv := grpc_dv.NewTrackingDelivery(
		deviceUsecase,
		gpsChannelStream,
		chanDeviceMonitoringById,
	)

	return &grpcSvc{
		grpcTrackingDlv: grpcTrackingDeliv,
		close: func() {
			log.Println("Take rest broh!!!, all connection has been closed")
			deviceManagerRepo.Close()
			cacheManagerRepo.Close()
			gcpPubsubManager.Close()
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
	opt := GetGrpcServerElasticApmOptions(
		SetUnaryMiddleware(grpc_deliv_mid.UnaryAuthMiddleware()),
		SetStreamMiddleware(grpc_deliv_mid.StreamAuthMiddleware()),
	)
	server := grpc.NewServer(opt...)
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

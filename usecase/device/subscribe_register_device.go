package device

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	logger "github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
	"github.com/aditya37/geospatial-tracking/usecase"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (du *DeviceUsecase) SubscribeRegisterDevice(c mqtt.Client, m mqtt.Message) {
	ctx := context.Background()
	payload, err := du.unmarshallRegisterPayload(m.Payload())
	if err != nil {
		logger.Logger().Error(err)
		return
	}
	// check device is registered or not
	if _, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, payload.Deviceid); err != nil {
		logger.Logger().Error(err)
		if err == repository.ErrDeviceNotFound {
			if err := du.registerDevice(ctx, payload); err != nil {
				logger.Logger().Error(err)
				// publish resp to device
				go du.publishRespRegister(
					"/device/resp/register",
					usecase.MqttRespRegisterDevice{
						Deviceid: payload.Deviceid,
						Message:  err.Error(),
						Status:   usecase.StatusFailedRegister.ToString(),
					},
				)
				m.Ack()
				return
			}

			go du.publishRespRegister(
				"/device/resp/register",
				usecase.MqttRespRegisterDevice{
					Deviceid: payload.Deviceid,
					Message:  "Register Device Success",
					Status:   usecase.StatusSuccessRegister.ToString(),
				},
			)
			m.Ack()
			return
		}
	}
	go du.publishRespRegister(
		"/device/resp/register",
		usecase.MqttRespRegisterDevice{
			Deviceid: payload.Deviceid,
			Message:  "Device id is valid",
			Status:   usecase.StatusValidDeviceId.ToString(),
		},
	)
	m.Ack()
}

// registerDevice....
func (du *DeviceUsecase) registerDevice(ctx context.Context, payload usecase.MqttRegisterDevicePayload) error {
	// do insert or register device
	if err := du.deviceManagerRepo.InsertDevice(
		ctx,
		entity.Device{
			DeviceId:   payload.Deviceid,
			DeviceType: payload.DeviceType,
			MacAddress: payload.MacAddress,
			ChipId:     payload.ChipId,
			I2cAddress: payload.I2cAddress,
		},
	); err != nil {
		return err
	}
	return nil
}

// unmarshallRegisterPayload....
func (du *DeviceUsecase) unmarshallRegisterPayload(data []byte) (usecase.MqttRegisterDevicePayload, error) {
	var payload usecase.MqttRegisterDevicePayload

	if err := json.Unmarshal(data, &payload); err != nil {
		return usecase.MqttRegisterDevicePayload{}, err
	}

	return payload, nil
}

// PublishNotify....
func (du *DeviceUsecase) publishRespRegister(topic string, data usecase.MqttRespRegisterDevice) error {
	logger.Logger().Info(fmt.Sprintf("Publish register response to %s", topic))
	j, _ := json.Marshal(data)
	log.Println("resp ", string(j))
	if err := du.mqttmanager.Publish(topic, 1, false, j); err != nil {
		return err
	}
	return nil
}

package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"

	logger "github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
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
			// reigster device if device id not found
			if err := du.registerDevice(ctx, payload); err != nil {
				logger.Logger().Error(err)
				// publish resp error...
				go du.publishRespRegister(
					"/device/resp/register",
					usecase.MqttRespRegisterDevice{
						Deviceid:      payload.Deviceid,
						Message:       err.Error(),
						Status:        usecase.StatusFailedRegister.ToString(),
						ValidSensorId: []int{},
					},
				)
			}
		} else {
			// publish another response error...
			go du.publishRespRegister(
				"/device/resp/register",
				usecase.MqttRespRegisterDevice{
					Deviceid:      payload.Deviceid,
					Message:       err.Error(),
					Status:        usecase.StatusFailedRegister.ToString(),
					ValidSensorId: []int{},
				},
			)
		}
	} else {
		// publish response if device is valid or has been registered....
		go du.publishRespRegister(
			"/device/resp/register",
			usecase.MqttRespRegisterDevice{
				Deviceid: payload.Deviceid,
				Message:  "Device id is valid",
				Status:   usecase.StatusValidDeviceId.ToString(),
			},
		)
	}
	m.Ack()
}

// registerDevice....
func (du *DeviceUsecase) registerDevice(ctx context.Context, payload usecase.MqttRegisterDevicePayload) error {
	if _, ok := proto.DeviceType_name[int32(payload.DeviceType)]; !ok {
		return errors.New("unknow device type")
	}

	if payload.NetworkMode == usecase.NETWORK_MODE_WLAN {
		// check attach sensor infornation or not
		if len(payload.EmbeddedSensor) <= 0 {
			// do insert or register device
			if _, err := du.deviceManagerRepo.InsertDevice(
				ctx,
				entity.Device{
					DeviceId: payload.Deviceid,
					DeviceType: int(
						proto.DeviceType(payload.DeviceType),
					),
					MacAddress:  payload.MacAddress,
					ChipId:      payload.ChipId,
					I2cAddress:  payload.I2cAddress,
					NetworkMode: payload.NetworkMode,
				},
			); err != nil {
				return err
			}

			// publish success response if success register device...
			du.publishRespRegister(
				"/device/resp/register",
				usecase.MqttRespRegisterDevice{
					Deviceid:              payload.Deviceid,
					Message:               "Register Device Success",
					Status:                usecase.StatusSuccessRegister.ToString(),
					CountSensorIdNotValid: 0,
					ValidSensorId:         []int{},
				},
			)
			return nil
		} else {
			// register device with embedded sensor
			if err := du.registerDeviceWithEmbeddedSensor(ctx, payload); err != nil {
				return err
			}
			return nil
		}

	} else if payload.NetworkMode == usecase.NETWORK_MODE_MOBILE_DATA {
		// process insert device with network mode mobile data
		if len(payload.EmbeddedSensor) <= 0 {
			// do insert or register device
			if _, err := du.deviceManagerRepo.InsertDevice(
				ctx,
				entity.Device{
					DeviceId: payload.Deviceid,
					DeviceType: int(
						proto.DeviceType(payload.DeviceType),
					),
					MacAddress:  payload.MacAddress,
					ChipId:      payload.ChipId,
					I2cAddress:  payload.I2cAddress,
					NetworkMode: payload.NetworkMode,
					SIM: entity.SIM{
						PhoneNo:     payload.PhoneNo,
						IMEI:        payload.IMEI,
						IMSI:        payload.IMSI,
						SIMOperator: payload.SimOperator,
						APN:         payload.APN,
					},
				},
			); err != nil {
				return err
			}
			// publish success response if success register device...
			du.publishRespRegister(
				"/device/resp/register",
				usecase.MqttRespRegisterDevice{
					Deviceid:              payload.Deviceid,
					Message:               "Register Device Success",
					Status:                usecase.StatusSuccessRegister.ToString(),
					CountSensorIdNotValid: 0,
					ValidSensorId:         []int{},
				},
			)
			return nil
		} else {
			// register device with embedded sensor
			if err := du.registerDeviceWithEmbeddedSensor(ctx, payload); err != nil {
				return err
			}
			return nil
		}
	} else {
		return errors.New("unknown network mode")
	}
}

// registerDeviceWithEmbeddedSensor....
func (du *DeviceUsecase) registerDeviceWithEmbeddedSensor(ctx context.Context, payload usecase.MqttRegisterDevicePayload) error {
	// attach sensor informataion and insert to db...
	validSensor, err := du.validateEmbeddedSensor(ctx, payload)
	if err != nil {
		logger.Logger().Error(err)
		return err
	}

	// get count embedded sensor id not valid or not available in database
	countNotValidSensor := math.Abs(
		float64(len(payload.EmbeddedSensor)) - float64(len(validSensor)),
	)

	// validate if all requested sensor not valid
	// will return error
	if len(validSensor) == 0 {
		return errors.New("all sensor id not found")
	}

	// if network mode mobile data
	// add information SIM
	var simOperator entity.SIM
	if payload.NetworkMode != usecase.NETWORK_MODE_WLAN {
		simOperator = entity.SIM{
			PhoneNo:     payload.PhoneNo,
			IMEI:        payload.IMEI,
			IMSI:        payload.IMSI,
			SIMOperator: payload.SimOperator,
			APN:         payload.APN,
		}
	}
	// do insert or register device
	lastInsertId, err := du.deviceManagerRepo.InsertDevice(
		ctx,
		entity.Device{
			DeviceId: payload.Deviceid,
			DeviceType: int(
				proto.DeviceType(payload.DeviceType),
			),
			MacAddress:  payload.MacAddress,
			ChipId:      payload.ChipId,
			I2cAddress:  payload.I2cAddress,
			NetworkMode: payload.NetworkMode,
			SIM:         simOperator,
		},
	)
	if err != nil {
		return err
	}

	var devicesensor []entity.DeviceSensor
	for _, v := range validSensor {
		devicesensor = append(devicesensor, entity.DeviceSensor{
			DeviceId: lastInsertId,
			SensorId: int64(v),
		})
	}

	if err := du.deviceManagerRepo.InsertEmbeddedSensorInDevice(ctx, devicesensor); err != nil {
		return err
	}

	// publish success response if success register device...
	du.publishRespRegister(
		"/device/resp/register",
		usecase.MqttRespRegisterDevice{
			Deviceid:              payload.Deviceid,
			Message:               "Register Device Success",
			Status:                usecase.StatusSuccessRegister.ToString(),
			CountSensorIdNotValid: int(countNotValidSensor),
			ValidSensorId:         validSensor,
		},
	)
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
	if err := du.mqttmanager.Publish(topic, 1, false, j); err != nil {
		return err
	}
	return nil
}

// validateEmbeddedSensor....
func (du *DeviceUsecase) validateEmbeddedSensor(ctx context.Context, payload usecase.MqttRegisterDevicePayload) ([]int, error) {
	var validSensorId []int

	// validate payload sensor in db
	sensors, err := du.deviceManagerRepo.GetSensorById(
		ctx,
		payload.EmbeddedSensor,
	)
	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}
	mapEmbeddedSensor := map[int]int{}
	for _, val := range payload.EmbeddedSensor {
		mapEmbeddedSensor[val] = val
	}
	for _, s := range sensors {
		v, ok := mapEmbeddedSensor[int(s.Id)]
		if ok {
			validSensorId = append(validSensorId, v)
		} else {
			log.Println("some key not found")
		}
	}
	return validSensorId, nil
}

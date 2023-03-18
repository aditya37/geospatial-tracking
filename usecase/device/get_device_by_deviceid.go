package device

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/geospatial-tracking/repository"
	getenv "github.com/aditya37/get-env"
	"github.com/aditya37/logger"
)

var (
	keyPrfxCacheDevice = "device.detail.%s"
)

// method for stream or read data from database or cache
func (du *DeviceUsecase) GetDeviceDetailByDeviceId(ctx context.Context, deviceId string) (proto.ResponseGetDeviceByDeviceId, error) {
	resp, err := du.loadDeviceDetailFromRedis(ctx, deviceId)
	if err != nil {
		device, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, deviceId)
		if err != nil {
			logger.Error(err)
			return proto.ResponseGetDeviceByDeviceId{}, err
		}

		// get attached sensor
		attachedSensor, err := du.deviceManagerRepo.GetAttachedSensorByDeviceId(ctx, deviceId)
		if err != nil {
			logger.Error(err)
			return proto.ResponseGetDeviceByDeviceId{}, err
		}

		// get qr code url
		deviceQrCode, err := du.deviceManagerRepo.GetDeviceQrCode(
			ctx,
			entity.QRDevice{
				DeviceId:  deviceId,
				EventType: 1,
			},
		)
		if err != nil {
			logger.Error(err)
			// return response success if device not have qr code...
			if err == repository.ErrDeviceNotFound {
				// set to redis
				result := proto.ResponseGetDeviceByDeviceId{
					Id:       device.Id,
					DeviceId: device.DeviceId,
					Device: &proto.Device{
						MacAddress:  device.MacAddress,
						DeviceType:  proto.DeviceType(device.DeviceType),
						ChipId:      device.ChipId,
						NetworkMode: device.NetworkMode,
						CreatedAt:   device.CreatedAt.Format(time.RFC3339),
					},
					NetworkDetail: &proto.Network{
						OperatorName: device.SIMOperator.Name,
						PhoneNo:      device.SIM.PhoneNo,
						Imei:         device.SIM.IMEI,
						Imsi:         device.SIM.IMSI,
						Apn:          device.SIM.APN,
						Status:       device.SIM.Status,
					},
					CreatedAt:           device.CreatedAt.Format(time.RFC3339),
					ModifiedAt:          device.ModifiedAt.Format(time.RFC3339),
					CountAttachedSensor: int64(len(attachedSensor.Sensor)),
					SystemUptime:        0, // TODO: add field system uptime and record from mqtt
					DeviceQrCode:        "",
					Sensors:             attachedSensor.Sensor,
				}
				// set redis cache....
				du.setCacheDeviceDetail(ctx, deviceId, result)
				return result, nil
			}
			return proto.ResponseGetDeviceByDeviceId{}, err
		}

		// set to redis
		result := proto.ResponseGetDeviceByDeviceId{
			Id:       device.Id,
			DeviceId: device.DeviceId,
			Device: &proto.Device{
				MacAddress:  device.MacAddress,
				DeviceType:  proto.DeviceType(device.DeviceType),
				ChipId:      device.ChipId,
				NetworkMode: device.NetworkMode,
				CreatedAt:   device.CreatedAt.Format(time.RFC3339),
			},
			NetworkDetail: &proto.Network{
				OperatorName: device.SIMOperator.Name,
				PhoneNo:      device.SIM.PhoneNo,
				Imei:         device.SIM.IMEI,
				Imsi:         device.SIM.IMSI,
				Apn:          device.SIM.APN,
				Status:       device.SIM.Status,
			},
			CreatedAt:           device.CreatedAt.Format(time.RFC3339),
			ModifiedAt:          device.ModifiedAt.Format(time.RFC3339),
			CountAttachedSensor: int64(len(attachedSensor.Sensor)),
			SystemUptime:        0, // TODO: add field system uptime and record from mqtt
			DeviceQrCode:        du.getDevicQrCodeUrl(deviceQrCode.QrFile),
			Sensors:             attachedSensor.Sensor,
		}
		// set redis cache....
		du.setCacheDeviceDetail(ctx, deviceId, result)
		return result, nil
	}
	return resp, nil
}

// getDevicQrCodeUrl....
func (du *DeviceUsecase) getDevicQrCodeUrl(filename string) string {
	return fmt.Sprintf(
		"https://firebasestorage.googleapis.com/v0/b/%s.appspot.com/o/%s?alt=media&token=%s",
		getenv.GetString("FIREBASE_PROJECT_ID", "device-service-1029d"),
		filename,
		filename,
	)
}

// setCacheDeviceDetail
func (du *DeviceUsecase) setCacheDeviceDetail(_ context.Context, deviceId string, data proto.ResponseGetDeviceByDeviceId) error {
	key := fmt.Sprintf(keyPrfxCacheDevice, deviceId)
	byteDeviceDetail, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		return err
	}
	return du.cacheManager.Set(key, byteDeviceDetail, time.Duration(86400*time.Second))
}

// get data to redis cache
func (du *DeviceUsecase) loadDeviceDetailFromRedis(_ context.Context, deviceId string) (proto.ResponseGetDeviceByDeviceId, error) {
	result, err := du.cacheManager.Get(fmt.Sprintf(keyPrfxCacheDevice, deviceId))
	if err != nil {
		logger.Error(err)
		return proto.ResponseGetDeviceByDeviceId{}, err
	}
	var record proto.ResponseGetDeviceByDeviceId
	if err := json.Unmarshal([]byte(result), &record); err != nil {
		logger.Error(err)
		return proto.ResponseGetDeviceByDeviceId{}, err
	}
	return record, nil
}

package device

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/proto"
)

var (
	keyPrfxCacheDevice = "device.detail.%s"
)

// method for stream or read data from database or cache
func (du *DeviceUsecase) GetDeviceDetailByDeviceId(ctx context.Context, deviceId string) (proto.Device, error) {
	resp, err := du.loadDeviceDetailFromRedis(ctx, deviceId)
	if err != nil {
		device, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, deviceId)
		if err != nil {
			util.Logger().Error(err)
			return proto.Device{}, err
		}

		// set to redis
		result := proto.Device{
			MacAddress: device.MacAddress,
			DeviceType: proto.DeviceType(device.DeviceType),
			ChipId:     device.ChipId,
			CreatedAt:  device.CreatedAt.Format(time.RFC3339),
		}
		du.setCacheDeviceDetail(ctx, deviceId, result)
		return result, nil

	}
	return resp, nil
}

// setCacheDeviceDetail
func (du *DeviceUsecase) setCacheDeviceDetail(_ context.Context, deviceId string, data proto.Device) error {
	key := fmt.Sprintf(keyPrfxCacheDevice, deviceId)
	byteDeviceDetail, err := json.Marshal(data)
	if err != nil {
		util.Logger().Error(err)
		return err
	}
	return du.cacheManager.Set(key, byteDeviceDetail, time.Duration(86400*time.Second))
}

// get data to redis cache
func (du *DeviceUsecase) loadDeviceDetailFromRedis(_ context.Context, deviceId string) (proto.Device, error) {
	result, err := du.cacheManager.Get(fmt.Sprintf(keyPrfxCacheDevice, deviceId))
	if err != nil {
		util.Logger().Error(err)
		return proto.Device{}, err
	}
	var record proto.Device
	if err := json.Unmarshal([]byte(result), &record); err != nil {
		util.Logger().Error(err)
		return proto.Device{}, err
	}
	return record, nil
}

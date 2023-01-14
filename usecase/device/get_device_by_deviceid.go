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
func (du *DeviceUsecase) GetDeviceDetailByDeviceId(ctx context.Context, deviceId string) (proto.ResponseGetDeviceByDeviceId, error) {
	resp, err := du.loadDeviceDetailFromRedis(ctx, deviceId)
	if err != nil {
		device, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, deviceId)
		if err != nil {
			util.Logger().Error(err)
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
			CreatedAt:  device.CreatedAt.Format(time.RFC3339),
			ModifiedAt: device.ModifiedAt.Format(time.RFC3339),
		}
		du.setCacheDeviceDetail(ctx, deviceId, result)
		return result, nil

	}
	return resp, nil
}

// setCacheDeviceDetail
func (du *DeviceUsecase) setCacheDeviceDetail(_ context.Context, deviceId string, data proto.ResponseGetDeviceByDeviceId) error {
	key := fmt.Sprintf(keyPrfxCacheDevice, deviceId)
	byteDeviceDetail, err := json.Marshal(data)
	if err != nil {
		util.Logger().Error(err)
		return err
	}
	return du.cacheManager.Set(key, byteDeviceDetail, time.Duration(86400*time.Second))
}

// get data to redis cache
func (du *DeviceUsecase) loadDeviceDetailFromRedis(_ context.Context, deviceId string) (proto.ResponseGetDeviceByDeviceId, error) {
	result, err := du.cacheManager.Get(fmt.Sprintf(keyPrfxCacheDevice, deviceId))
	if err != nil {
		util.Logger().Error(err)
		return proto.ResponseGetDeviceByDeviceId{}, err
	}
	var record proto.ResponseGetDeviceByDeviceId
	if err := json.Unmarshal([]byte(result), &record); err != nil {
		util.Logger().Error(err)
		return proto.ResponseGetDeviceByDeviceId{}, err
	}
	return record, nil
}

package device

import (
	"context"
	"errors"

	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/logger"
)

func (du *DeviceUsecase) GetSensorById(ctx context.Context, in *proto.RequestGetSemsorById) (*proto.ResponseGetSensorById, error) {
	// get sensor detail...
	if !in.CheckEmbedded && in.DeviceId == "" {
		return du.processGetSensorById(ctx, in)
	} else if in.CheckEmbedded && in.DeviceId != "" {
		return du.processGetSensorEmbeddedStatus(ctx, in)
	} else {
		return du.processGetSensorById(ctx, in)
	}

}
func (du *DeviceUsecase) processGetSensorEmbeddedStatus(ctx context.Context, in *proto.RequestGetSemsorById) (*proto.ResponseGetSensorById, error) {
	sensor, err := du.deviceManagerRepo.GetSensorById(
		ctx,
		[]int{int(in.SensorId)},
	)
	if err != nil {
		logger.Error(err)
		return &proto.ResponseGetSensorById{}, err
	}

	// if len sensor = 0
	// sensor not found....
	if len(sensor) == 0 {
		logger.Error("sensor id not found")
		return &proto.ResponseGetSensorById{}, errors.New("sensor id not found")
	}

	// validate device id...
	device, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, in.DeviceId)
	if err != nil {
		logger.Error(err)
		return &proto.ResponseGetSensorById{}, err
	}
	if err := du.deviceManagerRepo.GetDeviceSensorStatusEmbedded(ctx, sensor[0].Id, device.Id); err != nil {
		logger.Error(err)
		if err == errors.New("sensor not embedded in device") {
			return &proto.ResponseGetSensorById{}, err
		}
		return &proto.ResponseGetSensorById{}, err
	}
	return &proto.ResponseGetSensorById{
		Id:          sensor[0].Id,
		SensorName:  sensor[0].SensorName,
		Description: sensor[0].Description,
		SensorType:  sensor[0].SensorType,
	}, nil
}

// process get sensor by id
func (du *DeviceUsecase) processGetSensorById(ctx context.Context, in *proto.RequestGetSemsorById) (*proto.ResponseGetSensorById, error) {
	sensor, err := du.deviceManagerRepo.GetSensorById(
		ctx,
		[]int{int(in.SensorId)},
	)
	if err != nil {
		logger.Error(err)
		return &proto.ResponseGetSensorById{}, err
	}

	// if len sensor = 0
	// sensor not found....
	if len(sensor) == 0 {
		logger.Error("sensor id not found")
		return &proto.ResponseGetSensorById{}, errors.New("sensor id not found")
	}
	return &proto.ResponseGetSensorById{
		Id:          sensor[0].Id,
		SensorName:  sensor[0].SensorName,
		Description: sensor[0].Description,
		SensorType:  sensor[0].SensorType,
	}, nil
}

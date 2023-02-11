package device

import (
	"context"
	"errors"

	logger "github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/proto"
)

func (du *DeviceUsecase) GetSensorById(ctx context.Context, in *proto.RequestGetSemsorById) (*proto.ResponseGetSensorById, error) {
	sensor, err := du.deviceManagerRepo.GetSensorById(
		ctx,
		[]int{int(in.SensorId)},
	)
	if err != nil {
		logger.Logger().Error(err)
		return &proto.ResponseGetSensorById{}, err
	}

	// if len sensor = 0
	// sensor not found....
	if len(sensor) == 0 {
		return &proto.ResponseGetSensorById{}, errors.New("sensor id not found")
	}
	return &proto.ResponseGetSensorById{
		Id:          sensor[0].Id,
		SensorName:  sensor[0].SensorName,
		Description: sensor[0].Description,
		SensorType:  sensor[0].SensorType,
	}, nil
}

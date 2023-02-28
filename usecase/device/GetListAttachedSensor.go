package device

import (
	"context"

	"github.com/aditya37/geospatial-tracking/proto"
)

func (s *DeviceUsecase) GetListAttachedSensor(ctx context.Context, in *proto.RequestGetAttachedSensor) (proto.ResponseGetAttachedSensor, error) {
	if _, err := s.deviceManagerRepo.GetDeviceByDeviceId(ctx, in.DeviceId); err != nil {
		return proto.ResponseGetAttachedSensor{}, err
	}
	resp, err := s.deviceManagerRepo.GetAttachedSensorByDeviceId(ctx, in.DeviceId)
	if err != nil {
		return proto.ResponseGetAttachedSensor{}, err
	}
	return proto.ResponseGetAttachedSensor{
		DeviceId: in.DeviceId,
		Sensor:   resp.Sensor,
	}, nil
}

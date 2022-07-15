package grpc

import (
	"context"

	"github.com/aditya37/geospatial-tracking/proto"
	device_usecase "github.com/aditya37/geospatial-tracking/usecase/device"
)

type Trackingdeliv struct {
	deviceCase *device_usecase.DeviceUsecase
	proto.UnimplementedGeotrackingServer
}

func NewTrackingDelivery(
	deviceCase *device_usecase.DeviceUsecase,
) *Trackingdeliv {
	return &Trackingdeliv{
		deviceCase: deviceCase,
	}
}

//
func (td *Trackingdeliv) GetDeviceLogByDeviceId(ctx context.Context, req *proto.RequestGetDeviceLogByDeviceId) (*proto.ResponseGetDeviceLogByDeviceId, error) {
	return &proto.ResponseGetDeviceLogByDeviceId{}, nil
}

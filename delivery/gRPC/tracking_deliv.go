package grpc

import (
	"context"

	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/geospatial-tracking/repository"
	device_usecase "github.com/aditya37/geospatial-tracking/usecase/device"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Trackingdeliv struct {
	deviceCase *device_usecase.DeviceUsecase
	repostream *repository.GPSChannelStream
	proto.UnimplementedGeotrackingServer
}

func NewTrackingDelivery(
	deviceCase *device_usecase.DeviceUsecase,
	repostream *repository.GPSChannelStream,
) *Trackingdeliv {
	return &Trackingdeliv{
		deviceCase: deviceCase,
		repostream: repostream,
	}
}

//
func (td *Trackingdeliv) GetDeviceLogByDeviceId(ctx context.Context, req *proto.RequestGetDeviceLogByDeviceId) (*proto.ResponseGetDeviceLogByDeviceId, error) {
	return &proto.ResponseGetDeviceLogByDeviceId{}, nil
}

//
func (td *Trackingdeliv) GetGPSTracking(req *emptypb.Empty, stream proto.Geotracking_GetGPSTrackingServer) error {
	streamCtx := stream.Context()
	for {
		select {
		case d := <-td.repostream.Result:
			if err := stream.Send(&proto.ResponseStreamGPSTracking{
				DeviceId: d.DeviceId,
				Lat:      float32(d.Lat),
				Long:     float32(d.Long),
				Status:   d.Status,
			}); err != nil {
				return err
			}

		case <-streamCtx.Done():
			return nil
		}
	}
}

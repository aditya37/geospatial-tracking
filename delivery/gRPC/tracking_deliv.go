package grpc

import (
	"context"

	"github.com/aditya37/geofence-service/util"
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

//GetDeviceLogByDeviceId....
func (td *Trackingdeliv) GetDeviceLogByDeviceId(ctx context.Context, req *proto.RequestGetDeviceLogByDeviceId) (*proto.ResponseGetDeviceLogByDeviceId, error) {
	return &proto.ResponseGetDeviceLogByDeviceId{}, nil
}

//GetGPSTracking...
func (td *Trackingdeliv) GetGPSTracking(req *emptypb.Empty, stream proto.Geotracking_GetGPSTrackingServer) error {
	streamCtx := stream.Context()
	for {
		select {
		case d := <-td.repostream.StreamGPSTrack:
			// open channel
			device, err := td.deviceCase.GetDeviceDetailByDeviceId(streamCtx, d.DeviceId)
			if err != nil {
				util.Logger().Error(err)
				continue
			}
			if err := stream.Send(
				&proto.ResponseStreamGPSTracking{
					DeviceId: d.DeviceId,
					Status:   d.Status,
					GpsData: &proto.GPSData{
						Lat:   float32(d.Lat),
						Long:  float32(d.Long),
						Speed: float32(d.Speed),
					},
					DeviceInfo: &device,
					Sensors: &proto.Sensor{
						SignalStrength: float32(d.Signal),
						Temp:           float32(d.Temp),
					},
				},
			); err != nil {
				return err
			}
		case errStream := <-td.repostream.ChanStreamError:
			return errStream
		case <-streamCtx.Done():
			td.repostream.Done()
			return nil
		}
	}
}
func (td *Trackingdeliv) GetDeviceCounter(ctx context.Context, in *emptypb.Empty) (*proto.ResponseGetDeviceCounter, error) {
	resp, err := td.deviceCase.GetDeviceCounter(ctx, in)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

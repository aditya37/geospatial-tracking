package grpc

import (
	"context"

	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/geospatial-tracking/repository"
	device_usecase "github.com/aditya37/geospatial-tracking/usecase/device"
	"github.com/aditya37/logger"

	chan_repo "github.com/aditya37/geospatial-tracking/repository/channel"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Trackingdeliv struct {
	deviceCase           *device_usecase.DeviceUsecase
	repostream           *repository.GPSChannelStream
	chanStreamDeviceById *chan_repo.MonitoringDeviceByIdPool
	proto.UnimplementedGeotrackingServer
}

func NewTrackingDelivery(
	deviceCase *device_usecase.DeviceUsecase,
	repostream *repository.GPSChannelStream,
	chanStreamDeviceById *chan_repo.MonitoringDeviceByIdPool,
) *Trackingdeliv {
	return &Trackingdeliv{
		deviceCase:           deviceCase,
		repostream:           repostream,
		chanStreamDeviceById: chanStreamDeviceById,
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
				logger.Error(err)
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
					DeviceInfo: device.Device,
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

// GetDeviceCounter...
func (td *Trackingdeliv) GetDeviceCounter(ctx context.Context, in *emptypb.Empty) (*proto.ResponseGetDeviceCounter, error) {
	resp, err := td.deviceCase.GetDeviceCounter(ctx, in)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeviceLogs...
func (td *Trackingdeliv) GetDeviceLogs(ctx context.Context, in *proto.RequestGetDeviceLogs) (*proto.ResponseGetDeviceLogs, error) {
	if in.Paging.ItemPerPage == 0 {
		in.Paging.ItemPerPage = 5
	}
	if in.Paging.Page == 0 {
		in.Paging.Page = 1
	}
	resp, err := td.deviceCase.GetDeviceLogs(ctx, in)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (td *Trackingdeliv) GetDeviceByDeviceId(ctx context.Context, in *proto.RequestGetDeviceByDeviceId) (*proto.ResponseGetDeviceByDeviceId, error) {
	resp, err := td.deviceCase.GetDeviceDetailByDeviceId(ctx, in.DeviceId)
	if err != nil {
		return &proto.ResponseGetDeviceByDeviceId{}, err
	}
	return &resp, nil
}

func (td *Trackingdeliv) DeviceQrCode(ctx context.Context, in *proto.RequestDeviceQrCode) (*proto.ResponseDeviceQrCode, error) {
	resp, err := td.deviceCase.DeviceQrCode(ctx, in)
	if err != nil {
		return &proto.ResponseDeviceQrCode{}, err
	}
	return &resp, nil
}

func (td *Trackingdeliv) GetSensorById(ctx context.Context, in *proto.RequestGetSemsorById) (*proto.ResponseGetSensorById, error) {
	return td.deviceCase.GetSensorById(ctx, in)
}

func (td *Trackingdeliv) GetListAttachedSensor(ctx context.Context, in *proto.RequestGetAttachedSensor) (*proto.ResponseGetAttachedSensor, error) {
	resp, err := td.deviceCase.GetListAttachedSensor(ctx, in)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

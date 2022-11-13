package device

import (
	"context"
	"fmt"
	"time"

	"github.com/aditya37/geospatial-tracking/delivery/middleware"
	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/geospatial-tracking/repository/channel"
)

func (du *DeviceUsecase) GetDataMonitoringByDeviceId(ctx context.Context, request *proto.RequestMonitoringDeviceById) error {

	// validate device id
	if _, err := du.deviceManagerRepo.GetDeviceByDeviceId(
		ctx,
		request.DeviceId,
	); err != nil {
		return err
	}

	// get streamer id from context
	streamerId := fmt.Sprintf(
		"%s",
		ctx.Value(middleware.CtxKeyStreamerId.ToString()),
	)
	resp, err := du.deviceManagerRepo.GetDataMonitoringByDeviceId(
		ctx,
		request.DeviceId,
	)
	if err != nil {
		if err.Error() == "device record is empty" {
			du.chanStreamMonitoring.Send(
				channel.MonitoringDeviceByIdClient{
					StreamerId: streamerId,
					Data:       &proto.ResponseMonitoringDeviceById{},
				},
			)
			return nil
		}
		return err
	}

	// send data to channel and stream...
	du.chanStreamMonitoring.Send(
		channel.MonitoringDeviceByIdClient{
			StreamerId: streamerId,
			Data: &proto.ResponseMonitoringDeviceById{
				DeviceLog: &proto.LogItem{
					DeviceId:   resp.DeviceId,
					Status:     resp.LogStatus,
					Reason:     resp.LogReason,
					RecordedAt: resp.LogRecordedAt.Format(time.RFC3339),
				},
				GpsTracking: &proto.GPSData{
					Speed: float32(resp.GpsSpeed),
				},
			},
		},
	)

	return nil
}

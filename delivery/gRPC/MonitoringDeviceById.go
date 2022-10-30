package grpc

import (
	"fmt"
	"log"
	"time"

	"github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/delivery/middleware"
	"github.com/aditya37/geospatial-tracking/proto"
	"github.com/aditya37/geospatial-tracking/repository/channel"
)

func (td *Trackingdeliv) MonitoringDeviceById(request *proto.RequestMonitoringDeviceById, stream proto.Geotracking_MonitoringDeviceByIdServer) error {
	ctx := stream.Context()

	streamerId := fmt.Sprintf(
		"%s",
		ctx.Value(middleware.CtxKeyStreamerId.ToString()),
	)
	chanPool := channel.MonitoringDeviceByIdClient{
		StreamerId: streamerId,
		Streamer:   stream,
	}
	td.chanStreamDeviceById.Register(&chanPool)

	// validate device id
	tt := time.NewTicker(660 * time.Millisecond)
	defer func() {
		log.Println("closing stream.....")
		td.chanStreamDeviceById.Close(&chanPool)
		tt.Stop()
	}()

	for {
		select {
		case <-tt.C:
			if err := td.deviceCase.GetDataMonitoringByDeviceId(
				ctx,
				request,
			); err != nil {
				util.Logger().Error(err)
				return err
			}
		case <-ctx.Done():
			return nil
		}

	}

}

package device

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
	getenv "github.com/aditya37/get-env"
)

func (du *DeviceUsecase) GetDeviceLogs(ctx context.Context, request *proto.RequestGetDeviceLogs) (proto.ResponseGetDeviceLogs, error) {
	resp, err := du.deviceManagerRepo.GetDeviceLogs(ctx, request)
	if err != nil {
		util.Logger().Error(err)
		return proto.ResponseGetDeviceLogs{}, err
	}
	if len(resp) <= 0 {
		return proto.ResponseGetDeviceLogs{}, errors.New("empty device logs")
	}
	var items []*proto.LogItem
	for _, v := range resp {
		items = append(items, &proto.LogItem{
			DeviceId:   v.DeviceId,
			Status:     v.Status,
			Reason:     v.Reason,
			RecordedAt: v.RecordedAt.UTC().Format(time.RFC3339),
		})
	}
	// notify to stream service...
	go func() {
		j, _ := json.Marshal(items)
		if err := du.gcpManager.Publish(
			context.Background(),
			entity.PublishParam{
				TopicName: getenv.GetString("DEVICE_LOG_STREAM_TOPIC", "stream-device-logs"),
				Message:   j,
			},
		); err != nil {
			util.Logger().Error(err)
			return
		}
	}()
	return proto.ResponseGetDeviceLogs{
		DeviceLogs: items,
	}, nil
}

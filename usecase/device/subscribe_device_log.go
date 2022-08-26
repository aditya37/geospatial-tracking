package device

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/usecase"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (du *DeviceUsecase) SubscribeDeviceLog(c mqtt.Client, m mqtt.Message) {
	ctx := context.Background()
	var payload usecase.DeviceLogPayload
	if err := json.Unmarshal(m.Payload(), &payload); err != nil {
		util.Logger().Error(err)
		return
	}
	// validate device id
	if _, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, payload.DeviceId); err != nil {
		util.Logger().Error(err)
		return
	}

	// set to utc for save your life...
	epochToTime := time.Unix(payload.ReocordedAt, 0).UTC()
	if err := du.deviceManagerRepo.InsertDeviceLog(
		ctx,
		entity.DeviceLog{
			DeviceId:       payload.DeviceId,
			Status:         payload.Status,
			Reason:         payload.Reason,
			SignalStrength: payload.SignalStrength,
			RecordedAt:     epochToTime,
		},
	); err != nil {
		util.Logger().Error(err)
		return
	}
}

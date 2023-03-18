package device

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/usecase"
	"github.com/aditya37/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (du *DeviceUsecase) SubscribeDeviceDetect(c mqtt.Client, m mqtt.Message) {
	var payload usecase.PayloadInsertDeviceDetect
	ctx := context.Background()
	if err := json.Unmarshal(m.Payload(), &payload); err != nil {
		logger.Error(err)
		return
	}
	detectTime := time.Unix(payload.DetectTime, 0)
	if err := du.deviceManagerRepo.InsertDeviceDetect(
		ctx,
		entity.DetectDevice{
			DeviceId: payload.DeviceId,
			Detect:   payload.Detect,
			Lat:      payload.Lat,
			Long:     payload.Long,
			DetectAt: detectTime,
		},
	); err != nil {
		logger.Error(err)
		return
	}
	m.Ack()
}

package device

import (
	"github.com/aditya37/geospatial-tracking/repository"
)

type DeviceUsecase struct {
	mqttmanager       repository.MqttManager
	deviceManagerRepo repository.DeviceManager
	gpChannelStream   *repository.GPSChannelStream
}

func NewDeviceUsecase(
	mqttmanager repository.MqttManager,
	deviceManagerRepo repository.DeviceManager,
	gpsChannelStream *repository.GPSChannelStream,
) *DeviceUsecase {
	return &DeviceUsecase{
		mqttmanager:       mqttmanager,
		deviceManagerRepo: deviceManagerRepo,
		gpChannelStream:   gpsChannelStream,
	}
}

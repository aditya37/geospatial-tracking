package device

import (
	"github.com/aditya37/geospatial-tracking/repository"
)

type DeviceUsecase struct {
	mqttmanager       repository.MqttManager
	deviceManagerRepo repository.DeviceManager
	gpChannelStream   *repository.GPSChannelStream
	cacheManager      repository.CacheManager
	gpsChanForward    *repository.TrackingForward
}

func NewDeviceUsecase(
	mqttmanager repository.MqttManager,
	deviceManagerRepo repository.DeviceManager,
	gpsChannelStream *repository.GPSChannelStream,
	cacheManager repository.CacheManager,
	gpsChanForward *repository.TrackingForward,
) *DeviceUsecase {
	return &DeviceUsecase{
		mqttmanager:       mqttmanager,
		deviceManagerRepo: deviceManagerRepo,
		gpChannelStream:   gpsChannelStream,
		cacheManager:      cacheManager,
		gpsChanForward:    gpsChanForward,
	}
}

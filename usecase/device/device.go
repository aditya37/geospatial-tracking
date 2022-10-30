package device

import (
	"github.com/aditya37/geospatial-tracking/repository"
	"github.com/aditya37/geospatial-tracking/repository/channel"
)

type DeviceUsecase struct {
	mqttmanager          repository.MqttManager
	deviceManagerRepo    repository.DeviceManager
	gpChannelStream      *repository.GPSChannelStream
	cacheManager         repository.CacheManager
	gpsChanForward       *repository.TrackingForward
	gcpManager           repository.Pubsub
	chanStreamMonitoring *channel.MonitoringDeviceByIdPool
}

func NewDeviceUsecase(
	mqttmanager repository.MqttManager,
	deviceManagerRepo repository.DeviceManager,
	gpsChannelStream *repository.GPSChannelStream,
	cacheManager repository.CacheManager,
	gpsChanForward *repository.TrackingForward,
	gcpManager repository.Pubsub,
	chanStreamMonitoring *channel.MonitoringDeviceByIdPool,
) *DeviceUsecase {
	return &DeviceUsecase{
		mqttmanager:          mqttmanager,
		deviceManagerRepo:    deviceManagerRepo,
		gpChannelStream:      gpsChannelStream,
		cacheManager:         cacheManager,
		gpsChanForward:       gpsChanForward,
		gcpManager:           gcpManager,
		chanStreamMonitoring: chanStreamMonitoring,
	}
}

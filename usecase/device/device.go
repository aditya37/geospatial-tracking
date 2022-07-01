package device

import (
	"github.com/aditya37/geospatial-tracking/repository"
)

type DeviceUsecase struct {
	mqttmanager       repository.MqttManager
	deviceManagerRepo repository.DeviceManager
}

func NewDeviceUsecase(
	mqttmanager repository.MqttManager,
	deviceManagerRepo repository.DeviceManager,
) *DeviceUsecase {
	return &DeviceUsecase{
		mqttmanager:       mqttmanager,
		deviceManagerRepo: deviceManagerRepo,
	}
}

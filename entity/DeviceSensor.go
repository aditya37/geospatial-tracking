package entity

import "time"

type DeviceSensor struct {
	Id         int64
	DeviceId   int64
	SensorId   int64
	CreatedAt  time.Time
	ModifiedAt time.Time
}

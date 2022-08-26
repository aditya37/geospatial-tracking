package entity

import "time"

type (
	Device struct {
		DeviceId   string
		MacAddress string
		DeviceType string
		ChipId     string
		I2cAddress string
		CreatedAt  time.Time
		ModifiedAt time.Time
	}
	ResultGetCount struct {
		ActivatedDevice  int64
		RecordedTracking int64
	}
	DeviceLog struct {
		Id             int64
		DeviceId       string
		Status         string
		Reason         string
		SignalStrength float64
		RecordedAt     time.Time
	}
)

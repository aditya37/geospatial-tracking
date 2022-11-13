package entity

import "time"

type (
	Device struct {
		Id         int64
		DeviceId   string
		MacAddress string
		DeviceType int
		ChipId     string
		I2cAddress string
		CreatedAt  time.Time
		ModifiedAt time.Time
	}
	ResultGetCount struct {
		ActivatedDevice  int64
		RecordedTracking int64
		DetectCount      int64
		DeviceId         string
		Type             string
		LastDetect       time.Time
	}
	DeviceLog struct {
		Id             int64
		DeviceId       string
		Status         string
		Reason         string
		SignalStrength float64
		RecordedAt     time.Time
	}
	DetectDevice struct {
		Id       int64
		DeviceId int64
		Detect   string
		Lat      float64
		Long     float64
		DetectAt time.Time
	}
	// ResultMonitoringDeviceById...
	ResultMonitoringDeviceById struct {
		Id                int64
		DeviceId          string
		LogStatus         string
		LogReason         string
		LogSignalStrength float64
		LogRecordedAt     time.Time
		GpsSpeed          float64
	}
)

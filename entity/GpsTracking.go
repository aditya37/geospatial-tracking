package entity

import (
	"encoding/json"
	"time"
)

type (
	GPSTracking struct {
		Id             int64
		DeviceId       string
		Waypoints      json.RawMessage
		Status         string
		SignalStrength float64
		Speed          float64
		Temp           float64
		CreatedAt      time.Time
		ModifiedAt     time.Time
		Lat            float64
		Long           float64
	}
)

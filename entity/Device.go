package entity

import "time"

type (
	Device struct {
		DeviceId   string
		MacAddress string
		DeviceType string
		ChipId     int
		I2cAddress string
		CreatedAt  time.Time
		ModifiedAt time.Time
	}
)

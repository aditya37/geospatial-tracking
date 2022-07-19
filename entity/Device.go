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
)

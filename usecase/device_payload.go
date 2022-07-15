package usecase

import "fmt"

type Status string

var (
	StatusValidDeviceId            Status = "VALID_DEVICE_ID"
	StatusSuccessRegister          Status = "SUCCESS_REGISTER_DEVICE"
	StatusFailedRegister           Status = "FAILED_REGISETER_DEVICE"
	StatusGPSTrackingStart         Status = "START_RECORD_TRACKING"
	StatusGPSTrackingStop          Status = "STOP"
	StatusGPSTracingRecordTracking Status = "TRACKING_RECORDED"
)

func (s Status) ToString() string {
	return fmt.Sprintf("%s", s)
}

type (
	MqttRegisterDevicePayload struct {
		Deviceid   string `json:"device_id"`
		MacAddress string `json:"mac_address"`
		DeviceType string `json:"device_type"`
		ChipId     int64  `json:"chip_id"`
		I2cAddress string `json:"i2c_address"`
		Timestamp  int64  `json:"timestamp"`
	}

	MqttRespRegisterDevice struct {
		Deviceid string `json:"device_id"`
		Status   string `json:"status"`
		Message  string `json:"message"`
	}

	MqttGpsTrackingPayload struct {
		DeviceId  string  `json:"device_id"`
		Speed     float64 `json:"speed"`
		Status    string  `json:"status"`
		Temp      float64 `json:"temp"`
		Lat       float64 `json:"lat"`
		Long      float64 `json:"long"`
		Signal    float64 `json:"signal"`
		Timestamp int64   `json:"timestamp"`
	}
)

package usecase

import (
	"fmt"
)

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
		Angle     float64 `json:"angle"`
		Timestamp int64   `json:"timestamp"`
		Altitude  float64 `json:"altitude"`
	}

	ForwardTrackingPayload struct {
		Message string                 `json:"message"`
		GpsData MqttGpsTrackingPayload `json:"gps_data"`
	}

	// MQTTRespTracking....
	Message struct {
		Value  string `json:"value"`
		Reason string `json:"reason"`
	}
	GPSData struct {
		Lat      float64 `json:"lat"`
		Long     float64 `json:"long"`
		Altitude float64 `json:"altitude"`
		Speed    float64 `json:"speed"`
		Angle    float64 `json:"angle"`
	}
	Sensor struct {
		Temp float64 `json:"temp"`
	}
	MQTTRespTracking struct {
		DeviceId    string  `json:"device_id"`
		Status      string  `json:"status"`
		RespMessage Message `json:"message"`
		GPSData     GPSData `json:"gps_data"`
		Sensors     Sensor  `json:"sensor_data"`
	}
)

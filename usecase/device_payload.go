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
	StatusLowSignal                Status = "LOW_SIGNAL"
	StatusHeartBeat                Status = "HEARTBEAT"
)

func (s Status) ToString() string {
	return fmt.Sprintf("%s", s)
}

type (
	MqttRegisterDevicePayload struct {
		Deviceid   string `json:"device_id"`
		MacAddress string `json:"mac_address"`
		DeviceType int    `json:"device_type"`
		ChipId     string `json:"chip_id"`
		I2cAddress string `json:"i2c_address"`
		Timestamp  int64  `json:"timestamp"`
	}

	MqttRespRegisterDevice struct {
		Deviceid string `json:"device_id"`
		Status   string `json:"status"`
		Message  string `json:"message"`
	}

	// GPS Tracking from device
	MqttGpsTrackingPayload struct {
		DeviceId  string  `json:"device_id"`
		Speed     float64 `json:"speed"`
		Status    string  `json:"status"`
		Temp      float64 `json:"temp"`
		Lat       float64 `json:"lat"`
		Long      float64 `json:"long"`
		Signal    float64 `json:"signal"`
		Angle     float64 `json:"angle"`
		Altitude  float64 `json:"altitude"`
		Timestamp int64   `json:"timestamp"`
	}

	// Payload for forward from gps tracking topic
	// to channel...
	ForwardTrackingPayload struct {
		Message string                 `json:"message"`
		GpsData MqttGpsTrackingPayload `json:"gps_data"`
	}

	// MQTTRespTracking....
	// Payload for send to device or mqtt
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
		Temp   float64 `json:"temp"`
		Signal float64 `json:"signal"`
	}
	MQTTRespTracking struct {
		DeviceId    string  `json:"device_id"`
		DeviceType  int     `json:"device_type"`
		Id          int64   `json:"id"`
		Status      string  `json:"status"`
		RespMessage Message `json:"message"`
		GPSData     GPSData `json:"gps_data"`
		Sensors     Sensor  `json:"sensor_data"`
	}

	// DeviceLogPayload...
	DeviceLogPayload struct {
		DeviceId       string  `json:"device_id"`
		Status         string  `json:"status"`
		Reason         string  `json:"reason"`
		SignalStrength float64 `json:"signal_strength"`
		ReocordedAt    int64   `json:"recorded_at"`
	}

	// PayloadInsertDeviceDetect
	PayloadInsertDeviceDetect struct {
		DeviceId   int64   `json:"device_id"`
		Detect     string  `json:"detect"`
		Lat        float64 `json:"lat"`
		Long       float64 `json:"long"`
		DetectTime int64   `json:"detect_time"`
	}
)

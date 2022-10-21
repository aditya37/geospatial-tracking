package device

import (
	"context"
	"encoding/json"

	"github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/usecase"
)

// forward/response data tracking to device
func (du *DeviceUsecase) ForwardGPSTracking(m *usecase.ForwardTrackingPayload) {
	ctx := context.Background()
	if m.Message != "" {
		msg := du.mappingMessage(m)

		device, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, m.GpsData.DeviceId)
		if err != nil {
			util.Logger().Error(err)
			return
		}

		payload := usecase.MQTTRespTracking{
			DeviceId:    m.GpsData.DeviceId,
			Status:      m.GpsData.Status,
			Id:          device.Id,
			RespMessage: msg,
			DeviceType:  device.DeviceType,
			GPSData: usecase.GPSData{
				Lat:      m.GpsData.Lat,
				Long:     m.GpsData.Long,
				Altitude: 0,
				Speed:    m.GpsData.Speed,
				Angle:    0,
			},
		}
		jsonMsg, _ := json.Marshal(payload)
		if err := du.mqttmanager.Publish(
			"/device/resp/tracking",
			byte(1),
			false,
			jsonMsg,
		); err != nil {
			util.Logger().Error(err)
			return
		}
	} else {
		ms := du.mappingMessageByStatus(m)

		device, err := du.deviceManagerRepo.GetDeviceByDeviceId(ctx, m.GpsData.DeviceId)
		if err != nil {
			util.Logger().Error(err)
			return
		}

		payload := usecase.MQTTRespTracking{
			DeviceId:    m.GpsData.DeviceId,
			Status:      m.GpsData.Status,
			RespMessage: ms,
			Id:          device.Id,
			DeviceType:  device.DeviceType,
			GPSData: usecase.GPSData{
				Lat:      m.GpsData.Lat,
				Long:     m.GpsData.Long,
				Altitude: m.GpsData.Altitude,
				Speed:    m.GpsData.Speed,
				Angle:    m.GpsData.Angle,
			},
			Sensors: usecase.Sensor{
				Temp:   m.GpsData.Temp,
				Signal: m.GpsData.Signal,
			},
		}
		jsonMsg, _ := json.Marshal(payload)
		if err := du.mqttmanager.Publish(
			"/device/resp/tracking",
			byte(1),
			false,
			jsonMsg,
		); err != nil {
			util.Logger().Error(err)
			return
		}
	}
}

// mapping message by status
func (du *DeviceUsecase) mappingMessageByStatus(m *usecase.ForwardTrackingPayload) usecase.Message {
	if m.GpsData.Status == usecase.StatusGPSTrackingStart.ToString() {
		return usecase.Message{
			Value:  "Starting tracking",
			Reason: "Starting Tracking",
		}
	} else if m.GpsData.Status == usecase.StatusGPSTracingRecordTracking.ToString() {
		return usecase.Message{
			Value:  "Record GPS Tracking",
			Reason: "Recording GPS Tracking",
		}
	} else {
		return usecase.Message{
			Value: m.GpsData.Status,
		}
	}
}

// mapping message for increase memory usage in microcontroller
func (du *DeviceUsecase) mappingMessage(m *usecase.ForwardTrackingPayload) usecase.Message {
	if m.Message == "Can't insert tracking with status TRACKING_RECORDED please start tracking" {
		return usecase.Message{
			Value:  "Can't insert tracking with status TRACKING_RECORDED please start tracking",
			Reason: "Please Start Tracking",
		}
	}
	return usecase.Message{}
}

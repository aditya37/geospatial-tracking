package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aditya37/geofence-service/util"
	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
	"github.com/aditya37/geospatial-tracking/usecase"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	geojson "github.com/paulmach/go.geojson"
)

func (du *DeviceUsecase) SubscribeGPSTracking(c mqtt.Client, m mqtt.Message) {
	ctx := context.Background()
	payload, err := du.unmarshallTrackingPayload(m.Payload())
	if err != nil {
		util.Logger().Error(err)
		return
	}

	status, err := du.getLastTrackingStatus(
		ctx,
		payload.DeviceId,
		20,
	)
	if err != nil {
		if err == repository.ErrLastTrackingNotFound {
			if validRequest := du.validateRequestPayloadBeforeInsert(payload, status); !validRequest {
				util.Logger().Error("Wrong request for insert")
				return
			} else {
				// do insert with last status "STOP"
				util.Logger().Info(fmt.Sprintf("Insert tracking from %s", payload.DeviceId))
				if err := du.insertGPSTracking(ctx, payload); err != nil {
					util.Logger().Error(err)
					return
				}
				m.Ack()
				return
			}
		}
		util.Logger().Error(err)
		return
	}

	if availableCurrentTrack := du.evaluateLastTrackingStatus(status.Status); !availableCurrentTrack {
		if validRequest := du.validateRequestPayloadBeforeInsert(payload, status); !validRequest {
			util.Logger().Error("Wrong request for insert")
			return
		} else {
			// do insert with last status "STOP"
			util.Logger().Info(fmt.Sprintf("Insert tracking with last status %s from %s", payload.DeviceId, status.Status))
			if err := du.insertGPSTracking(ctx, payload); err != nil {
				util.Logger().Error(err)
				return
			}
			m.Ack()
		}
	} else {
		// do update
		util.Logger().Info(fmt.Sprintf("Do tracking from %s", payload.DeviceId))
		if err := du.updateGPSTracking(ctx, payload, status.Id); err != nil {
			return
		}
		m.Ack()
	}
}

// validateRequestPayloadBeforeInsert
func (du *DeviceUsecase) validateRequestPayloadBeforeInsert(payload usecase.MqttGpsTrackingPayload, laststatus entity.GPSTracking) bool {
	if payload.Status == usecase.StatusGPSTrackingStop.ToString() {
		return false
	} else if payload.Status == usecase.StatusGPSTrackingStart.ToString() && laststatus.Status == usecase.StatusGPSTrackingStop.ToString() {
		return true
	} else {
		return true
	}
}

// updateGPSTracking....
func (du *DeviceUsecase) updateGPSTracking(ctx context.Context, payload usecase.MqttGpsTrackingPayload, idTracking int64) error {
	if err := du.deviceManagerRepo.UpdateTracking(ctx, entity.GPSTracking{
		Id:             idTracking,
		DeviceId:       payload.DeviceId,
		Lat:            payload.Lat,
		Long:           payload.Long,
		Status:         payload.Status,
		Speed:          payload.Speed,
		Temp:           payload.Temp,
		SignalStrength: payload.Signal,
	}); err != nil {
		util.Logger().Error(err)
		return err
	}
	return nil
}

// insertGPSTracking...
func (du *DeviceUsecase) insertGPSTracking(ctx context.Context, payload usecase.MqttGpsTrackingPayload) error {

	if payload.Status != usecase.StatusGPSTrackingStart.ToString() && payload.Status != usecase.StatusGPSTrackingStop.ToString() {
		return errors.New(fmt.Sprintf("Can't insert tracking with status %s please start tracking", payload.Status))
	}

	// do insert
	var latLong []float64
	latLong = append(latLong, payload.Long, payload.Lat)
	jsonLineString := du.generateLinestring(latLong)
	if _, err := du.deviceManagerRepo.InsertTracking(
		ctx,
		entity.GPSTracking{
			DeviceId:       payload.DeviceId,
			SignalStrength: payload.Signal,
			Speed:          payload.Speed,
			Status:         payload.Status,
			Temp:           payload.Temp,
			Waypoints:      jsonLineString,
		},
	); err != nil {
		util.Logger().Error(err)
		return err
	}
	return nil
}

// generateLinestring...
func (du *DeviceUsecase) generateLinestring(data []float64) []byte {
	coordinate := [][]float64{}
	coordinate = append(coordinate, data)
	gjson := geojson.NewLineStringFeature(coordinate)
	res, _ := gjson.MarshalJSON()
	return res
}

// getLastTrackingStatus...
func (du *DeviceUsecase) getLastTrackingStatus(ctx context.Context, device_id string, interval int) (entity.GPSTracking, error) {
	lastTracking, err := du.deviceManagerRepo.GetLastTrackingByInterval(
		ctx,
		device_id,
		interval,
	)
	if err != nil {
		util.Logger().Error(err)
		return entity.GPSTracking{}, err
	}

	return entity.GPSTracking{
		Id:     lastTracking.Id,
		Status: lastTracking.Status,
	}, nil

}

// evaluateLastTrackingStatus
func (du *DeviceUsecase) evaluateLastTrackingStatus(status string) bool {
	if status == usecase.StatusGPSTrackingStart.ToString() || status == usecase.StatusGPSTracingRecordTracking.ToString() {
		// do update last tracking
		return true
	} else {
		return false
	}
}

// unmarshallTrackingPayload...
func (du *DeviceUsecase) unmarshallTrackingPayload(data []byte) (usecase.MqttGpsTrackingPayload, error) {
	var payload usecase.MqttGpsTrackingPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		util.Logger().Error(err)
		return usecase.MqttGpsTrackingPayload{}, err
	}
	return payload, nil
}

package devicemanager

import (
	"context"

	"github.com/aditya37/geospatial-tracking/proto"
)

func (dm *device) GetAttachedSensorByDeviceId(ctx context.Context, device_id string) (proto.ResponseGetAttachedSensor, error) {
	arg := []interface{}{
		&device_id,
	}

	rows, err := dm.db.QueryContext(ctx, mysqlQueryGetAttachedSensorByDeviceId, arg...)
	if err != nil {
		return proto.ResponseGetAttachedSensor{}, err
	}
	defer rows.Close()

	var result proto.ResponseGetAttachedSensor
	for rows.Next() {
		var record proto.ResponseGetSensorById
		var deviceId string
		if err := rows.Scan(
			&deviceId,
			&record.Id,
			&record.SensorName,
			&record.Description,
			&record.SensorType,
		); err != nil {
			return proto.ResponseGetAttachedSensor{}, err
		}
		result.DeviceId = deviceId
		result.Sensor = append(result.Sensor, &record)
	}
	return result, nil
}

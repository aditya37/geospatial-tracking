package devicemanager

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (d *device) GetDataMonitoringByDeviceId(ctx context.Context, device_id string) (*entity.ResultMonitoringDeviceById, error) {
	arg := []interface{}{
		&device_id,
	}

	row := d.db.QueryRowContext(ctx, mysqlQueryGetDeviceMonitoringById, arg...)
	var record entity.ResultMonitoringDeviceById
	if err := row.Scan(
		&record.Id,
		&record.DeviceId,
		&record.LogStatus,
		&record.LogReason,
		&record.LogSignalStrength,
		&record.LogRecordedAt,
		&record.GpsSpeed,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("device record is empty")
		}
		return nil, err
	}

	return &record, nil
}

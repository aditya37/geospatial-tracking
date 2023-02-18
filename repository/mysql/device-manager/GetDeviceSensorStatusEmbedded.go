package devicemanager

import (
	"context"
	"database/sql"
	"errors"
)

func (d *device) GetDeviceSensorStatusEmbedded(ctx context.Context, sensorid, id_device int64) error {
	arg := []interface{}{
		&sensorid,
		&id_device,
	}

	row := d.db.QueryRowContext(ctx, mysqlQueryGetEmbeddedSensorStatus, arg...)
	var id int64
	if err := row.Scan(
		&id,
	); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("sensor not embedded in device")
		}
	}
	return nil
}

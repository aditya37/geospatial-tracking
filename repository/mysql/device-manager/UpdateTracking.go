package devicemanager

import (
	"context"
	"errors"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (d *device) UpdateTracking(ctx context.Context, data entity.GPSTracking) error {
	arg := []interface{}{
		&data.Long,
		&data.Lat,
		&data.Status,
		&data.Temp,
		&data.Speed,
		&data.SignalStrength,
		&data.Id,
		&data.DeviceId,
	}

	row, err := d.db.ExecContext(ctx, mysqlQueryUpdateTracking, arg...)
	if err != nil {
		return err
	}
	if isAffacted, _ := row.RowsAffected(); isAffacted == 0 {
		return errors.New("Failed update tracking data")
	}
	return nil
}

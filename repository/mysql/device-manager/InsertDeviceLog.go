package devicemanager

import (
	"context"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (d *device) InsertDeviceLog(ctx context.Context, data entity.DeviceLog) error {
	arg := []interface{}{
		&data.DeviceId,
		&data.Status,
		&data.Reason,
		&data.SignalStrength,
		&data.RecordedAt,
	}
	if _, err := d.db.ExecContext(ctx, mysqlQueryInsertDeviceLog, arg...); err != nil {
		return err
	}
	return nil
}

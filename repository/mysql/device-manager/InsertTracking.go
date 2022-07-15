package devicemanager

import (
	"context"
	"errors"

	"github.com/aditya37/geospatial-tracking/entity"
)

// TODO:Query insert json
func (d *device) InsertTracking(ctx context.Context, data entity.GPSTracking) (int64, error) {
	arg := []interface{}{
		&data.DeviceId,
		&data.SignalStrength,
		&data.Speed,
		&data.Status,
		&data.Temp,
		&data.Waypoints,
	}
	row, err := d.db.ExecContext(ctx, mysqlQueryInsertTracking, arg...)
	if err != nil {
		return 0, nil
	}
	if isAffacted, _ := row.RowsAffected(); isAffacted == 0 {
		return 0, errors.New("Failed insert tracking")
	}
	id, _ := row.LastInsertId()
	return id, nil
}

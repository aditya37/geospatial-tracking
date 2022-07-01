package devicemanager

import (
	"context"
	"database/sql"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
)

func (d *device) GetLastTrackingByInterval(ctx context.Context, deviceid string, interval int) (*entity.GPSTracking, error) {
	arg := []interface{}{
		&deviceid,
		&interval,
	}
	row := d.db.QueryRowContext(ctx, mysqlQueryGetLastTrackingByInterval, arg...)

	var record entity.GPSTracking
	if err := row.Scan(
		&record.Status,
		&record.Id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrLastTrackingNotFound
		}
		return nil, err
	}
	return &record, nil
}

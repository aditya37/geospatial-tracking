package devicemanager

import (
	"context"
	"database/sql"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
)

func (dm *device) GetDeviceByDeviceId(ctx context.Context, deviceid string) (*entity.Device, error) {
	arg := []interface{}{
		&deviceid,
	}
	row := dm.db.QueryRowContext(ctx, mysqlQueryGetDeviceByDeviceId, arg...)
	var record entity.Device
	if err := row.Scan(
		&record.DeviceId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrDeviceNotFound
		}
		return nil, err
	}
	return &record, nil
}

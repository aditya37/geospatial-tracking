package devicemanager

import (
	"context"
	"database/sql"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (dm *device) GetDeviceCounter(ctx context.Context) (*entity.ResultGetCount, error) {
	row := dm.db.QueryRowContext(ctx, mysqlQueryGetCounter)
	var record entity.ResultGetCount
	if err := row.Scan(
		&record.RecordedTracking,
	); err != nil {
		if err == sql.ErrNoRows {
			return &entity.ResultGetCount{}, nil
		}
		return nil, err
	}
	return &record, nil
}

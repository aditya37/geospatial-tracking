package devicemanager

import (
	"context"
	"time"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
)

func (d *device) GetDeviceLogs(ctx context.Context, data *proto.RequestGetDeviceLogs) ([]*entity.DeviceLog, error) {
	offset := (data.Paging.Page - 1) * data.Paging.ItemPerPage
	recordedAt := time.Unix(data.RecordedAt.Seconds, 0).UTC()
	arg := []interface{}{
		&recordedAt,
		&offset,
		&data.Paging.ItemPerPage,
	}

	rows, err := d.db.QueryContext(ctx, mysqlQueryGetDeviceLogs, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.DeviceLog
	for rows.Next() {
		var record entity.DeviceLog
		if err := rows.Scan(
			&record.DeviceId,
			&record.Status,
			&record.Reason,
			&record.RecordedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, &record)
	}
	return result, nil
}

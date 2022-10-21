package devicemanager

import (
	"context"
	"time"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
)

func (dm *device) GetCountDeviceDetect(ctx context.Context) ([]*proto.DetectDeviceItem, error) {

	rows, err := dm.db.QueryContext(ctx, mysqlQueryGetCountDeviceDetect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*proto.DetectDeviceItem
	for rows.Next() {
		var record entity.ResultGetCount
		if err := rows.Scan(
			&record.DetectCount,
			&record.DeviceId,
			&record.Type,
			&record.LastDetect,
		); err != nil {
			return nil, err
		}
		result = append(result, &proto.DetectDeviceItem{
			DeviceType: record.Type,
			Count:      record.DetectCount,
			LastDetect: record.LastDetect.Format(time.RFC3339),
		})
	}

	return result, nil
}

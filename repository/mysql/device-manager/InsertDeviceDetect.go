package devicemanager

import (
	"context"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (dm *device) InsertDeviceDetect(ctx context.Context, data entity.DetectDevice) error {
	arg := []interface{}{
		&data.DeviceId,
		&data.Detect,
		&data.Lat,
		&data.Long,
		&data.DetectAt,
	}
	if _, err := dm.db.ExecContext(ctx, mysqlQueryInsertDeviceDetect, arg...); err != nil {
		return err
	}
	return nil
}

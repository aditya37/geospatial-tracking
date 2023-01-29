package devicemanager

import (
	"context"
	"database/sql"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
)

func (dm *device) GetDeviceQrCode(ctx context.Context, data entity.QRDevice) (*entity.QRDevice, error) {
	arg := []interface{}{
		&data.DeviceId,
		&data.EventType,
	}

	row := dm.db.QueryRowContext(ctx, mysqlQueryGetDeviceQr, arg...)

	var record entity.QRDevice
	if err := row.Scan(
		&record.QrFile,
		&record.Url,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrDeviceNotFound
		}
		return nil, err
	}
	return &record, nil
}

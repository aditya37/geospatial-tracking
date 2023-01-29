package devicemanager

import (
	"context"
	"errors"

	"github.com/aditya37/geospatial-tracking/entity"
	sqldriver "github.com/go-sql-driver/mysql"
)

func (dm *device) InsertDeviceQr(ctx context.Context, data entity.QRDevice) error {
	arg := []interface{}{
		&data.EventType,
		&data.DeviceId,
		&data.Description,
		&data.QrFile,
		&data.Url,
	}
	if _, err := dm.db.ExecContext(ctx, mysqlQueryInsertQr, arg...); err != nil {
		if errCode, ok := err.(*sqldriver.MySQLError); ok {
			if errCode.Number == 1062 {
				return errors.New("qr code has been generated")
			}
		}
		return err
	}
	return nil
}

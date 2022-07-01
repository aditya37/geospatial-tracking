package devicemanager

import (
	"context"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
	sqldriver "github.com/go-sql-driver/mysql"
)

func (dm *device) InsertDevice(ctx context.Context, data entity.Device) error {
	arg := []interface{}{
		&data.DeviceId,
		&data.MacAddress,
		&data.DeviceType,
		&data.ChipId,
		&data.I2cAddress,
	}
	row, err := dm.db.ExecContext(
		ctx,
		mysqlQueryInsertDevice,
		arg...,
	)
	if err != nil {
		if errCode, ok := err.(*sqldriver.MySQLError); ok {
			if errCode.Number == 1062 {
				return repository.ErrDeviceHasBeenRegistered
			}
		}
		return err
	}
	if isAffacted, _ := row.RowsAffected(); isAffacted == 0 {
		return repository.ErrInsertDeviceNotAffacted
	}
	return nil
}

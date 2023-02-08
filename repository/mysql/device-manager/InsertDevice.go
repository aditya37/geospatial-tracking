package devicemanager

import (
	"context"
	"errors"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
	"github.com/aditya37/geospatial-tracking/usecase"
	sqldriver "github.com/go-sql-driver/mysql"
)

func (dm *device) InsertDevice(ctx context.Context, data entity.Device) (int64, error) {
	// insert or register device if network mode wlan
	if data.NetworkMode == usecase.NETWORK_MODE_WLAN {
		arg := []interface{}{
			&data.DeviceId,
			&data.MacAddress,
			&data.DeviceType,
			&data.ChipId,
			&data.I2cAddress,
			&data.NetworkMode,
		}
		return dm.registerDeviceNetworkModeWlan(ctx, arg)
		// register device if mode mobile data
	} else if data.NetworkMode == usecase.NETWORK_MODE_MOBILE_DATA {
		return dm.registerDeviceNetworkModeMobile(ctx, data)
	} else {
		return 0, errors.New("unknow network mode,failed register device")
	}
}

// registerDeviceNetworkModeWlan...
func (dm *device) registerDeviceNetworkModeWlan(ctx context.Context, arg []interface{}) (int64, error) {
	row, err := dm.db.ExecContext(
		ctx,
		mysqlQueryInsertDevice,
		arg...,
	)
	if err != nil {
		if errCode, ok := err.(*sqldriver.MySQLError); ok {
			if errCode.Number == 1062 {
				return 0, repository.ErrDeviceHasBeenRegistered
			}
		}
		return 0, err
	}
	if isAffacted, _ := row.RowsAffected(); isAffacted == 0 {
		return 0, repository.ErrInsertDeviceNotAffacted
	}
	return row.LastInsertId()
}

// registerDeviceNetworkModeMobile...
func (dm *device) registerDeviceNetworkModeMobile(ctx context.Context, data entity.Device) (int64, error) {
	tx, err := dm.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	arg := []interface{}{
		&data.DeviceId,
		&data.MacAddress,
		&data.DeviceType,
		&data.ChipId,
		&data.I2cAddress,
		&data.NetworkMode,
	}
	// insert to tbl mst_device....
	row, err := tx.ExecContext(
		ctx,
		mysqlQueryInsertDevice,
		arg...,
	)
	if err != nil {
		if errCode, ok := err.(*sqldriver.MySQLError); ok {
			if errCode.Number == 1062 {
				return 0, repository.ErrDeviceHasBeenRegistered
			}
		}
		return 0, err
	}
	// id after insert
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	argInsertSim := []interface{}{
		&id,
		&data.SIM.PhoneNo,
		&data.SIM.IMEI,
		&data.SIM.IMSI,
		&data.SIM.SIMOperator,
		&data.SIM.APN,
	}
	rowInsertSim, err := tx.ExecContext(ctx, mysqlQueryInsertSim, argInsertSim...)
	if err != nil {
		return 0, err
	}
	if isAffacted, _ := rowInsertSim.RowsAffected(); isAffacted == 0 {
		return 0, errors.New("failed register device")
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return row.LastInsertId()
}

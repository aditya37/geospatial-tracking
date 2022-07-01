package devicemanager

import (
	"database/sql"

	"github.com/aditya37/geospatial-tracking/repository"
)

type device struct {
	db *sql.DB
}

func NewDeviceManager(
	db *sql.DB,
) (repository.DeviceManager, error) {
	return &device{
		db: db,
	}, nil
}

func (dm *device) Close() error {
	return dm.db.Close()
}

package repository

import (
	"context"
	"errors"
	"io"

	"github.com/aditya37/geospatial-tracking/entity"
)

var (
	ErrInsertDeviceNotAffacted = errors.New("Failed insert device")
	ErrDeviceHasBeenRegistered = errors.New("Device has been registered")
	ErrDeviceNotFound          = errors.New("Device id not found")
	ErrLastTrackingNotFound    = errors.New("Last tracking not found")
)

type DeviceManager interface {
	io.Closer
	// Write....
	InsertDevice(ctx context.Context, data entity.Device) error
	InsertTracking(ctx context.Context, data entity.GPSTracking) (int64, error)
	UpdateTracking(ctx context.Context, data entity.GPSTracking) error
	// Read....
	GetDeviceByDeviceId(ctx context.Context, deviceid string) (*entity.Device, error)
	GetLastTrackingByInterval(ctx context.Context, deviceid string, interval int) (*entity.GPSTracking, error)
	GetDeviceCounter(ctx context.Context) (*entity.ResultGetCount, error)
}

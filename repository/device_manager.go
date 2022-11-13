package repository

import (
	"context"
	"errors"
	"io"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
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
	InsertDeviceLog(ctx context.Context, data entity.DeviceLog) error
	InsertDeviceDetect(ctx context.Context, data entity.DetectDevice) error
	GetCountDeviceDetect(ctx context.Context) ([]*proto.DetectDeviceItem, error)

	// Read....
	GetDeviceByDeviceId(ctx context.Context, deviceid string) (*entity.Device, error)
	GetLastTrackingByInterval(ctx context.Context, deviceid string, interval int) (*entity.GPSTracking, error)
	GetDeviceCounter(ctx context.Context) (*entity.ResultGetCount, error)
	GetDeviceLogs(ctx context.Context, data *proto.RequestGetDeviceLogs) ([]*entity.DeviceLog, error)
	GetDataMonitoringByDeviceId(ctx context.Context, device_id string) (*entity.ResultMonitoringDeviceById, error)
}

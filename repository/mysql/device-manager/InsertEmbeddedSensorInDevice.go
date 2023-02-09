package devicemanager

import (
	"context"
	"fmt"
	"strings"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (dm *device) InsertEmbeddedSensorInDevice(ctx context.Context, data []entity.DeviceSensor) error {
	arg := []interface{}{}
	stmt := []string{}

	for i, _ := range data {
		arg = append(arg, &data[i].DeviceId, &data[i].SensorId)
		stmt = append(stmt, "(?,?)")
	}
	query := fmt.Sprintf(
		mysqlQueryInsertEmbeddedSensor,
		strings.Join(stmt, ","),
	)

	if _, err := dm.db.ExecContext(ctx, query, arg...); err != nil {
		return err
	}
	return nil
}

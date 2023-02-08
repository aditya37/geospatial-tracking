package devicemanager

import (
	"context"
	"fmt"
	"strings"

	"github.com/aditya37/geospatial-tracking/entity"
)

func (dm *device) GetSensorById(ctx context.Context, sensorid []int) ([]*entity.Sensor, error) {
	stmt := []string{}
	arg := []interface{}{}
	for i, _ := range sensorid {
		arg = append(arg, &sensorid[i])
		stmt = append(stmt, "?")
	}
	query := fmt.Sprintf(
		mysqlQueryGetSensorById,
		strings.Join(stmt, ","),
	)

	rows, err := dm.db.QueryContext(ctx, query, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entity.Sensor
	for rows.Next() {
		var record entity.Sensor
		if err := rows.Scan(
			&record.Id,
		); err != nil {
			return nil, err
		}
		result = append(
			result,
			&entity.Sensor{
				Id: record.Id,
			},
		)
	}
	return result, nil
}

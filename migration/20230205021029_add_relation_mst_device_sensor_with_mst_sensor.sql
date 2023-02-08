-- +goose Up
-- +goose StatementBegin
ALTER TABLE 
	mst_device_sensor 
ADD CONSTRAINT 
	`fk_mst_device_sensor_mst_sensor`  
FOREIGN KEY (`sensor_id`) REFERENCES `mst_sensor` (`id`) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_device_sensor DROP CONSTRAINT `fk_mst_device_sensor_mst_sensor`;
-- +goose StatementEnd

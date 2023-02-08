-- +goose Up
-- +goose StatementBegin
ALTER TABLE 
	mst_device_sensor 
ADD CONSTRAINT 
	`fk_mst_device_sensor_mst_device`  
FOREIGN KEY (`device_id`) REFERENCES `mst_device` (`id`) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_device_sensor DROP CONSTRAINT `fk_mst_device_sensor_mst_device`;
-- +goose StatementEnd

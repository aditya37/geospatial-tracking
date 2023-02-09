-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_device_sensor` (
	   id BIGINT NOT NULL AUTO_INCREMENT,
	   device_id BIGINT NOT NULL,
	   sensor_id BIGINT NOT NULL,
	   created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   	   modified_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	   PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_device_sensor`;
-- +goose StatementEnd

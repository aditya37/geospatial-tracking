-- +goose Up
DROP TABLE IF EXISTS `trx_gps_tracking`;
-- +goose StatementBegin
CREATE TABLE `trx_gps_tracking`(
	   id BIGINT NOT NULL AUTO_INCREMENT,
	   device_id VARCHAR(100),
	   waypoints JSON,
	   status VARCHAR(200) NOT NULL,
	   signal_strength FLOAT,
	   speed FLOAT,
	   temp FLOAT,
	   created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	   modified_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	   PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `trx_gps_tracking`;
-- +goose StatementEnd

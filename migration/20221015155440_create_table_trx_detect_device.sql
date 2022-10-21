-- +goose Up
-- +goose StatementBegin
CREATE TABLE `trx_detect_device` (
	   id BIGINT NOT NULL AUTO_INCREMENT,
	   device_id BIGINT,
	   detect VARCHAR(32),
	   lat FLOAT,
	   `long` FLOAT,
	   detected_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	   PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `trx_detect_device`;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_qr_code` (
	   id BIGINT NOT NULL AUTO_INCREMENT,
	   event_type INT(4),
	   device_id VARCHAR(100) NOT NULL, 
	   description VARCHAR(100),
	   qr_file VARCHAR(100),
	   url VARCHAR(700),
	   PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_qr_code`;
-- +goose StatementEnd

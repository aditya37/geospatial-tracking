-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_sensor` (
	   id BIGINT NOT NULL AUTO_INCREMENT,
	   sensor_name VARCHAR(32) NOT NULL,
	   description VARCHAR(100),
	   PRIMARY KEY(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_sensor`;
-- +goose StatementEnd

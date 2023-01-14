-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_sim` (
    id BIGINT NOT NULL AUTO_INCREMENT,
    device_id BIGINT NOT NULL,
    phone_no VARCHAR(32),
    imei VARCHAR(32),
    imsi VARCHAR(32),
    sim_operator BIGINT NOT NULL,
    apn VARCHAR(32),
    `status` BOOLEAN NOT NULL DEFAULT true,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_sim`;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_sim_operator` (
    id BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32),
    `status` BOOLEAN NOT NULL DEFAULT true,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id),
    UNIQUE KEY (`name`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_sim_operator`;
-- +goose StatementEnd

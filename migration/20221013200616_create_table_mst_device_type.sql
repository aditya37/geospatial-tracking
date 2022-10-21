-- +goose Up
DROP TABLE IF EXISTS `mst_device_type`;
-- +goose StatementBegin
CREATE TABLE `mst_device_type` (
  `id` bigint NOT NULL DEFAULT 0,
  `type` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `mst_device_type`;
-- +goose StatementEnd

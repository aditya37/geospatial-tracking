-- +goose Up
DROP TABLE IF EXISTS `trx_device_log`;
-- +goose StatementBegin
CREATE TABLE `trx_device_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `device_id` varchar(100) NOT NULL,
  `status` varchar(100) NOT NULL,
  `reason` varchar(200) NOT NULL,
  `signal_strength` int NOT NULL DEFAULT '0',
  `recorded_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `trx_device_log`;
-- +goose StatementEnd

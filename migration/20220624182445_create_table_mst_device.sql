-- +goose Up
DROP TABLE IF EXISTS `mst_device`;
-- +goose StatementBegin
CREATE TABLE `mst_device` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `device_id` varchar(100) NOT NULL,
  `mac_address` varchar(50) DEFAULT NULL,
  `device_type` varchar(50) DEFAULT NULL,
  `chip_id` int NOT NULL DEFAULT '0',
  `i2c_address` varchar(5) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `device_id` (`device_id`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_device`;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_device ADD COLUMN network_mode enum("MOBILE_DATA","WLAN") NOT NULL DEFAULT 'WLAN';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_device DROP COLUMN network_mode;
-- +goose StatementEnd

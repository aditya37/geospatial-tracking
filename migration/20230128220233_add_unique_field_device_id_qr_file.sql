-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_qr_code ADD UNIQUE(`device_id`,`qr_file`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_qr_code ADD UNIQUE(`device_id`,`qr_file`);
-- +goose StatementEnd

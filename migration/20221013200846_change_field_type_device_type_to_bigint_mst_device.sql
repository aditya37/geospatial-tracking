-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_device MODIFY COLUMN device_type BIGINT; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_device MODIFY COLUMN device_type VARCHAR(20); 
-- +goose StatementEnd

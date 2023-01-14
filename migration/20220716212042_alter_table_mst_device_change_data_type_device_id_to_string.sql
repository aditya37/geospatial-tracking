-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_device MODIFY COLUMN chip_id VARCHAR(20); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_device MODIFY COLUMN chip_id INTEGER(10); 
-- +goose StatementEnd
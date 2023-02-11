-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_sensor ADD COLUMN sensor_type enum("TEMPERATURE","ELECTRICAL") NOT NULL DEFAULT 'TEMPERATURE';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_sensor DROP COLUMN sensor_type;
-- +goose StatementEnd

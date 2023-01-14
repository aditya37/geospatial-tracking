-- +goose Up
-- +goose StatementBegin
ALTER TABLE 
	mst_sim 
ADD CONSTRAINT 
	`fk_mst_sim_mst_device`  
FOREIGN KEY (`device_id`) REFERENCES `mst_device` (`id`) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_sim DROP CONSTRAINT `fk_mst_sim_mst_device`;
-- +goose StatementEnd
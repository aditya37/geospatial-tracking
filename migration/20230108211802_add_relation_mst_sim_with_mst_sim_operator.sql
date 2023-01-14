-- +goose Up
-- +goose StatementBegin
ALTER TABLE 
	mst_sim 
ADD CONSTRAINT 
	`fk_mst_sim_mst_operator`  
FOREIGN KEY (`sim_operator`) REFERENCES `mst_sim_operator` (`id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_sim DROP CONSTRAINT `fk_mst_sim_mst_operator`;
-- +goose StatementEnd
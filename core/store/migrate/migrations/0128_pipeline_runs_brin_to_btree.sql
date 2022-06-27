-- +goose Up
	-- Migrate idx_pipeline_runs_created_at to BTREE
    DROP INDEX IF EXISTS idx_pipeline_runs_created_at;
    CREATE INDEX idx_pipeline_runs_created_at on public.pipeline_runs BTREE (created_at);

	-- Migrate idx_pipeline_runs_finished_at to BTREE
    DROP INDEX IF EXISTS idx_pipeline_runs_finished_at;
    CREATE INDEX idx_pipeline_runs_finished_at on public.pipeline_runs BTREE (finished_at);

    -- Migrate idx_pipeline_runs_pipeline_spec_id to HASH index
    DROP INDEX IF EXISTS idx_pipeline_runs_pipeline_spec_id 
    CREATE INDEX idx_pipeline_runs_pipeline_spec_id on public.pipeline_runs HASH (pipeline_spec_id);

-- +goose Down
    DROP INDEX IF EXISTS idx_pipeline_runs_created_at;
    CREATE INDEX idx_pipeline_runs_created_at on public.pipeline_runs BRIN (created_at);

    DROP INDEX IF EXISTS idx_pipeline_runs_finished_at;
    CREATE INDEX idx_pipeline_runs_finished_at on public.pipeline_runs BRIN (finished_at);

    DROP INDEX IF EXISTS idx_pipeline_runs_pipeline_spec_id 
    CREATE INDEX idx_pipeline_runs_pipeline_spec_id on public.pipeline_runs BTREE (pipeline_spec_id);

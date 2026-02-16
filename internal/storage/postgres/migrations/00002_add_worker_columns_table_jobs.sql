-- +goose Up
-- +goose StatementBegin
alter table jobs 
    add column worker_id varchar(100),
    add column claimed_at timestamp,
    add column completed_at timestamp,
    add column failed_at timestamp,
    add column error_message text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table job
    drop column if exists worker_id,
    add column if exists failed_at,
    add column if exists completed_at,
    add column if exists claimed_at,
    add column if exists worker_id;
-- +goose StatementEnd

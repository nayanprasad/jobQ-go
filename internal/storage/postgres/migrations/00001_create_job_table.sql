-- +goose Up
-- +goose StatementBegin

create type job_type as enum ('email.send', 'webhook.deliver');
create type job_status as enum ('pending', 'running', 'completed', 'failed', 'dead');

create table if not exists jobs (
    id serial primary key,
    type job_type not null,
    payload jsonb not null,
    status job_status not null,
    available_at timestamp not null,
    max_retries int not null,
    retry_count int not null default 0,
    priority int not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp default null   
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists jobs;
drop type if exists job_type;
drop type if exists job_status;
-- +goose StatementEnd

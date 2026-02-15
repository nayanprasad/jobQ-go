-- +goose Up
-- +goose StatementBegin

create type job_type as enum ('email.send', 'webhook.deliver');
create type job_status as enum ('pending', 'running', 'completed', 'failed', 'dead');

create table jobs if not exists (
    id serial primary key,
    type job_type not null,
    payload jsonb not null,
    status job_status not null,
    available_at timestamp not null,
    max_retries int not null,
    retry_count int not null,
    priority int not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp null   
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table jobs if exists;
drop type job_type if exists;
drop type job_status if exists;
-- +goose StatementEnd

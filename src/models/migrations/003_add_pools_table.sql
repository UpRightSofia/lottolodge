-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists pools (
    id uuid not null default uuid_generate_v4(),
    details JSON not null,
    updated_at timestamp not null default CURRENT_TIMESTAMP,
    constraint primary key(id)
);

-- +goose Down
drop table pools;
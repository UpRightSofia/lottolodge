-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists pools (
    id uuid not null default uuid_generate_v4(),
    details json,
    created_at timestamptz not null default CURRENT_TIMESTAMP,
    is_active boolean not null default true,
    constraint pk_pools primary key (id)
);

CREATE INDEX IF NOT EXISTS idx_pools_is_active_created ON pools (is_active, created_at);

-- +goose Down
drop table pools;
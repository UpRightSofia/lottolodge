-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists users (
    id uuid not null default uuid_generate_v4(),
    updated_at timestamp not null default CURRENT_TIMESTAMP,
    balance_e5 bigint not null default 0,
    constraint pk_users primary key (id)
);

-- +goose Down
drop table users;
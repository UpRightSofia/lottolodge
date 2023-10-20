-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists users (
    id uuid primary key default uuid_generate_v4() not null
);

-- +goose Down
drop table huds;
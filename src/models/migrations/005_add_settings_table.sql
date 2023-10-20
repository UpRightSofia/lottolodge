-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists settings (
    id serial primary key,
    ticket_prize_e5 bigint not null default 0,
    payout_percent smallint not null default 0,
    updated_at timestamptz not null default CURRENT_TIMESTAMP
);

-- +goose Down
drop table settings;
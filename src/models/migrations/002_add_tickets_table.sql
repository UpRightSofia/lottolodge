-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists tickets (
    id uuid not null default uuid_generate_v4(),
    user_id not null,
    details JSON not null,
    is_hand_picked boolean not null default false,
    updated_at timestamp not null default CURRENT_TIMESTAMP,
    constraint primary key(id),
    constraint foreign key(user_id) references users(id)
);

-- +goose Down
drop table tickets;
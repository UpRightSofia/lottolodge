-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists winnings (
    id uuid not null default uuid_generate_v4(),
    user_id uuid not null,
    ticket_id uuid not null,
    pool_id uuid not null,
    prize_e5 bigint not null default 0,
    updated_at timestamp not null default CURRENT_TIMESTAMP,
    constraint primary key(id),
    constraint foreign key(user_id) references users(id),
    constraint foreign key(ticket_id) references tickets(id),
    constraint foreign key(pool_id) references pools(id)
);

-- +goose Down
drop table winnings;
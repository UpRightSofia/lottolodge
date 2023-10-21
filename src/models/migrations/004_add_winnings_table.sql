-- +goose Up
create table if not exists winnings (
    id uuid not null default uuid_generate_v4(),
    user_id uuid not null,
    ticket_id uuid not null,
    pool_id uuid not null,
    prize_e5 bigint not null default 0,
    updated_at timestamptz not null default CURRENT_TIMESTAMP,
    constraint pk_winnings primary key (id),
    constraint fk_users foreign key (user_id) references users (id),
    constraint fk_tickets foreign key (ticket_id) references tickets (id),
    constraint fk_pools foreign key (pool_id) references pools (id)
);

-- +goose Down
drop table winnings;
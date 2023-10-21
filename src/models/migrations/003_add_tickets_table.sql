-- +goose Up
create table if not exists tickets (
    id uuid not null default uuid_generate_v4() primary key,
    user_id uuid not null,
    pool_id uuid not null,
    details json not null,
    is_hand_picked boolean not null default false,
    is_used bool not null default false,
    updated_at timestamptz not null default CURRENT_TIMESTAMP,
    constraint fk_users foreign key (user_id) references users (id),
    constraint fk_pool foreign key (pool_id) references pools (id)
);

CREATE INDEX IF NOT EXISTS idx_tickets_user_id_pool ON tickets (user_id, pool_id);

create index if not exists idx_is_used on tickets (is_used);

create index if not exists idx_pool_id on tickets (pool_id);

-- +goose Down
drop table tickets;

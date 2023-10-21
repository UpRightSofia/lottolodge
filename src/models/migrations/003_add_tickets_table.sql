-- +goose Up
create table if not exists tickets (
    id uuid not null default uuid_generate_v4(),
    user_id uuid not null,
    pool_id uuid not null,
    details json not null,
    is_hand_picked boolean not null default false,
    updated_at timestamptz not null default CURRENT_TIMESTAMP,
    constraint pk_tickets primary key (id),
    constraint fk_users foreign key (user_id) references users (id),
    constraint fk_pool foreign key (pool_id) references pools (id)
);

CREATE INDEX IF NOT EXISTS idx_tickets_user_id_pool ON tickets (user_id, pool_id);

-- +goose Down
drop table tickets;
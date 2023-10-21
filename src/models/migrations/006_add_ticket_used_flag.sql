-- +goose Up
alter table tickets
add column is_used bool default false;

create index if not exists idx_is_used
on tickets (is_used);

create index if not exists idx_pool_id
on tickets (pool_id);

-- +goose Down
drop index idx_is_used;
drop index idx_pool_id;

alter table tickets
drop column is_used;
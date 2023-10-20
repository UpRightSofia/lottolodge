-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists usertickets (
    user_id uuid not null,
    ticket_id uuid not null,
    constraint primary key(user_id, ticket_id),
    constraint foreign key(user_id) references users(id),
    constraint foreign key(ticket_id) references tickets(id)
);
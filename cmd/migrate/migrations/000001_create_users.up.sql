create extension if not exists citext;

create table if not exists users (
    id bigserial primary key,
    email citext unique not null,
    username varchar(255) unique not null,
    password bytea not null,
    timestamep timestamp(0) with time zone not null default now()
)
create extension if not exists citext;

create table if not exists posts (
    id bigserial primary key,
    title text not null,
    user_id bigint not null,
    content text not null,
    created_at timestamp(0) with time zone not null default now()
)
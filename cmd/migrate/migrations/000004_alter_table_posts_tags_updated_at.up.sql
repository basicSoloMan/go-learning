alter table posts
add column tags varchar(100) [];

alter table posts
add column updated_at timestamp(0) with time zone not null default now();
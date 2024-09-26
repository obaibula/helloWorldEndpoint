create table if not exists accounts
(
    id            bigserial primary key,
    first_name    text          not null,
    last_name     text          not null,
    email         text unique not null,
    birthday      date          not null,
    creation_date date          not null default date(now()),
    balance       decimal       not null default 0,
    version       integer       not null default 1
);
begin;

create table if not exists users_ (
    id_ integer primary key generated always as identity,
    username_ varchar(50) unique not null,
    email_ varchar(100) unique not null,
    password_ varchar(100) not null,
    is_admin_ boolean default false not null
);

create table if not exists content_ (
    id_ integer primary key generated always as identity,
    title_ varchar(255) not null,
    subtitle_ varchar(255) not null,
    slug_ varchar(255) not null,
    body_ text,
    created_at_ timestamp without time zone default current_timestamp not null,
    updated_at_ timestamp without time zone default current_timestamp not null,
    type_ varchar(4) check (type_ in ('page', 'post')) default 'post' not null,
    user_id_ integer references users_(id_)
);

create table if not exists sessions_ (
    id_ integer primary key generated always as identity,
    token_ text,
    expires_at_ bigint,
    issued_at_ bigint,
    user_id_ integer references users_(id_)
);

commit;


create table config_ (
    id integer primary key generated always as identity,
    name_ varchar(50),
    value_ varchar(255)
);

create table user_ (
    id integer primary key generated always as identity,
    username_ varchar(50) unique not null,
    email_ varchar(50) unique not null,
    is_admin_ boolean not null default false,
    password_ varchar(50) not null,
    registered_at_ time not null, 
    last_login_ time, 
    role_ varchar(10),
    profile_ varchar(255)
);

create table meta_ (
    id integer primary key generated always as identity,
    name_ varchar(50),
    value_ varchar(255)
);

create table user_meta_ (
    id integer primary key generated always as identity,
    user_id_ integer references user_(id),
    meta_id_ integer references meta_(id)
);

create table template_ (
    id integer primary key generated always as identity,
    name_ varchar(50),
    tmpl_ varchar(255)
);

create table category_ (
    id integer primary key generated always as identity,
    name_ varchar(50),
    template_id_ integer references template_(id)
);

create table page_ (
    id integer primary key generated always as identity,
    title_ varchar(255) not null,
    body_ text,
    slug_ varchar(255) not null,
    status_ varchar(10) not null default 'draft',
    created_at_ timestamp without time zone not null default current_timestamp,
    published_at_ timestamp without time zone,
    updated_at_ timestamp without time zone not null default current_timestamp,
    user_id_ integer references user_(id),
    category_id_ integer references category_(id)
);

create table page_meta_ (
    id integer primary key generated always as identity,
    page_id_ integer references page_(id),
    meta_id_ integer references meta_(id)
);

create table category_meta_ (
    id integer primary key generated always as identity,
    category_id_ integer references category_(id),
    meta_id_ integer references meta_(id)
);

create table post_ (
    id integer primary key generated always as identity,
    title_ varchar(255),
    subtitle_ varchar(255),
    body_ text,
    slug_ varchar(255) not null,
    status_ varchar(10) not null default 'draft',
    created_at_ timestamp without time zone not null default current_timestamp,
    published_at_ timestamp without time zone,
    updated_at_ timestamp without time zone not null default current_timestamp,
    user_id_ integer references user_(id),
    category_id_ integer references category_(id)
);


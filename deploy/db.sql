create database nixpigweb_;

# ------------------

create table content_ (
	id_ integer primary key generated always as identity,
	title_ varchar(255) not null,
	subtitle_ varchar(255) not null,
	slug_ varchar(255) not null,
	body_ text,
	created_at_ timestamp without time zone default current_timestamp not null,
	updated_at_ timestamp without time zone default current_timestamp not null,
	type_ varchar(4) check (type_ in ('page', 'post')) default 'post' not null
);

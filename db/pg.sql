---------------------------------------------------------------------------------
--                      Copyright (C) 2016 wystan
--
--        filename: pg.sql
--     description: 
--         created: 2016-02-24 17:07:18
--          author: wystan
-- 
---------------------------------------------------------------------------------

drop database if exists fixer;
create database fixer;
\c fixer;

drop user if exists pgtest;
create user pgtest with CREATEDB LOGIN PASSWORD '123456';

create table users (
    id serial not null,
    nicky varchar(64) not null,
    email varchar(64) not null,
    pwd char(128) not null,
    portrait varchar(128),
    register_date timestamp,
    last_login_time timestamp,

    primary key(id)
);

create table team (
    id serial not null, 
    name varchar(64) not null,
    leader int not null,
    goal varchar(1024),
    created_date timestamp,
    bug_table char(64),
    bug_table_status char,
    status char,
    logo varchar(128),

    primary key(id)
);

create unique index users_nicky on users(nicky);
create unique index team_name on team(name);
create index team_leader on team(leader);

alter table users owner to pgtest;
alter table team owner to pgtest;
alter sequence users_id_seq owner to pgtest;
alter sequence team_id_seq owner to pgtest;

---------------------------------------------------------------------------------
-- test data here

insert into users(nicky, email, pwd, register_date, last_login_time) 
       values ('john', 'john@gmail.com', '123456', now(), now());
insert into users(nicky, email, pwd, register_date, last_login_time) 
       values ('sherlock', 'sherlock@gmail.com', '123456', now(), now());


insert into team(name, leader, created_date, bug_table, bug_table_status, status) 
    values('frog', 1, now(), '3eefab', '1', '1');
insert into team(name, leader, created_date, bug_table, bug_table_status, status) 
    values('lion', 2, now(), '3aefab', '1', '1');

select * from users order by id;
select * from team order by id;
---------------------------------------------------------------------------------

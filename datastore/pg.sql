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
    pwd char(32) not null,
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
    bug_table char(32),
    bug_table_status char,
    status char,
    logo varchar(128),

    primary key(id)
);

create table user_team (
    user_id int not null,
    team_id int not null,
    join_date timestamp not null,

    primary key(user_id, team_id)
);

create unique index users_nicky on users(nicky);
create unique index team_name on team(name);
create index team_leader on team(leader);

alter table users owner to pgtest;
alter table team owner to pgtest;
alter sequence users_id_seq owner to pgtest;
alter sequence team_id_seq owner to pgtest;

---------------------------------------------------------------------------------

---------------------------------------------------------------------------------
--                      Copyright (C) 2016 wystan
--
--        filename: pg-data.sql
--     description: test data generation
--         created: 2016-02-24 20:07:19
--          author: wystan
-- 
---------------------------------------------------------------------------------

\c fixer;

-- Add new users
insert into users(nicky, email, pwd, register_date, last_login_time, portrait) 
       values ('john', 'john@gmail.com', '123456', now(), now(), 'static/images/2.jpg');
insert into users(nicky, email, pwd, register_date, last_login_time, portrait) 
       values ('sherlock', 'sherlock@gmail.com', '123456', now(), now(), 'static/images/1.jpg');

-- Add new teams
insert into team(name, leader_id, created_date, bug_table, bug_table_status, status) 
    values('john-frog', 1, now(), '1aaaaa', 1, 1);
insert into team(name, leader_id, created_date, bug_table, bug_table_status, status) 
    values('john-shark', 1, now(), '2aaaaa', 1, 1);
insert into team(name, leader_id, created_date, bug_table, bug_table_status, status) 
    values('john-whale', 1, now(), '3aaaaa', 1, 1);
insert into team(name, leader_id, created_date, bug_table, bug_table_status, status, goal) 
    values('sherlock-lion', 2, now(), '1bbbbb', 1, 0, 'lion team');
insert into team(name, leader_id, created_date, bug_table, bug_table_status, status, goal) 
    values('sherlock-tiger', 2, now(), '2bbbbb', 1, 1, 'tiger team');

-- Add user team relationships
insert into user_team values(1,4,now());
insert into user_team values(1,5,now());
insert into user_team values(2,1,now());
insert into user_team values(2,2,now());
insert into user_team values(2,3,now());

select id, nicky, email, pwd from users order by id;
select id, name, leader_id, status, bug_table from team order by id;
select * from user_team order by user_id, team_id;


---------------------------------------------------------------------------------

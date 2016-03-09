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

BEGIN;

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

-- Add bugs and log

CREATE TABLE bugs_1aaaaa (
    id              SERIAL NOT NULL,
    created_by      INT NOT NULL,
    current_handler INT NOT NULL,
    priority        INT NOT NULL,
    status          INT, 
    created_time    TIMESTAMP,
    last_update     TIMESTAMP,
    title           VARCHAR(128) not null,
    attachments     VARCHAR(512), 
    detail          VARCHAR(4096), 

    PRIMARY KEY(id)
);

CREATE TABLE buglog_1aaaaa (
    bug_id          INT NOT NULL,
    who             INT NOT NULL,
    action_type     INT NOT NULL,
    action_time     TIMESTAMP,
    action          VARCHAR(4096) 
);

CREATE INDEX idx_bugs_createdby ON bugs_1aaaaa(created_by);
CREATE INDEX idx_bugs_priority ON bugs_1aaaaa(priority);
CREATE INDEX idx_bugs_handler ON bugs_1aaaaa(current_handler);
CREATE INDEX idx_bugs_status ON bugs_1aaaaa(status);
CREATE INDEX idx_bugs_created_time ON bugs_1aaaaa(created_time);
CREATE INDEX idx_bugs_last_update ON bugs_1aaaaa(last_update);
CREATE INDEX idx_buglog_bugid ON buglog_1aaaaa(bug_id);

ALTER TABLE bugs_1aaaaa OWNER TO fixer;
ALTER TABLE buglog_1aaaaa OWNER TO fixer;

insert into bugs_1aaaaa
    (created_by, current_handler, priority, status, created_time, last_update, title, attachments, detail) 
    values(1, 2, 1, 1, now(), now(), 'system crashed', null, 
    'the program coredumped, need to fixed AS QUICKLY AS POSSIBLE');
insert into bugs_1aaaaa
    (created_by, current_handler, priority, status, created_time, last_update, title, attachments, detail) 
    values(2, 1, 1, 1, now(), now(), 'we can not login to the system', null, 
    'login page can not be shown, so fix it.');
insert into buglog_1aaaaa(bug_id, who, action_type, action_time, action)
    values(1, 1, 1, now(), 'assigned to user sherklock(2)');
insert into buglog_1aaaaa(bug_id, who, action_type, action_time, action)
    values(2, 1, 1, now(), 'assigned to user john(1)');

select id, nicky, email, pwd from users order by id;
select id, name, leader_id, status, bug_table from team order by id;
select * from user_team order by user_id, team_id;
select id, created_by, current_handler, priority, status, title from bugs_1aaaaa;
select * from buglog_1aaaaa;

-- commit tx
COMMIT;


---------------------------------------------------------------------------------

---------------------------------------------------------------------------------
--                      Copyright (C) 2016 wystan
--
--        filename: pg.sql
--     description: 
--         created: 2016-02-24 17:07:18
--          author: wystan
-- 
---------------------------------------------------------------------------------

DROP DATABASE IF EXISTS fixer;
CREATE DATABASE fixer;
\c fixer;

DROP USER IF EXISTS pgtest;
CREATE USER pgtest WITH CREATEDB LOGIN PASSWORD '123456';

CREATE TABLE users (
    id                  SERIAL NOT NULL,
    nicky               VARCHAR(64) NOT NULL,
    email               VARCHAR(64) NOT NULL,
    pwd                 CHAR(32) NOT NULL,
    portrait            VARCHAR(128),
    register_date       TIMESTAMP,
    last_login_time     TIMESTAMP,

    PRIMARY KEY(id)
);

CREATE TABLE team (
    id                  SERIAL NOT NULL, 
    name                VARCHAR(64) NOT NULL,
    leader_id           INT NOT NULL,
    goal                VARCHAR(1024),
    created_date        TIMESTAMP,
    bug_table           CHAR(32),
    bug_table_status    INT,
    status              INT,
    logo                VARCHAR(128),

    PRIMARY KEY(id)
);

CREATE TABLE user_team (
    user_id             INT NOT NULL,
    team_id             INT NOT NULL,
    join_date           TIMESTAMP NOT NULL,

    PRIMARY KEY(user_id, team_id)
);

CREATE TABLE bugs (
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

CREATE TABLE buglog (
    bug_id          INT NOT NULL,
    who             INT NOT NULL,
    action_type     INT NOT NULL,
    action_time     TIMESTAMP,
    action          VARCHAR(4096) 
);

-- index of tables
CREATE UNIQUE INDEX idx_users_nicky ON users(nicky);
CREATE UNIQUE INDEX idx_team_name ON team(name);
CREATE INDEX idx_team_leader ON team(leader_id);
CREATE INDEX idx_bugs_createdby ON bugs(created_by);
CREATE INDEX idx_bugs_priority ON bugs(priority);
CREATE INDEX idx_bugs_handler ON bugs(current_handler);
CREATE INDEX idx_bugs_status ON bugs(status);
CREATE INDEX idx_bugs_created_time ON bugs(created_time);
CREATE INDEX idx_bugs_last_update ON bugs(last_update);
CREATE INDEX idx_buglog_bugid ON buglog(bug_id);

ALTER TABLE users OWNER TO pgtest;
ALTER TABLE team OWNER TO pgtest;
ALTER TABLE user_team OWNER TO pgtest;
ALTER TABLE bugs OWNER TO pgtest;
ALTER TABLE buglog OWNER TO pgtest;

ALTER SEQUENCE users_id_seq OWNER TO pgtest;
ALTER SEQUENCE team_id_seq OWNER TO pgtest;

---------------------------------------------------------------------------------

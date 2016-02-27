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

-- index of tables
CREATE UNIQUE INDEX idx_users_nicky ON users(nicky);
CREATE UNIQUE INDEX idx_team_name ON team(name);
CREATE INDEX idx_team_leader ON team(leader_id);

ALTER TABLE users OWNER TO pgtest;
ALTER TABLE team OWNER TO pgtest;
ALTER TABLE user_team OWNER TO pgtest;

CREATE FUNCTION get_nicky(user_id INT) RETURNS VARCHAR AS $$
DECLARE
    name VARCHAR;
BEGIN
    SELECT nicky FROM users WHERE id = user_id INTO name;
    RETURN name;
END;
$$ LANGUAGE PLPGSQL;

---------------------------------------------------------------------------------

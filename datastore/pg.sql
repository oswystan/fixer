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

DROP USER IF EXISTS fixer;
CREATE USER fixer WITH CREATEDB LOGIN PASSWORD 'fixer';

CREATE TABLE users (
    id                  SERIAL NOT NULL,
    nicky               VARCHAR(64) NOT NULL,
    email               VARCHAR(64) NOT NULL,
    pwd                 CHAR(32) NOT NULL,
    portrait            VARCHAR(128),
    register_date       TIMESTAMP WITH TIME ZONE,
    last_login_time     TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id)
);

CREATE TABLE team (
    id                  SERIAL NOT NULL, 
    name                VARCHAR(64) NOT NULL,
    leader_id           INT NOT NULL,
    goal                VARCHAR(1024),
    created_date        TIMESTAMP WITH TIME ZONE,
    bug_table           CHAR(40),
    bug_table_status    INT,
    status              INT,
    logo                VARCHAR(128),

    PRIMARY KEY(id)
);

CREATE TABLE user_team (
    user_id             INT NOT NULL,
    team_id             INT NOT NULL,
    join_date           TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY(user_id, team_id)
);

-- index of tables
CREATE UNIQUE INDEX idx_users_nicky ON users(nicky);
CREATE UNIQUE INDEX idx_team_name ON team(name);
CREATE INDEX idx_team_leader ON team(leader_id);

ALTER TABLE users OWNER TO fixer;
ALTER TABLE team OWNER TO fixer;
ALTER TABLE user_team OWNER TO fixer;

-- FUNCTIONS

CREATE FUNCTION get_nicky(user_id INT) RETURNS VARCHAR AS $$
DECLARE
    name VARCHAR;
BEGIN
    SELECT nicky FROM users WHERE id = user_id INTO name;
    IF NOT FOUND THEN
        name = '';
    END IF;
    RETURN name;
END;
$$ LANGUAGE PLPGSQL;

CREATE FUNCTION get_team_name(team_id INT) RETURNS VARCHAR AS $$
DECLARE
    name VARCHAR;
BEGIN
    SELECT name FROM team WHERE id = team_id INTO name;
    IF NOT FOUND THEN
        name = '';
    END IF;
    RETURN name;
END;
$$ LANGUAGE PLPGSQL;

---------------------------------------------------------------------------------

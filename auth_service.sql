CREATE DATABASE auth_microservice_mysql_db;

use auth_microservice_mysql_db;

CREATE TABLE session (
    id         VARCHAR(255),
    user_id    VARCHAR(50),
    token      VARCHAR(50),
    created_at DATETIME,
    expires_at DATETIME
);

CREATE TABLE user (
    id         VARCHAR(255),
    name       VARCHAR(50),
    email      VARCHAR(50),
    password   VARCHAR(50),
    created_at DATETIME
);

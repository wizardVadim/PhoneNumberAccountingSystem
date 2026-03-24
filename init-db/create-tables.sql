CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS user_role (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "user" (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_id BIGINT NOT NULL REFERENCES user_role(id)
);

CREATE TABLE IF NOT EXISTS phone_number_type (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    type_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS physical_person (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    city VARCHAR(255) NOT NULL,
    person_address VARCHAR(255),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255),
    born_year INTEGER
);

CREATE TABLE IF NOT EXISTS phone_number (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    phone_number_value VARCHAR(20) NOT NULL UNIQUE,
    person_id BIGINT NOT NULL REFERENCES physical_person(id),
    phone_number_type_id BIGINT NOT NULL REFERENCES phone_number_type(id),
    comment VARCHAR(255)
);
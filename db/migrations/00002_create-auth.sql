-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE auth (
    id serial NOT NULL,
    username varchar(40) NOT NULL,
    password varchar(60) NOT NULL,
    updated_at timestamp with time zone DEFAULT now(),
    created_at timestamp with time zone DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE auth;
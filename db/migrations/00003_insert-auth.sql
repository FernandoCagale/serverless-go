-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO auth (username, password, updated_at, created_at) VALUES ('Golang', '$2y$12$Q9UO22/89NHQ0Fkgqn.qF.VvUCAXkW1fUomMZTLvIdOvZ/.UqZDPC', now(), now());

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE FROM auth;
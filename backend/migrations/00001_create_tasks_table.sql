-- +goose Up
CREATE TABLE tasks (
    id         bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title      text NOT NULL CHECK (btrim(title) <> ''),
    done       boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE tasks;
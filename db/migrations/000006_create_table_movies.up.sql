CREATE TABLE movies (
    id serial NOT NULL PRIMARY KEY,
    title varchar(100) NOT NULL UNIQUE,
    watched boolean NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

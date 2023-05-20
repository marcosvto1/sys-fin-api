CREATE TABLE categories (
    id serial NOT NULL PRIMARY KEY,
    name varchar(150) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);
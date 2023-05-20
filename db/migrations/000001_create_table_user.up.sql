CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    email character varying(255) NULL,
    name character varying(255) NULL,
    password character varying(255) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

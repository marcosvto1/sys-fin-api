CREATE TABLE subscriptions (
    id serial NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL,
    price decimal(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

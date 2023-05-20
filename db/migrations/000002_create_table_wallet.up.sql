CREATE TABLE wallets (
    id serial NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL,
    amount decimal(10, 2) DEFAULT 0,
    user_id serial NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
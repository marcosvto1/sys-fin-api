CREATE TABLE transactions (
    id serial NOT NULL PRIMARY KEY,
    short_description varchar(100) NOT NULL,
    description TEXT NULL,
    amount decimal(10, 2) NOT NULL,
    category_id serial NOT NULL,
    wallet_id serial NOT NULL,
    transaction_at DATE NOT NULL,
    transaction_type VARCHAR(20),
    paid BOOLEAN NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,

   FOREIGN KEY (wallet_id) REFERENCES wallets(id),
   FOREIGN KEY (category_id) REFERENCES categories(id)
);
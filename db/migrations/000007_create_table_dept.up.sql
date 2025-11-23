CREATE TABLE IF NOT EXISTS debts (
    id serial NOT NULL PRIMARY KEY,
    category_id integer NOT NULL,
	amount     decimal(10, 2) NOT NULL, 
    date_debit date NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX unique_debt_per_month ON debts (category_id, EXTRACT(YEAR FROM date_debit), EXTRACT(MONTH FROM date_debit));
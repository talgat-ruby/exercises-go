CREATE TYPE category AS ENUM ('groceries', 'leisure', 'electronics', 'utilities', 'clothing', 'health');
CREATE TABLE IF NOT EXISTS transactions (
    id_transaction SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    id_budget INT NOT NULL,
    expense DECIMAL(10, 2) NOT NULL DEFAULT 0,
    comment TEXT,
    expense_category category NOT NULL DEFAULT 'others',
    created TIMESTAMPTZ NOT NULL DEFAULT NOW() ,
    FOREIGN KEY (id_user) REFERENCES users_(id_user),
    FOREIGN KEY (id_budget) REFERENCES budget(id_budget)
);
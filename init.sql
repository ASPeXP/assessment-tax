-- Creation of product table

CREATE TABLE IF NOT EXISTS personal_deduct (
	id SERIAL PRIMARY KEY,
	type VARCHAR(500) NOT NULL,
	amount DECIMAL(10, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- INSERT INTO personal_deduct (id, type, amount, created_at) VALUES
-- (1, 'donation',60000.00,  transaction_timestamp()),
-- (2, 'k-receipt',100000.00,  transaction_timestamp());


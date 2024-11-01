CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    block_number BIGINT,
    transaction_hash VARCHAR(66) UNIQUE,
    sender_address VARCHAR(42),
    recipient_address VARCHAR(42),
    value NUMERIC,
    gas_price NUMERIC,
    gas_limit BIGINT,
    nonce BIGINT,
    data TEXT,
    timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
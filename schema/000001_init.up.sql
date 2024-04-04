-- Table for users
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    username   VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL
);

-- Table for wallets (linked to users)
CREATE TABLE wallets
(
    id        SERIAL PRIMARY KEY,
    user_id   INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    balance   DECIMAL(10, 2) NOT NULL DEFAULT 0.00 -- Assuming balance is in bitcoins
);

-- Table for wallet transactions (linked to wallets)
CREATE TABLE wallet_transactions
(
    id            SERIAL PRIMARY KEY,
    wallet_id     INT NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,
    amount        DECIMAL(10, 2) NOT NULL, -- Positive for deposits, negative for withdrawals
    transaction_type  VARCHAR(50) NOT NULL, -- e.g., 'deposit', 'withdrawal'
    timestamp     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Dropping tables in reverse order of creation due to dependencies

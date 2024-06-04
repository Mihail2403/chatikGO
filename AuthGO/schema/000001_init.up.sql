CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255)
);

-- migrate -path schema/ -database postgres://postgres:postgres@localhost:6543/auth?sslmode=disable up
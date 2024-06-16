CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    img_id VARCHAR(24),
    user_id INT
);

-- migrate -path schema/ -database postgres://postgres:postgres@db_users:5432/users?sslmode=disable up
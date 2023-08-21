CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT null,
    email VARCHAR(32),
    created_at TIMESTAMP with time zone DEFAULT CURRENT_TIMESTAMP NOT null
);
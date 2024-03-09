CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email text NOT NULL UNIQUE,
    encrypted_password text NOT NULL 
);
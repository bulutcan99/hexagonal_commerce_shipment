SET TIME ZONE 'Europe/Istanbul';

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    permission_id INTEGER NOT NULL,
    name VARCHAR(20) NOT NULL,
    surname VARCHAR(20) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    notification_radius INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE UNIQUE INDEX "email" ON "users" ("email");

ALTER TABLE users ADD CONSTRAINT fk_users_permissons FOREIGN KEY (permission_id) REFERENCES permissions (id);
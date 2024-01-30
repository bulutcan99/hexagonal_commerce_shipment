CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    entry INTEGER NOT NULL,
    add_flag BOOLEAN NOT NULL,
    admin_flag BOOLEAN NOT NULL
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_permissions (
    user_id INTEGER REFERENCES users(id),
    permission_id INTEGER REFERENCES permissions(id),
    PRIMARY KEY (user_id, permission_id)
);

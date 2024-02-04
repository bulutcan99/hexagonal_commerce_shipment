SET TIME ZONE 'Europe/Istanbul';

CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    entry INTEGER NOT NULL,
    add_flag BOOLEAN NOT NULL,
    remove_flag BOOLEAN NOT NULL,
    admin_flag BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


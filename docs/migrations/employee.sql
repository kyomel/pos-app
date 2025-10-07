CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    public_id TEXT NOT NULL,
    name TEXT NOT NULL,
    profile TEXT NOT NULL,
    auth_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
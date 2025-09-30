CREATE TABLE IF NOT EXISTS auth (
     id SERIAL PRIMARY KEY,
     public_id TEXT NOT NULL,
     email TEXT NOT NULL,
     password TEXT NOT NULL,
     is_active BOOLEAN NOT NULL,
     role TEXT NOT NULL,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW()
);
    
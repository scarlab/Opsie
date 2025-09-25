-- =====================
-- Extensions
-- =====================
CREATE EXTENSION IF NOT EXISTS "pgcrypto";


-- =====================
-- Auto-updated timestamps
-- =====================
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- =====================
-- Users
-- =====================
CREATE TYPE user_role AS ENUM ('owner', 'admin', 'staff');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    display_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    avatar TEXT,
    role user_role DEFAULT 'staff',
    preference JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER set_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- =====================
-- Sessions
-- =====================
CREATE TABLE sessions (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    key TEXT UNIQUE NOT NULL,
    ip TEXT,
    os TEXT,
    device TEXT,
    browser TEXT,
    is_valid BOOLEAN DEFAULT true,
    expiry TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


-- =====================
-- Nodes
-- =====================
CREATE TABLE nodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    hostname TEXT,
    ip_address TEXT,
    os TEXT,
    kernel TEXT,
    arch TEXT,
    cores INT,
    threads INT,
    memory BIGINT,
    online BOOLEAN DEFAULT false,
    token TEXT,
    verified BOOLEAN DEFAULT false,
    last_seen TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER set_nodes_updated_at
BEFORE UPDATE ON nodes
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- =====================
-- Drop triggers
-- =====================
DROP TRIGGER IF EXISTS set_users_updated_at ON users;
DROP TRIGGER IF EXISTS set_nodes_updated_at ON nodes;

-- =====================
-- Drop tables
-- =====================
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS nodes CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- =====================
-- Drop types
-- =====================
DROP TYPE IF EXISTS user_role;

-- =====================
-- Drop functions
-- =====================
DROP FUNCTION IF EXISTS set_updated_at();
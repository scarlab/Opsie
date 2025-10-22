-- -----------------------------------------------------------------------
-- Drop triggers first
-- -----------------------------------------------------------------------
DROP TRIGGER IF EXISTS set_users_updated_at ON users;
DROP TRIGGER IF EXISTS set_nodes_updated_at ON nodes;
DROP TRIGGER IF EXISTS set_teams_updated_at ON teams;
DROP TRIGGER IF EXISTS set_roles_updated_at ON roles;
DROP TRIGGER IF EXISTS set_projects_updated_at ON projects;
DROP TRIGGER IF EXISTS set_resources_updated_at ON resources;

-- -----------------------------------------------------------------------
-- Drop tables in reverse dependency order
-- -----------------------------------------------------------------------
DROP TABLE IF EXISTS resource_nodes;
DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS user_teams;
DROP TABLE IF EXISTS nodes;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS users;

-- -----------------------------------------------------------------------
-- Drop enums
-- -----------------------------------------------------------------------
DROP TYPE IF EXISTS resource_status;
DROP TYPE IF EXISTS user_system_role;

-- -----------------------------------------------------------------------
-- Drop trigger function
-- -----------------------------------------------------------------------
DROP FUNCTION IF EXISTS set_updated_at();

-- -----------------------------------------------------------------------
-- Timestamp trigger
-- -----------------------------------------------------------------------
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- -----------------------------------------------------------------------
-- Enum types
-- -----------------------------------------------------------------------
CREATE TYPE user_system_role AS ENUM ('owner', 'admin', 'staff');

CREATE TYPE resource_status AS ENUM (
    'stopped',
    'starting',
    'running',
    'restarting',
    'degraded',
    'failed'
);

-- -----------------------------------------------------------------------
-- Users
-- -----------------------------------------------------------------------
CREATE TABLE users (
    id              BIGINT PRIMARY KEY,
    display_name    TEXT NOT NULL,
    email           TEXT UNIQUE NOT NULL,
    password        TEXT NOT NULL,
    avatar          TEXT,
    system_role     user_system_role DEFAULT 'staff',
    preference      JSONB DEFAULT '{}',
    is_active       BOOLEAN DEFAULT true,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_users_email ON users(email);

CREATE TRIGGER set_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- -----------------------------------------------------------------------
-- Sessions
-- -----------------------------------------------------------------------
CREATE TABLE sessions (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    key             TEXT UNIQUE NOT NULL,
    ip              TEXT,
    os              TEXT,
    device          TEXT,
    browser         TEXT,
    is_valid        BOOLEAN DEFAULT true,
    expiry          TIMESTAMP WITH TIME ZONE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_sessions_user_id ON sessions(user_id);


-- -----------------------------------------------------------------------
-- Nodes
-- -----------------------------------------------------------------------
CREATE TABLE nodes (
    id              BIGINT PRIMARY KEY,
    name            TEXT NOT NULL,
    hostname        TEXT,
    ip_address      TEXT,
    os              TEXT,
    kernel          TEXT,
    arch            TEXT,
    cores           INT,
    threads         INT,
    memory          BIGINT,
    online          BOOLEAN DEFAULT false,
    token           TEXT,
    verified        BOOLEAN DEFAULT false,
    last_seen       TIMESTAMP WITH TIME ZONE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER set_nodes_updated_at
BEFORE UPDATE ON nodes
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- -----------------------------------------------------------------------
-- Organizations
-- -----------------------------------------------------------------------
CREATE TABLE organizations (
    id              BIGINT PRIMARY KEY,
    name            TEXT NOT NULL,
    description     TEXT,
    logo            TEXT,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT now(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER set_organizations_updated_at
BEFORE UPDATE ON organizations
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- -----------------------------------------------------------------------
-- User <-> Organizations
-- -----------------------------------------------------------------------
CREATE TABLE user_organizations (
    id                  BIGSERIAL PRIMARY KEY,
    user_id             BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_id     BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    invited_by          BIGINT REFERENCES users(id) ON DELETE SET NULL,
    joined_at           TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE UNIQUE INDEX idx_user_org_unique ON user_organizations(user_id, organization_id);
CREATE INDEX idx_user_org_user_id ON user_organizations(user_id);
CREATE INDEX idx_user_org_org_id ON user_organizations(organization_id);


-- -----------------------------------------------------------------------
-- Roles
-- -----------------------------------------------------------------------
CREATE TABLE roles (
    id                BIGINT PRIMARY KEY,
    organization_id   BIGINT REFERENCES organizations(id) ON DELETE CASCADE,
    title             TEXT NOT NULL,
    color             TEXT,
    is_default        BOOLEAN DEFAULT false,
    is_active         BOOLEAN DEFAULT true,
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT now(),
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_roles_org_id ON roles(organization_id);

CREATE TRIGGER set_roles_updated_at
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- -----------------------------------------------------------------------
-- Permissions
-- -----------------------------------------------------------------------
CREATE TABLE permissions (
    id              BIGSERIAL PRIMARY KEY,
    code            TEXT UNIQUE NOT NULL,
    title           TEXT NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now()
);


-- -----------------------------------------------------------------------
-- Role <-> Permissions
-- -----------------------------------------------------------------------
CREATE TABLE role_permissions (
    id              BIGSERIAL PRIMARY KEY,
    role_id         BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id   BIGINT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE (role_id, permission_id)
);

CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);


-- -----------------------------------------------------------------------
-- User <-> Roles
-- -----------------------------------------------------------------------
CREATE TABLE user_roles (
    id                BIGSERIAL PRIMARY KEY,
    user_id           BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id           BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    organization_id   BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    assigned_by       BIGINT REFERENCES users(id) ON DELETE SET NULL,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE (user_id, role_id, organization_id)
);

CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_org_id ON user_roles(organization_id);


-- -----------------------------------------------------------------------
-- Projects
-- -----------------------------------------------------------------------
CREATE TABLE projects (
    id                BIGINT PRIMARY KEY,
    organization_id   BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name              TEXT NOT NULL,
    description       TEXT,
    status            TEXT,
    is_archived       BOOLEAN DEFAULT false,
    archived_at       TIMESTAMP,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_projects_org_id ON projects(organization_id);

CREATE TRIGGER set_projects_updated_at
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- -----------------------------------------------------------------------
-- Resources
-- -----------------------------------------------------------------------
CREATE TABLE resources (
    id                BIGINT PRIMARY KEY,
    organization_id   BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id        BIGINT REFERENCES projects(id) ON DELETE CASCADE,
    name              TEXT NOT NULL,
    description       TEXT,
    type              TEXT,
    ports             JSONB,
    env               JSONB,
    replicas          INT DEFAULT 1,
    status            resource_status DEFAULT 'stopped',
    is_archived       BOOLEAN DEFAULT false,
    archived_at       TIMESTAMP,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_resources_org_id ON resources(organization_id);
CREATE INDEX idx_resources_project_id ON resources(project_id);

CREATE TRIGGER set_resources_updated_at
BEFORE UPDATE ON resources
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- -----------------------------------------------------------------------
-- Resource <-> Nodes (deployment instances)
-- -----------------------------------------------------------------------
CREATE TABLE resource_nodes (
    id              BIGSERIAL PRIMARY KEY,
    resource_id     BIGINT NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    node_id         BIGINT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    status          resource_status DEFAULT 'stopped',
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE (resource_id, node_id)
);

CREATE INDEX idx_resource_nodes_resource_id ON resource_nodes(resource_id);
CREATE INDEX idx_resource_nodes_node_id ON resource_nodes(node_id);

CREATE TABLE IF NOT EXISTS platform_users (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT,
    password_hash TEXT NOT NULL,
    role TEXT CHECK (role IN ('admin', 'operator', 'pending')) NOT NULL DEFAULT 'operator',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 2. Clinics
CREATE TABLE IF NOT EXISTS clinics (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    db_name TEXT NOT NULL,
    connection_uri TEXT,
    created_by UUID REFERENCES platform_users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- 3. Admin Role Elevation Keys (One-time keys for privileged role assignment)
CREATE TABLE IF NOT EXISTS admin_keys (
    key TEXT PRIMARY KEY,
    user_email TEXT NOT NULL,
    description TEXT,
    used BOOLEAN DEFAULT FALSE,
    used_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 4. Audit Logs (optional but useful)
CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES platform_users(id),
    action TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 5. Migration Tracker (if not using golang-migrate, keep track of manual SQL versions)
CREATE TABLE IF NOT EXISTS db_schema_migrations (
    version TEXT PRIMARY KEY,
    applied_at TIMESTAMP DEFAULT NOW()
);
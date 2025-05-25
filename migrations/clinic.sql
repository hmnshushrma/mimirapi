-- This is for rx_<clinic>_<uuid> databases

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT CHECK (role IN ('doctor', 'receptionist', 'helper', 'admin')) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS patients (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT,
    phone TEXT,
    dob DATE,
    gender TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS appointments (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    doctor_id UUID REFERENCES users(id),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    status TEXT CHECK (status IN ('scheduled', 'in_progress', 'completed', 'cancelled')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS visits (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    doctor_id UUID REFERENCES users(id),
    notes TEXT,
    diagnosis TEXT,
    prescription TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS documents (
    id UUID PRIMARY KEY,
    patient_id UUID REFERENCES patients(id),
    uploaded_by UUID REFERENCES users(id),
    file_url TEXT,
    file_type TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

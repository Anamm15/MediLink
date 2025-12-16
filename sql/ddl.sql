-- Gunakan ENUM untuk tipe data yang pilihannya terbatas, lebih efisien & aman
CREATE TYPE user_role AS ENUM ('patient', 'doctor', 'admin', 'staff', 'clinic');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'suspended', 'banned');
CREATE TYPE gender_enum AS ENUM ('male', 'female');
CREATE TYPE appointment_status AS ENUM ('pending', 'confirmed', 'completed', 'canceled', 'expired');
CREATE TYPE appointment_type AS ENUM ('video_call', 'chat', 'onsite');
CREATE TYPE payment_status AS ENUM ('unpaid', 'paid', 'failed', 'refunded');
CREATE TYPE schedule_day AS ENUM ('monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday');

-- --- 1. Core User ---

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    address TEXT,
    role user_role NOT NULL,
    status user_status NOT NULL DEFAULT 'active',
    birth_place VARCHAR(100),
    birth_date DATE,
    gender gender_enum,
    is_verified BOOLEAN DEFAULT FALSE,
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE clinics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE,
    type VARCHAR(100), 
    address TEXT NOT NULL,
    city VARCHAR(100),
    province VARCHAR(100),
    postal_code VARCHAR(10),
    latitude NUMERIC(9, 6),
    longitude NUMERIC(9, 6),
    phone_number VARCHAR(20),
    insurance_partners JSONB, -- ["BPJS", "Prudential", "AXA"]
    facilities JSONB,         -- {"icu":true,"lab":true,"radiology":false}
    email VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    accreditation VARCHAR(100),
    established_date DATE,
    opening_time JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE doctors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    clinic_id UUID REFERENCES clinics(id) ON DELETE SET NULL,
    specialization VARCHAR(100) NOT NULL,
    license_number VARCHAR(100) UNIQUE NOT NULL,
    consultation_fee NUMERIC(12, 2) DEFAULT 0,
    experience JSONB, -- [ { "place": "RS Harapan Bunda", "years": "2018-2020" } ]
    education JSONB, -- [ { "institution": "Universitas Indonesia", "degree": "Sp.A" } ]
    is_active BOOLEAN DEFAULT TRUE,
    rating NUMERIC(3, 2) DEFAULT 0,
    total_reviews INT DEFAULT 0,
    available_for_telemedicine BOOLEAN DEFAULT FALSE,
    bio TEXT
);

CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_number VARCHAR(50) UNIQUE, -- Primary NIK
    blood_type VARCHAR(5),
    weight_kg NUMERIC(5, 2),
    height_cm NUMERIC(5, 2),
    allergies TEXT,
    chronic_diseases TEXT,
    emergency_contact TEXT,
    insurance_provider VARCHAR(100),
    insurance_number VARCHAR(100),
    occupation VARCHAR(100),
);


-- --- 2. Transactional & Activity Tables ---

CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    schedule_start_time TIMESTAMPTZ NOT NULL,
    schedule_end_time TIMESTAMPTZ NOT NULL,
    duration_minutes INT,
    status appointment_status NOT NULL DEFAULT 'pending',
    type appointment_type NOT NULL,
    canceled_reason TEXT,
    complaint TEXT,
    location TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE doctor_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    doctor_id UUID NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
    day_of_week schedule_day NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    type appointment_type NOT NULL,
    location TEXT,
    max_appointments INT,
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(doctor_id, day_of_week, start_time, end_time) -- Mencegah jadwal duplikat
);

CREATE TABLE medicines (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    dosage VARCHAR(100),
    price NUMERIC(12, 2) NOT NULL,
    stock INT DEFAULT 0,
    requires_prescription BOOLEAN DEFAULT TRUE
);

CREATE TABLE prescriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    clinic_id UUID REFERENCES clinics(id),
    notes TEXT,
    is_redeemed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Join table for Prescriptions and Medicines (Many-to-Many)
CREATE TABLE prescription_medicines (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prescription_id UUID NOT NULL REFERENCES prescriptions(id) ON DELETE CASCADE,
    medicine_id UUID NOT NULL REFERENCES medicines(id),
    instructions TEXT NOT NULL, 
    quantity INT NOT NULL,
    UNIQUE(prescription_id, medicine_id)
);

-- --- 3. Supporting & Other Tables ---

CREATE TABLE medical_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    appointment_id UUID REFERENCES appointments(id),
    diagnosis TEXT,
    treatment_plan TEXT,
    next_appointment_date DATE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE billings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    appointment_id UUID UNIQUE REFERENCES appointments(id),
    amount NUMERIC(12, 2) NOT NULL,
    status payment_status NOT NULL DEFAULT 'unpaid',
    due_date DATE,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    billing_id UUID NOT NULL REFERENCES billings(id),
    payment_method VARCHAR(50),
    amount NUMERIC(12, 2) NOT NULL,
    status payment_status NOT NULL,
    payment_gateway_order_id VARCHAR(255),
    payment_gateway_transaction_id VARCHAR(255),
    payment_url TEXT,
    va_number VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID REFERENCES doctors(id), -- Can be null if reviewing a clinic
    clinic_id UUID REFERENCES clinics(id), -- Can be null if reviewing a doctor
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT chk_review_target CHECK (doctor_id IS NOT NULL OR clinic_id IS NOT NULL) -- A review must be for a doctor OR a clinic
);

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Penerima
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    type VARCHAR(100), -- 'APPOINTMENT_REMINDER', 'PRESCRIPTION_READY'
    priority VARCHAR(20) DEFAULT 'normal',
    read_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Polymorphic Table for Files (CENTRAL)
CREATE TABLE files (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    file_name VARCHAR(255) NOT NULL,
    url TEXT NOT NULL, -- URL dari cloud storage (S3, GCS, dll)
    fileable_id UUID NOT NULL,
    fileable_type VARCHAR(100) NOT NULL, -- e.g., 'User', 'Medicine', 'MedicalRecord'
    mime_type VARCHAR(100),
    size_bytes INT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Create indexes on frequently queried columns
CREATE INDEX idx_appointments_patient_id ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor_id ON appointments(doctor_id);
CREATE INDEX idx_files_fileable ON files(fileable_id, fileable_type);
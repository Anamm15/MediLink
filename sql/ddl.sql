-- =============================================
-- SECTION 0: SETUP & EXTENSIONS
-- =============================================
-- We need pgcrypto for UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =============================================
-- SECTION 1: ENUMS (Type Safety)
-- Using Enums prevents invalid data entry (e.g., preventing typos)
-- =============================================
CREATE TYPE user_role AS ENUM ('patient', 'doctor', 'admin', 'pharmacist', 'nurse', 'super_admin');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'suspended', 'banned');
CREATE TYPE gender_enum AS ENUM ('male', 'female');
CREATE TYPE appointment_status AS ENUM ('pending', 'confirmed', 'in_progress', 'completed', 'canceled', 'expired');
CREATE TYPE appointment_type AS ENUM ('video_call', 'onsite', 'chat');
CREATE TYPE payment_status AS ENUM ('unpaid', 'paid', 'failed', 'refunded');
CREATE TYPE schedule_day AS ENUM ('monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday');

-- =============================================
-- SECTION 2: IDENTITY & CORE ENTITIES
-- The foundation of the system
-- =============================================

-- 2.1 USERS
-- Base table for authentication.
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    
    -- Profile Data
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    phone_number VARCHAR(20) UNIQUE,
    
    is_email_verified BOOLEAN DEFAULT FALSE,
    status user_status NOT NULL DEFAULT 'active',
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 2.2 CLINICS
-- Represents physical or digital branches.
CREATE TABLE clinics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(100),
    coordinates POINT, -- Latitude/Longitude 
    accreditation VARCHAR(100),
    
    phone_number VARCHAR(20),
    email VARCHAR(255),
    
    is_active BOOLEAN DEFAULT TRUE,
    insurance_partners JSONB, -- Array for accepted insurances
    facilities JSONB, -- Array of strings: ["Pharmacy", "Lab"]
    opening_time JSONB, -- e.g., {"monday": "08:00-17:00"}
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 2.3 DOCTORS (Profile)
-- Note: No clinic_id here. A doctor is an entity independent of a specific clinic.
CREATE TABLE doctors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    specialization VARCHAR(100) NOT NULL, -- e.g., "Cardiologist"
    license_number VARCHAR(50) UNIQUE NOT NULL, -- STR/SIP
    
    bio TEXT,
    experience_years INT DEFAULT 0,
    education JSONB, -- Array of objects: [{"degree": "MD", "school": "UI"}]

    rating_total NUMERIC(3, 2) DEFAULT 0, -- Cached average rating (e.g. 4.8)
    rating_count INT DEFAULT 0,
    review_count INT DEFAULT 0,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 2.4 DOCTOR PLACEMENTS (The "Multi-Clinic" Fix)
-- Connects a Doctor to a Clinic. A doctor can exist in 5 clinics with 5 different prices.
CREATE TABLE doctor_clinic_placements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    doctor_id UUID NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
    clinic_id UUID NOT NULL REFERENCES clinics(id) ON DELETE CASCADE,
    
    -- Financials specific to this location
    consultation_fee NUMERIC(12, 2) NOT NULL DEFAULT 0,
    
    is_active BOOLEAN DEFAULT TRUE, -- Doctor might take a break from just this clinic
    created_at TIMESTAMPTZ DEFAULT NOW(),
    
    UNIQUE(doctor_id, clinic_id) -- Prevent duplicate assignment
);

-- 2.5 PATIENTS
-- Medical profile separated from User auth data.
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    identity_number VARCHAR(50) UNIQUE, -- NIK or Country ID
    birth_date DATE,
    gender gender_enum,
    blood_type VARCHAR(5),
    
    height_cm NUMERIC(5, 2),
    weight_kg NUMERIC(5, 2),
    allergies TEXT, 
    
    emergency_contact VARCHAR(20),
    history_chronic_diseases TEXT,

    insurance_provider VARCHAR(100),
    insurance_number VARCHAR(100),
    occupation VARCHAR(100),
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- SECTION 3: INVENTORY & PHARMACY
-- Scalable "Catalog vs Stock" approach
-- =============================================

-- 3.1 MEDICINE CATALOG (Global Definition)
-- What the medicine IS (e.g., "Panadol 500mg"). Shared across all clinics.
CREATE TABLE medicines (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    generic_name VARCHAR(255), -- e.g., "Paracetamol"
    category VARCHAR(100), -- "Antibiotic", "Analgesic"
    description TEXT,
    manufacturer VARCHAR(100),
    
    is_prescription_required BOOLEAN DEFAULT FALSE,
    base_price NUMERIC(12, 2),
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 3.2 CLINIC INVENTORY (Local Stock)
-- What a specific clinic HAS (e.g., "Clinic A has 50 boxes of Batch X").
CREATE TABLE clinic_inventories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    clinic_id UUID NOT NULL REFERENCES clinics(id) ON DELETE CASCADE,
    medicine_id UUID NOT NULL REFERENCES medicines(id),
    
    current_stock INT NOT NULL DEFAULT 0 CHECK (current_stock >= 0),
    low_stock_threshold INT DEFAULT 10, -- trigger alert
    price NUMERIC(12, 2) NOT NULL, 

    batch_number VARCHAR(100),
    expiry_date DATE,
    
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- SECTION 4: SCHEDULING & TRANSACTIONS
-- The core business logic
-- =============================================

-- 4.1 SCHEDULES
-- Defined templates (e.g., "Dr. Budi works Mondays 09:00-12:00 at Clinic A")
CREATE TABLE doctor_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    clinic_id UUID NOT NULL REFERENCES clinics(id),
    
    day_of_week schedule_day NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    
    -- Helps frontend generate slots (e.g., 09:00, 09:15, 09:30)
    slot_duration_minutes INT DEFAULT 15,
    max_quota INT,
    
    is_active BOOLEAN DEFAULT TRUE,
    
    UNIQUE(doctor_id, clinic_id, day_of_week, start_time, end_time)
);

-- 4.2 APPOINTMENTS
CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    clinic_id UUID NOT NULL REFERENCES clinics(id),
    
    appointment_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    
    status appointment_status NOT NULL DEFAULT 'pending_payment',
    type appointment_type NOT NULL,
    
    -- === FINANCIAL SNAPSHOTS (CRITICAL) ===
    -- We record the fee *at the moment of booking*. 
    -- If doctor changes price later, this historical record remains accurate.
    consultation_fee_snapshot NUMERIC(12, 2) NOT NULL,
    
    -- === EXECUTION DETAILS ===
    queue_number INT, -- For onsite
    meeting_link TEXT, -- For video call
    
    -- === MEDICAL CONTEXT ===
    symptom_complaint TEXT, -- Patient's reason for visiting
    doctor_notes TEXT, -- Private notes for doctor
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- SECTION 5: MEDICAL RECORDS & OUTCOMES
-- =============================================

-- 5.1 MEDICAL RECORDS (EMR)
-- Follows SOAP Standard (Subjective, Objective, Assessment, Plan)
CREATE TABLE medical_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    appointment_id UUID UNIQUE REFERENCES appointments(id), -- 1-to-1 with appointment
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    
    -- SOAP Data
    subjective TEXT, -- "Patient complains of headache..."
    objective TEXT,  -- "Blood pressure 120/80..."
    assessment TEXT, -- "Diagnosis: Migraine"
    plan TEXT,       -- "Rest, Prescription X..."
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 5.2 PRESCRIPTIONS
CREATE TABLE prescriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    medical_record_id UUID NOT NULL REFERENCES medical_records(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    patient_id UUID NOT NULL REFERENCES patients(id),
    
    notes TEXT, -- "Take after meals"
    status BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 5.3 PRESCRIPTION ITEMS
-- Links prescription to the Catalog (not Inventory, because Inventory is decided at Checkout)
CREATE TABLE prescription_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prescription_id UUID NOT NULL REFERENCES prescriptions(id) ON DELETE CASCADE,
    medicines_id UUID NOT NULL REFERENCES medicines(id),
    
    quantity INT NOT NULL,
    dosage_instruction VARCHAR(255), -- "3x1 per day"
    notes TEXT
);

-- =============================================
-- SECTION 6: FINANCE (Billing)
-- =============================================

-- 6.1 BILLINGS (Invoice)
CREATE TABLE billings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    appointment_id UUID UNIQUE REFERENCES appointments(id),
    patient_id UUID NOT NULL REFERENCES patients(id),
    
    total_amount NUMERIC(12, 2) NOT NULL,
    status payment_status DEFAULT 'unpaid',
    issued_at TIMESTAMPTZ DEFAULT NOW(),
    paid_at TIMESTAMPTZ
);

-- 6.2 PAYMENTS (Transactions)
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    billing_id UUID NOT NULL REFERENCES billings(id),
    
    external_id VARCHAR(255), -- Payment Gateway ID (e.g. from Midtrans/Stripe)
    payment_method VARCHAR(50), -- "credit_card", "gopay", "bca_va"
    amount NUMERIC(12, 2) NOT NULL,
    status payment_status NOT NULL,
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- SECTION 7: INDEXING (Performance)
-- Add indexes on columns frequently used in WHERE, JOIN, and ORDER BY
-- =============================================

-- Users
CREATE INDEX idx_users_email ON users(email);

-- Doctors & Clinics
CREATE INDEX idx_doctor_placements_clinic ON doctor_clinic_placements(clinic_id);
CREATE INDEX idx_doctor_placements_doctor ON doctor_clinic_placements(doctor_id);

-- Appointments
CREATE INDEX idx_appointments_patient ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor_date ON appointments(doctor_id, appointment_date);
CREATE INDEX idx_appointments_status ON appointments(status);

-- Inventory
CREATE INDEX idx_inventory_clinic_catalog ON clinic_inventories(clinic_id, catalog_id);

-- Search functionality
CREATE INDEX idx_medicine_name ON medicines USING gin(to_tsvector('english', name));
CREATE INDEX idx_doctor_specialization ON doctors(specialization);
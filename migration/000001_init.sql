CREATE TABLE login_in (
    'id' SERIAL PRIMARY KEY,
    'username' VARCHAR(50) UNIQUE NOT NULL,
    'password' VARCHAR(255) NOT NULL,
    'email' VARCHAR(100) UNIQUE NOT NULL,
    'user_id' INT REFERENCES users(id), -- Foreign key referencing the users table
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS 'roles' (
    'id' SERIAL PRIMARY KEY,
    'name' VARCHAR(30) NOT NULL,
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS departments (
    'id' SERIAL PRIMARY KEY,
    'name' VARCHAR(45) NOT NULL,
    'address' VARCHAR(100) NOT NULL,
    'status' SMALLINT NOT NULL CHECK (status IN (0, 1)), -- 0: inactive, 1: active
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS countries (
    'id' SERIAL PRIMARY KEY,
    'name' VARCHAR(45) NOT NULL,
    'status' SMALLINT NOT NULL CHECK (status IN (0, 1)),
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    'id' SERIAL PRIMARY KEY,
    'role_id' INT NOT NULL REFERENCES roles(id),
    'department_id' INT DEFAULT NULL REFERENCES departments(id),
    'email' VARCHAR(100) NOT NULL UNIQUE,
    'password' VARCHAR(255) NOT NULL, -- Adjusted to use VARCHAR(255) for storing hashed passwords
    'name' VARCHAR(45) NOT NULL,
    'surname' VARCHAR(45) NOT NULL,
    'gender' VARCHAR(20) NOT NULL CHECK (gender IN ('male', 'female', 'other')), -- Ensuring gender can only be 'male', 'female', or 'other'
    'date_of_birth' DATE NOT NULL,
    'mobile' VARCHAR(15) NOT NULL,
    'country_id' INT NOT NULL REFERENCES countries(id),
    'resident_country_id' INT NOT NULL REFERENCES countries(id),
    'avatar' VARCHAR(100) DEFAULT NULL,
    'verification_status' SMALLINT DEFAULT 0 CHECK (verification_status IN (0, 1)), -- 0: unverified, 1: verified
    'status' SMALLINT NOT NULL CHECK (status IN (0, 1)), -- 0: inactive, 1: active
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS volunteer_details (
    'id' SERIAL PRIMARY KEY,
    'user_id' INT NOT NULL REFERENCES users(id),
    'department_id' INT NOT NULL REFERENCES departments(id),
    'status' SMALLINT NOT NULL CHECK (status IN (0, 1)),
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS requests (
    'id' SERIAL PRIMARY KEY,
    'user_id' INT NOT NULL REFERENCES users(id),
    'type' VARCHAR(45) NOT NULL,
    'status' SMALLINT NOT NULL CHECK (status IN (0, 1)),
    'reject_notes' VARCHAR(255) DEFAULT NULL,
    'verifier_id' INT DEFAULT NULL REFERENCES users(id),
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_identities (
    'id' SERIAL PRIMARY KEY,
    'user_id' INT NOT NULL REFERENCES users(id),
    'number' VARCHAR(30) NOT NULL,
    'type' VARCHAR(45) NOT NULL,
    'status' SMALLINT NOT NULL CHECK (status IN (0, 1)),
    'expiry_date' DATE NOT NULL,
    'place_issued' VARCHAR(100) NOT NULL,
    'created_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    'updated_at' TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

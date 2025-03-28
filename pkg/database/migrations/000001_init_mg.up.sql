-- 001_create_users_table.up.sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       phone VARCHAR(20),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 002_create_doctors_table.up.sql
CREATE TABLE doctors (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         specialty VARCHAR(100) NOT NULL,
                         working_hour_start TIMESTAMP WITH TIME ZONE,
                         working_hour_end TIMESTAMP WITH TIME ZONE
);

-- 003_create_reservations_table.up.sql
CREATE TABLE reservations (
                              id SERIAL PRIMARY KEY,
                              user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                              doctor_id INTEGER REFERENCES doctors(id) ON DELETE CASCADE,
                              reservation_time TIMESTAMP WITH TIME ZONE NOT NULL,
                              status VARCHAR(20) DEFAULT 'scheduled',
                              UNIQUE(doctor_id, reservation_time)
);
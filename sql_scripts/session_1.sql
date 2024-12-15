DROP TABLE IF EXISTS Client CASCADE;
DROP TABLE IF EXISTS Realtor CASCADE;

CREATE TABLE Client (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(255),
    first_name VARCHAR(255),
    middle_name VARCHAR(255),
    phone_number VARCHAR(20),
    email VARCHAR(255),
    CONSTRAINT contact_info_check CHECK (phone_number IS NOT NULL OR email IS NOT NULL)
);

CREATE TABLE Realtor (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255) NOT NULL,
    commission_share NUMERIC,
    CONSTRAINT commission_share_check CHECK (commission_share >= 0 AND commission_share <= 100)
);

CREATE INDEX idx_client_phone ON Client(phone_number);
CREATE INDEX idx_client_email ON Client(email);
CREATE INDEX idx_realtor_lastname ON Realtor(last_name);

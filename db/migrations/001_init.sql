-- Create roles
CREATE TABLE IF NOT EXISTS roles (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
    description TEXT
    );

-- Create users
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     company_name VARCHAR(255) NOT NULL,
    bin_iin VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role_id INT REFERENCES roles(id)
    );

-- Create messages
CREATE TABLE IF NOT EXISTS messages (
                                        id SERIAL PRIMARY KEY,
                                        sender_id INT REFERENCES users(id),
    receiver_id INT REFERENCES users(id),
    content TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL
    );

-- Create leads
CREATE TABLE IF NOT EXISTS leads (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at DATE NOT NULL,
    owner_id INT REFERENCES users(id),
    status VARCHAR(100)
    );

-- Create deals
CREATE TABLE IF NOT EXISTS deals (
                                     id SERIAL PRIMARY KEY,
                                     lead_id INT REFERENCES leads(id),
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    status VARCHAR(100),
    created_at DATE NOT NULL
    );

-- Create tasks
CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     creator_id INT REFERENCES users(id),
    assignee_id INT REFERENCES users(id),
    entity_id INT NOT NULL,
    entity_type VARCHAR(100) NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    due_date TIMESTAMP NOT NULL,
    status VARCHAR(100)
    );

-- Create documents
CREATE TABLE IF NOT EXISTS documents (
                                         id SERIAL PRIMARY KEY,
                                         deal_id INT REFERENCES deals(id),
    doc_type VARCHAR(100),
    file_path VARCHAR(255),
    status VARCHAR(100),
    signed_at TIMESTAMP
    );

-- Create sms_confirmations
CREATE TABLE IF NOT EXISTS sms_confirmations (
                                                 id SERIAL PRIMARY KEY,
                                                 document_id INT REFERENCES documents(id),
    sms_code VARCHAR(100),
    sent_at TIMESTAMP,
    confirmed BOOLEAN DEFAULT FALSE,
    confirmed_at TIMESTAMP
    );

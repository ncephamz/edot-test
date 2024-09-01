CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(50) NOT NULL,  
    phone_number VARCHAR(14) NOT NULL,      
    email VARCHAR(100) NOT NULL,       
    password TEXT NOT NULL,       
    name VARCHAR(50) NULL,
    photo_profile VARCHAR(225) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id)
);

CREATE INDEX users_idx_phone_number ON users(phone_number)
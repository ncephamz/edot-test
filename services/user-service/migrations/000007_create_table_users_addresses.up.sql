CREATE TABLE IF NOT EXISTS users_addresses (
    user_address_id VARCHAR(50) NOT NULL,  
    user_id VARCHAR(50) NOT NULL,  
    province_id VARCHAR(2) NOT NULL,      
    city_id VARCHAR(2) NOT NULL,      
    district_id VARCHAR(2) NOT NULL,      
    sub_district_id VARCHAR(2) NOT NULL,      
    zipcode VARCHAR(6) NOT NULL,       
    address TEXT NOT NULL,       
    note TEXT NULL,
    google_map VARCHAR(225) NULL,
    is_main BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_address_id)
);

CREATE INDEX user_address_user_id ON users_addresses(user_id);
CREATE INDEX user_address_province_id ON users_addresses(province_id);
CREATE INDEX user_address_city_id ON users_addresses(city_id);
CREATE INDEX user_address_district_id ON users_addresses(district_id);
CREATE INDEX user_address_sub_district_id ON users_addresses(sub_district_id);
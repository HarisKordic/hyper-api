-- Create Carbon Footprint History table
CREATE TABLE carbon_footprint_history (
    id SERIAL PRIMARY KEY,
    month DATE NOT NULL,
    amount DECIMAL(4,2) NOT NULL
);

-- Create Pollution Levels table
CREATE TABLE pollution_levels (
    id SERIAL PRIMARY KEY,
    pollutant VARCHAR(50) NOT NULL,
    level INTEGER NOT NULL,
    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Map Users table
CREATE TABLE map_users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    activity TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    icon_type VARCHAR(50) NOT NULL,
    emission_amount DECIMAL(4,1) NOT NULL
);
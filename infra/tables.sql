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
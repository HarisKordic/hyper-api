-- Insert Carbon Footprint History
INSERT INTO carbon_footprint_history (month, amount) VALUES
    ('2025-01-01', 3.2),
    ('2025-02-01', 3.0),
    ('2025-03-01', 2.8),
    ('2025-04-01', 2.7),
    ('2025-05-01', 2.5),
    ('2025-06-01', 2.4),
    ('2025-07-01', 2.4);

-- Insert Pollution Levels
INSERT INTO pollution_levels (pollutant, level) VALUES
    ('PM2.5', 35),
    ('PM10', 42),
    ('NO2', 28),
    ('SO2', 15),
    ('CO', 20);

-- Insert mock data
INSERT INTO map_users (name, avatar, latitude, longitude, activity, icon_type, emission_amount) VALUES
    ('Sarah K.', '/placeholder.svg?height=40&width=40', 43.3438, 17.8078, 'Drove to Mostar in Tesla Model 3', 'car', -2.5),
    ('Alex M.', '/placeholder.svg?height=40&width=40', 43.3538, 17.8178, 'Rode e-scooter to work', 'bike', -1.2),
    ('Mike R.', '/placeholder.svg?height=40&width=40', 43.3338, 17.8278, 'Charged EV at green station', 'zap', -3.0);
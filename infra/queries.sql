-- Get carbon footprint trend
SELECT
    to_char(month, 'Mon') as month,
    amount
FROM carbon_footprint_history
ORDER BY month;

-- Get current pollution levels
SELECT
    pollutant,
    level
FROM pollution_levels;
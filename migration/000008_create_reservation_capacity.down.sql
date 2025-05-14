-- Drop the reservations table if it exists
DROP TABLE IF EXISTS reservations;

-- Remove the added columns from reservationcapacity table
ALTER TABLE IF EXISTS reservationcapacity
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;


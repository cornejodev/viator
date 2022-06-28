GRANT ALL PRIVILEGES ON DATABASE viator TO johndoe;
DROP TABLE vehicle;
CREATE TABLE IF NOT EXISTS vehicle (
	id SERIAL PRIMARY KEY,
	type VARCHAR(255) NOT NULL,
	license_plate VARCHAR(255) NOT NULL,
	passenger_capacity INT NOT NULL,
	make VARCHAR(255) NOT NULL,
	model VARCHAR(255) NOT NULL,
	year INT NOT NULL,
	mileage INT NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
);
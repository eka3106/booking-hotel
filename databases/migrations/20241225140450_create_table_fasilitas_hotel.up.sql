CREATE TABLE fasilitas_hotel (
    fasilitas_hotel_id SERIAL PRIMARY KEY,
    fasilitas_id INT NOT NULL,
    hotel_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
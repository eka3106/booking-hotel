ALTER TABLE users
ADD COLUMN hotel_id INT;

ALTER TABLE users
ADD CONSTRAINT users_hotel_id
FOREIGN KEY (hotel_id) REFERENCES hotel(hotel_id);
CREATE TABLE status_pembayaran (
    status_pembayaran_id SERIAL PRIMARY KEY,
    status_pembayaran VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE status_booking (
    status_booking_id SERIAL PRIMARY KEY,
    status_booking_nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL     DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE status_kamar (
    status_kamar_id SERIAL PRIMARY KEY,
    status_kamar_nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tipe_kamar (
    tipe_kamar_id SERIAL PRIMARY KEY,
    tipe_kamar_nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE hotel (
    hotel_id SERIAL PRIMARY KEY,
    nama_hotel VARCHAR(255) NOT NULL,
    alamat_hotel TEXT NOT NULL,    
    telp_hotel VARCHAR(255) NOT NULL,
    email_hotel VARCHAR(255) NOT NULL,    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE fasilitas (
    fasilitas_id SERIAL PRIMARY KEY,
    jenis_fasilitas VARCHAR(255) NOT NULL,
    hotel_id INT NOT NULL,
    deskripsi TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (hotel_id) REFERENCES hotel(hotel_id)
);


CREATE TABLE hak_akses (
    hak_akses_id SERIAL PRIMARY KEY,
    hak_akses VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE kamar (
    kamar_id SERIAL PRIMARY KEY,
    hotel_id INT NOT NULL,
    nomor_kamar VARCHAR(255) NOT NULL,
    tipe_kamar_id INT NOT NULL,
    harga INT NOT NULL,
    status_kamar_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (hotel_id) REFERENCES hotel(hotel_id),
    FOREIGN KEY (status_kamar_id) REFERENCES status_kamar(status_kamar_id),
    FOREIGN KEY (tipe_kamar_id) REFERENCES tipe_kamar(tipe_kamar_id)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,    
    email_user VARCHAR(255) NOT NULL UNIQUE,
    nama VARCHAR(255) NOT NULL,
    password_user VARCHAR(255) NOT NULL,
    hak_akses_id INT NOT NULL,
    token TEXT[],
    created_at TIMESTAMP NOT NULL    DEFAULT CURRENT_TIMESTAMP ,

    FOREIGN KEY (hak_akses_id) REFERENCES hak_akses(hak_akses_id)
);

CREATE TABLE rating (
    rating_id SERIAL PRIMARY KEY,
    hotel_id INT NOT NULL,
    rating INT NOT NULL,
    user_id INT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (hotel_id) REFERENCES hotel(hotel_id)
);

CREATE TABLE booking (
    booking_id SERIAL PRIMARY KEY,
    kamar_id INT NOT NULL,
    user_id INT NOT NULL,
    hotel_id INT NOT NULL,
    tanggal_check_in DATE ,    
    tanggal_check_out DATE ,
    total_biaya INT NOT NULL,
    status_booking_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL     DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (hotel_id) REFERENCES hotel(hotel_id),
    FOREIGN KEY (kamar_id) REFERENCES kamar(kamar_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (status_booking_id) REFERENCES status_booking(status_booking_id)   
);

CREATE TABLE pembayaran (
    pembayaran_id SERIAL PRIMARY KEY,
    booking_id INT NOT NULL,
    total_pembayaran INT NOT NULL,
    tanggal_pembayaran TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_pembayaran_id INT NOT NULL,    

    FOREIGN KEY (booking_id)  REFERENCES booking(booking_id),    
    FOREIGN KEY (status_pembayaran_id) REFERENCES status_pembayaran(status_pembayaran_id)    
);
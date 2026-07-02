-- Database Schema for AWS AI Academy Card RPG MVP

CREATE TABLE IF NOT EXISTS participants (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    whatsapp TEXT,
    nama_lengkap TEXT NOT NULL,
    nama_panggilan TEXT,
    gender TEXT,
    kota TEXT,
    provinsi TEXT,
    pekerjaan TEXT,
    bio TEXT,
    rpg_class TEXT,
    rpg_stats TEXT, -- JSON string: {"coding": 80, "bugs": 75, "architecture": 70, "cloud": 85}
    image_path TEXT,
    commitment INTEGER DEFAULT 0, -- 0=belum, 1=iya farming pahala, -1=tidak
    linkedin TEXT,
    background_it TEXT,
    motivasi TEXT,
    slug TEXT UNIQUE,
    is_registered INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS provinces (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS regencies (
    code TEXT PRIMARY KEY,
    province_code TEXT NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (province_code) REFERENCES provinces(code)
);

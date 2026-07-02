package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)

type MenteeJSON struct {
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	Whatsapp    string `json:"whatsapp"`
}

func InitDB(dbPath string, schemaPath string, seedJSONPath string) (*sql.DB, error) {
	// Buat direktori folder database jika belum ada
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("gagal membuat direktori database: %v", err)
	}

	// Buka koneksi database SQLite
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka database SQLite: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %v", err)
	}

	// Jalankan schema.sql
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file schema: %v", err)
	}

	_, err = db.Exec(string(schemaBytes))
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan schema: %v", err)
	}

	// Migrasi dinamis untuk database lama agar tidak error
	_, _ = db.Exec("ALTER TABLE participants ADD COLUMN linkedin TEXT")
	_, _ = db.Exec("ALTER TABLE participants ADD COLUMN background_it TEXT")
	_, _ = db.Exec("ALTER TABLE participants ADD COLUMN motivasi TEXT")
	_, _ = db.Exec("ALTER TABLE participants ADD COLUMN slug TEXT")
	_, _ = db.Exec("ALTER TABLE participants ADD COLUMN commitment INTEGER DEFAULT 0")
	_, _ = db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_participants_slug ON participants(slug)")

	// Isi slug yang kosong untuk data lama jika ada
	if rows, err := db.Query("SELECT id, nama_lengkap FROM participants WHERE slug IS NULL OR slug = ''"); err == nil {
		type TempPart struct {
			ID   int64
			Name string
		}
		var updates []TempPart
		for rows.Next() {
			var tp TempPart
			if err := rows.Scan(&tp.ID, &tp.Name); err == nil {
				updates = append(updates, tp)
			}
		}
		rows.Close()

		if tx, err := db.Begin(); err == nil {
			for _, tp := range updates {
				baseSlug := generateSlug(tp.Name)
				slug := makeUniqueSlug(tx, baseSlug)
				_, _ = tx.Exec("UPDATE participants SET slug = ? WHERE id = ?", slug, tp.ID)
			}
			_ = tx.Commit()
		}
	}

	// Lakukan seeding jika tabel participants masih kosong
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM participants").Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("gagal memeriksa data participants: %v", err)
	}

	if count == 0 {
		fmt.Println("Melakukan seeding data awal mentee...")
		err = seedMentees(db, seedJSONPath)
		if err != nil {
			fmt.Printf("Peringatan seeding: %v\n", err)
		}
	}

	// Pastikan data fasilitator khusus selalu diperbarui/disisipkan
	err = upsertFacilitator(db)
	if err != nil {
		fmt.Printf("Peringatan inisialisasi fasilitator: %v\n", err)
	}

	// Pastikan data Sumbul Sumbil selalu ada
	if err := upsertSumbul(db); err != nil {
		fmt.Printf("Peringatan inisialisasi Sumbul: %v\n", err)
	}

	// Seed data wilayah (provinces + regencies) jika belum ada
	if err := seedWilayah(db); err != nil {
		fmt.Printf("Peringatan seed wilayah: %v\n", err)
	}

	return db, nil
}

func upsertFacilitator(db *sql.DB) error {
	email := "mochraflyherdianto@gmail.com"
	namaLengkap := "Mochammad Rafly Herdianto"
	whatsapp := "+6285155070929"
	isRegistered := 0

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM participants WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err = db.Exec(`
			UPDATE participants 
			SET nama_lengkap = ?, whatsapp = ?, is_registered = ?,
			nama_panggilan = NULL, gender = NULL, kota = NULL, provinsi = NULL, pekerjaan = NULL, bio = NULL, rpg_class = NULL, rpg_stats = NULL, linkedin = NULL, motivasi = NULL, slug = NULL
			WHERE email = ?
		`, namaLengkap, whatsapp, isRegistered, email)
	} else {
		_, err = db.Exec(`
			INSERT INTO participants(email, whatsapp, nama_lengkap, is_registered)
			VALUES(?, ?, ?, ?)
		`, email, whatsapp, namaLengkap, isRegistered)
	}
	return err
}

func upsertSumbul(db *sql.DB) error {
	email := "sumbulsumbil01@gmail.com"
	namaLengkap := "Sumbul Sumbil"
	whatsapp := "+6281234567890"

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM participants WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		slug := "sumbul-sumbil"
		// Cek slug konflik
		var slugExists bool
		_ = db.QueryRow("SELECT EXISTS(SELECT 1 FROM participants WHERE slug = ?)", slug).Scan(&slugExists)
		if slugExists {
			slug = "sumbul-sumbil-1"
		}
		_, err = db.Exec(`
			INSERT INTO participants(email, whatsapp, nama_lengkap, slug, is_registered)
			VALUES(?, ?, ?, ?, 0)
		`, email, whatsapp, namaLengkap, slug)
	}
	return err
}

func seedMentees(db *sql.DB, seedJSONPath string) error {
	byteValue, err := os.ReadFile(seedJSONPath)
	if err != nil {
		return fmt.Errorf("gagal membaca file mentee.json: %v", err)
	}

	var mentees []MenteeJSON
	if err := json.Unmarshal(byteValue, &mentees); err != nil {
		return fmt.Errorf("gagal unmarshal mentee.json: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO participants(email, whatsapp, nama_lengkap, slug) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, mentee := range mentees {
		baseSlug := generateSlug(mentee.NamaLengkap)
		slug := makeUniqueSlug(tx, baseSlug)
		_, err = stmt.Exec(mentee.Email, mentee.Whatsapp, mentee.NamaLengkap, slug)
		if err != nil {
			// Lewati jika email/slug duplikat
			continue
		}
	}

	return tx.Commit()
}

func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")

	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func makeUniqueSlug(tx *sql.Tx, baseSlug string) string {
	slug := baseSlug
	counter := 1
	for {
		var exists bool
		err := tx.QueryRow("SELECT EXISTS(SELECT 1 FROM participants WHERE slug = ?)", slug).Scan(&exists)
		if err != nil || !exists {
			break
		}
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}
	return slug
}

// --- Wilayah Seeder ---

type wilayahItem struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code,omitempty"`
}

type wilayahLocalFile struct {
	Provinces []wilayahItem `json:"provinces"`
	Regencies []wilayahItem `json:"regencies"`
}

// wilayahJSONPaths adalah daftar path kandidat untuk wilayah.json
var wilayahJSONPaths = []string{
	"./data/wilayah.json",
	"../data/wilayah.json",
	"./wilayah.json",
}

func loadWilayahFromFile() (*wilayahLocalFile, string, error) {
	for _, p := range wilayahJSONPaths {
		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		var result wilayahLocalFile
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, p, fmt.Errorf("gagal parse %s: %v", p, err)
		}
		return &result, p, nil
	}
	return nil, "", fmt.Errorf("wilayah.json tidak ditemukan di: %v", wilayahJSONPaths)
}

func seedWilayah(db *sql.DB) error {
	// Cek apakah sudah ada data
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM provinces").Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil // sudah ada, skip
	}

	// Baca dari file JSON lokal (pre-fetched oleh scripts/fetch_wilayah.py)
	wilayah, filePath, err := loadWilayahFromFile()
	if err != nil {
		return fmt.Errorf("gagal memuat data wilayah: %v", err)
	}
	fmt.Printf("Seeding wilayah dari file lokal: %s\n", filePath)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmtProv, err := tx.Prepare("INSERT OR IGNORE INTO provinces(code, name) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmtProv.Close()

	stmtReg, err := tx.Prepare("INSERT OR IGNORE INTO regencies(code, province_code, name) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtReg.Close()

	for _, prov := range wilayah.Provinces {
		if _, err := stmtProv.Exec(prov.Code, prov.Name); err != nil {
			fmt.Printf("Skip province %s: %v\n", prov.Code, err)
		}
	}

	for _, reg := range wilayah.Regencies {
		if _, err := stmtReg.Exec(reg.Code, reg.ProvinceCode, reg.Name); err != nil {
			fmt.Printf("Skip regency %s: %v\n", reg.Code, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	fmt.Printf("Seeding wilayah selesai: %d provinsi, %d kota/kab\n",
		len(wilayah.Provinces), len(wilayah.Regencies))
	return nil
}

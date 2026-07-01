package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"backend/models"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	DB        *sql.DB
	UploadDir string
	BaseURL   string // URL dasar untuk open graph (misal: http://localhost:8080 atau https://domain.com)
}

func NewProfileHandler(db *sql.DB, uploadDir string, baseURL string) *ProfileHandler {
	// Buat direktori upload jika belum ada
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		fmt.Printf("Gagal membuat direktori upload: %v\n", err)
	}
	return &ProfileHandler{
		DB:        db,
		UploadDir: uploadDir,
		BaseURL:   baseURL,
	}
}

func (h *ProfileHandler) RegisterProfile(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validasi input gagal: " + err.Error()})
		return
	}

	// [Authorization Check] Pastikan email yang diupdate milik pemilik token
	tokenEmail, exists := c.Get("user_email")
	if !exists || tokenEmail.(string) != req.Email {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Anda tidak diizinkan mengubah profil ini."})
		return
	}

	// Serialize RPG Stats map to JSON string
	statsJSON, err := json.Marshal(req.RPGStats)
	if err != nil {
		fmt.Printf("[Security Log] Gagal memproses rpg_stats: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data statistik"})
		return
	}

	// Generate unique slug based on their nickname (nama_panggilan)
	slugVal := makeUniqueSlug(h.DB, generateSlug(req.NamaPanggilan), req.Email)

	// Update database
	res, err := h.DB.Exec(`
		UPDATE participants
		SET nama_panggilan = ?, gender = ?, kota = ?, provinsi = ?, pekerjaan = ?, bio = ?, rpg_class = ?, rpg_stats = ?, linkedin = ?, background_it = ?, motivasi = ?, slug = ?, is_registered = 1, updated_at = CURRENT_TIMESTAMP
		WHERE LOWER(TRIM(email)) = LOWER(TRIM(?))
	`, req.NamaPanggilan, req.Gender, req.Kota, req.Provinsi, req.Pekerjaan, req.Bio, req.RPGClass, string(statsJSON), req.Linkedin, req.BackgroundIT, req.Motivasi, slugVal, req.Email)

	if err != nil {
		fmt.Printf("[Security Log] Gagal menyimpan profil (DB): %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data profil ke server"})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email tidak terdaftar sebagai peserta"})
		return
	}

	// Ambil profil yang baru saja diperbarui
	var p models.Participant
	var namaPanggilan, gender, kota, provinsi, pekerjaan, bio, rpgClass, rpgStats, imagePath, linkedin, backgroundIT, motivasi, slug sql.NullString
	err = h.DB.QueryRow(`
		SELECT id, email, whatsapp, nama_lengkap, nama_panggilan, gender, kota, provinsi, pekerjaan, bio, rpg_class, rpg_stats, image_path, linkedin, background_it, motivasi, slug, is_registered
		FROM participants
		WHERE LOWER(TRIM(email)) = LOWER(TRIM(?))
	`, req.Email).Scan(
		&p.ID, &p.Email, &p.Whatsapp, &p.NamaLengkap,
		&namaPanggilan, &gender, &kota, &provinsi, &pekerjaan, &bio, &rpgClass, &rpgStats, &imagePath, &linkedin, &backgroundIT, &motivasi, &slug,
		&p.IsRegistered,
	)

	if err == nil {
		if namaPanggilan.Valid { p.NamaPanggilan = &namaPanggilan.String }
		if gender.Valid { p.Gender = &gender.String }
		if kota.Valid { p.Kota = &kota.String }
		if provinsi.Valid { p.Provinsi = &provinsi.String }
		if pekerjaan.Valid { p.Pekerjaan = &pekerjaan.String }
		if bio.Valid { p.Bio = &bio.String }
		if rpgClass.Valid { p.RPGClass = &rpgClass.String }
		if rpgStats.Valid { p.RPGStats = &rpgStats.String }
		if imagePath.Valid { p.ImagePath = &imagePath.String }
		if linkedin.Valid { p.Linkedin = &linkedin.String }
		if backgroundIT.Valid { p.BackgroundIT = &backgroundIT.String }
		if motivasi.Valid { p.Motivasi = &motivasi.String }
		if slug.Valid { p.Slug = &slug.String }
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Profil berhasil didaftarkan",
		"participant": p,
	})
}

func (h *ProfileHandler) UploadCardImage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// [Authorization Check] Pastikan ID yang diupload milik pemilik token (Cegah IDOR)
	tokenID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak."})
		return
	}
	// Perlu convert tokenID karena JWT parser (.sub) by default menyimpannya sebagai float64 pada JSON numeric
	var authID int64
	switch v := tokenID.(type) {
	case float64:
		authID = int64(v)
	case int64:
		authID = v
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Tipe data ID tidak valid."})
		return
	}

	if authID != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Anda tidak diizinkan mengubah gambar profil ini."})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File gambar diperlukan"})
		return
	}

	// Validasi ekstensi
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya file PNG/JPG yang diperbolehkan"})
		return
	}

	// Simpan file secara lokal dengan nama ID peserta
	filename := fmt.Sprintf("card_%d%s", id, ext)
	filePath := filepath.Join(h.UploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		fmt.Printf("[Security Log] Gagal menyimpan file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file gambar di server"})
		return
	}

	// Update path di database (gunakan relative URL path untuk web client)
	webPath := fmt.Sprintf("/uploads/%s", filename)
	_, err = h.DB.Exec("UPDATE participants SET image_path = ? WHERE id = ?", webPath, id)
	if err != nil {
		fmt.Printf("[Security Log] Gagal menyimpan path DB: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data gambar ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Gambar kartu berhasil diunggah",
		"image_path": webPath,
	})
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	idOrSlug := c.Param("id")

	var p models.Participant
	var namaPanggilan, gender, kota, provinsi, pekerjaan, bio, rpgClass, rpgStats, imagePath, linkedin, backgroundIT, motivasi, slug sql.NullString

	// Coba cari sebagai ID (integer) terlebih dahulu, jika gagal cari sebagai Slug (string)
	id, err := strconv.ParseInt(idOrSlug, 10, 64)
	var query string
	var param interface{}
	if err == nil {
		query = `
			SELECT id, email, whatsapp, nama_lengkap, nama_panggilan, gender, kota, provinsi, pekerjaan, bio, rpg_class, rpg_stats, image_path, linkedin, background_it, motivasi, slug, is_registered
			FROM participants
			WHERE id = ?`
		param = id
	} else {
		query = `
			SELECT id, email, whatsapp, nama_lengkap, nama_panggilan, gender, kota, provinsi, pekerjaan, bio, rpg_class, rpg_stats, image_path, linkedin, background_it, motivasi, slug, is_registered
			FROM participants
			WHERE slug = ?`
		param = idOrSlug
	}

	err = h.DB.QueryRow(query, param).Scan(
		&p.ID, &p.Email, &p.Whatsapp, &p.NamaLengkap,
		&namaPanggilan, &gender, &kota, &provinsi, &pekerjaan, &bio, &rpgClass, &rpgStats, &imagePath, &linkedin, &backgroundIT, &motivasi, &slug,
		&p.IsRegistered,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Peserta tidak ditemukan"})
		return
	} else if err != nil {
		fmt.Printf("[Security Log] Kesalahan database GetProfile: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan internal server"})
		return
	}

	if namaPanggilan.Valid { p.NamaPanggilan = &namaPanggilan.String }
	if gender.Valid { p.Gender = &gender.String }
	if kota.Valid { p.Kota = &kota.String }
	if provinsi.Valid { p.Provinsi = &provinsi.String }
	if pekerjaan.Valid { p.Pekerjaan = &pekerjaan.String }
	if bio.Valid { p.Bio = &bio.String }
	if rpgClass.Valid { p.RPGClass = &rpgClass.String }
	if rpgStats.Valid { p.RPGStats = &rpgStats.String }
	if imagePath.Valid { p.ImagePath = &imagePath.String }
	if linkedin.Valid { p.Linkedin = &linkedin.String }
	if backgroundIT.Valid { p.BackgroundIT = &backgroundIT.String }
	if motivasi.Valid { p.Motivasi = &motivasi.String }
	if slug.Valid { p.Slug = &slug.String }

	c.JSON(http.StatusOK, p)
}

func (h *ProfileHandler) SharePage(c *gin.Context) {
	idOrSlug := c.Param("id")

	var id int64
	var namaLengkap, rpgClass, imagePath, slugStr sql.NullString
	
	idParsed, err := strconv.ParseInt(idOrSlug, 10, 64)
	var query string
	var param interface{}
	if err == nil {
		query = "SELECT id, nama_lengkap, rpg_class, image_path, slug FROM participants WHERE id = ?"
		param = idParsed
	} else {
		query = "SELECT id, nama_lengkap, rpg_class, image_path, slug FROM participants WHERE slug = ?"
		param = idOrSlug
	}

	err = h.DB.QueryRow(query, param).Scan(&id, &namaLengkap, &rpgClass, &imagePath, &slugStr)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	userAgent := c.GetHeader("User-Agent")
	isCrawler := false
	crawlers := []string{"whatsapp", "facebookexternalhit", "twitterbot", "linkedinbot", "telegrambot", "slackbot", "googlebot"}
	for _, crawler := range crawlers {
		if strings.Contains(strings.ToLower(userAgent), crawler) {
			isCrawler = true
			break
		}
	}

	fullNameSlug := "developer"
	if namaLengkap.Valid && namaLengkap.String != "" {
		fullNameSlug = generateSlug(namaLengkap.String)
	}

	redirectTarget := fmt.Sprintf("/share/%d/%s", id, fullNameSlug)
	if strings.Contains(h.BaseURL, "localhost") {
		redirectTarget = fmt.Sprintf("http://localhost:5173/share/%d/%s", id, fullNameSlug)
	}

	if isCrawler {
		name := "Developer"
		if namaLengkap.Valid {
			name = namaLengkap.String
		}
		class := "Cloud Adventurer"
		if rpgClass.Valid {
			class = rpgClass.String
		}
		imgURL := h.BaseURL + "/assets/placeholder-card.jpg"
		if imagePath.Valid && imagePath.String != "" {
			imgURL = h.BaseURL + imagePath.String
		}
		shareURL := fmt.Sprintf("%s/share/%d/%s", h.BaseURL, id, fullNameSlug)

		html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Developer RPG Card - %s</title>
    <meta property="og:title" content="Developer RPG Card: %s" />
    <meta property="og:description" content="Class: %s | Intip kartu karakter developer dan ikuti ice-breaking kami di AWS AI Academy!" />
    <meta property="og:image" content="%s" />
    <meta property="og:url" content="%s" />
    <meta property="og:type" content="website" />
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:title" content="Developer RPG Card: %s" />
    <meta name="twitter:description" content="Class: %s | Intip kartu karakter developer dan ikuti ice-breaking kami di AWS AI Academy!" />
    <meta name="twitter:image" content="%s" />
</head>
<body>
    <p>Mengalihkan Anda ke halaman kartu perkenalan...</p>
    <script>window.location.href = "%s";</script>
</body>
</html>`, name, name, class, imgURL, shareURL, name, class, imgURL, redirectTarget)

		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, html)
	} else {
		c.Redirect(http.StatusFound, redirectTarget)
	}
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

func makeUniqueSlug(db *sql.DB, baseSlug string, currentEmail string) string {
	slug := baseSlug
	counter := 1
	for {
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM participants WHERE slug = ? AND email != ?)", slug, currentEmail).Scan(&exists)
		if err != nil || !exists {
			break
		}
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}
	return slug
}

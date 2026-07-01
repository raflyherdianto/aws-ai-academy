package handlers

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "super-secret-aws-ai-academy-2026-key"
	}
	return secret
}

// GenerateJWT creates a new JWT token for a participant
func GenerateJWT(participantID int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   participantID,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiry
		"iat":   time.Now().Unix(),
	})

	return token.SignedString(jwtSecret)
}

func (h *AuthHandler) ValidateEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter email diperlukan"})
		return
	}

	var p models.Participant
	var namaPanggilan, gender, kota, provinsi, pekerjaan, bio, rpgClass, rpgStats, imagePath, linkedin, backgroundIT, motivasi, slug sql.NullString

	err := h.DB.QueryRow(`
		SELECT id, email, whatsapp, nama_lengkap, nama_panggilan, gender, kota, provinsi, pekerjaan, bio, rpg_class, rpg_stats, image_path, linkedin, background_it, motivasi, slug, is_registered 
		FROM participants 
		WHERE LOWER(TRIM(email)) = LOWER(TRIM(?))
	`, email).Scan(
		&p.ID, &p.Email, &p.Whatsapp, &p.NamaLengkap,
		&namaPanggilan, &gender, &kota, &provinsi, &pekerjaan, &bio, &rpgClass, &rpgStats, &imagePath, &linkedin, &backgroundIT, &motivasi, &slug,
		&p.IsRegistered,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"exists": false, "message": "Email tidak terdaftar sebagai peserta"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan internal server"})
		return
	}

	// Map null strings to models pointers
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

	// Generate Token
	token, err := GenerateJWT(p.ID, p.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat sesi autentikasi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists":        true,
		"is_registered": p.IsRegistered == 1,
		"participant":   p,
		"token":         token,
	})
}


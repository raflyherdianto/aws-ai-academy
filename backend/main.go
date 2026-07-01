package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"backend/database"
	"backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// CORSMiddleware mengatur izin CORS menggunakan whitelist dari .env
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
		
		isAllowed := false
		if allowedOrigins == "" {
			isAllowed = true // Fallback untuk development lokal
		} else {
			originsList := strings.Split(allowedOrigins, ",")
			for _, allowed := range originsList {
				if origin == strings.TrimSpace(allowed) {
					isAllowed = true
					break
				}
			}
		}

		if origin != "" && isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SecurityHeadersMiddleware menambahkan HTTP Security Headers untuk mitigasi XSS, Clickjacking, dan MIME-sniffing
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:;")
		c.Next()
	}
}

// JWTMiddleware memastikan request memiliki token JWT yang valid
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses ditolak. Token tidak ditemukan."})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "super-secret-aws-ai-academy-2026-key"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode penandatanganan tidak terduga: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak valid atau telah kadaluarsa."})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_email", claims["email"])
			c.Set("user_id", claims["sub"])
		}

		c.Next()
	}
}

func findSeedFile() string {
	// Cari file mentee.json di beberapa tempat potensial
	paths := []string{
		"./data/mentee.json", // Prioritas utama untuk Docker volume
		"../mentee.json",
		"./mentee.json",
		"mentee.json",
		"../../mentee.json",
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			abs, _ := filepath.Abs(p)
			fmt.Printf("Menemukan berkas seed di: %s\n", abs)
			return p
		}
	}
	return "mentee.json" // default fallback
}

func main() {
	// Muat file .env jika ada (ignore error jika tidak ada)
	_ = godotenv.Load()

	// Set paths
	dbPath := "./data/academy.db"
	schemaPath := "./database/schema.sql"
	seedPath := findSeedFile()
	uploadDir := "./uploads"

	// Tentukan Base URL untuk Open Graph preview (gunakan domain production jika kosong)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "https://fasilrafly.cobalab.web.id"
	}

	// Inisialisasi basis data
	db, err := database.InitDB(dbPath, schemaPath, seedPath)
	if err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	defer db.Close()

	// Inisialisasi Gin router
	r := gin.Default()
	r.Use(SecurityHeadersMiddleware())
	r.Use(CORSMiddleware())

	// Sajikan file statis kartu profil PNG dari folder uploads
	r.Static("/uploads", uploadDir)

	// Tambahkan rute untuk placeholder/fallback asset jika dibutuhkan
	r.StaticFS("/assets", http.Dir("./assets"))

	// Pastikan folder assets ada dan berisi gambar default jika di-crawl
	_ = os.MkdirAll("./assets", 0755)
	placeholderPath := "./assets/placeholder-card.jpg"
	if _, err := os.Stat(placeholderPath); os.IsNotExist(err) {
		// Buat placeholder kosong jika tidak disalin
		_ = os.WriteFile(placeholderPath, []byte("placeholder"), 0644)
	}

	// Inisialisasi Handlers
	authHandler := handlers.NewAuthHandler(db)
	profileHandler := handlers.NewProfileHandler(db, uploadDir, baseURL)
	gameHandler := handlers.NewGameHandler(db)
	wilayahHandler := handlers.NewWilayahHandler(db)

	// API Routes
	api := r.Group("/api")
	{
		api.GET("/auth", authHandler.ValidateEmail)
		api.GET("/profile/:id", profileHandler.GetProfile)
		api.GET("/game/quiz", gameHandler.GetQuiz)
		api.GET("/wilayah/provinces", wilayahHandler.GetProvinces)
		api.GET("/wilayah/regencies/:province_code", wilayahHandler.GetRegencies)

		// Protected Routes
		protected := api.Group("/")
		protected.Use(JWTMiddleware())
		{
			protected.POST("/profile", profileHandler.RegisterProfile)
			protected.POST("/profile/:id/upload", profileHandler.UploadCardImage)
		}
	}

	// Route khusus link sharing (WhatsApp Open Graph)
	r.GET("/share/:id", profileHandler.SharePage)
	r.GET("/share/:id/:nama_lengkap", profileHandler.SharePage)

	// Route proxy avatar to prevent CORS taint in html2canvas (returns PNG format for 100% canvas compatibility)
	r.GET("/api/avatar", func(c *gin.Context) {
		seed := c.Query("seed")
		if seed == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter seed diperlukan"})
			return
		}

		targetURL := fmt.Sprintf("https://api.dicebear.com/7.x/pixel-art/png?seed=%s", url.QueryEscape(seed))
		resp, err := http.Get(targetURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil avatar: " + err.Error()})
			return
		}
		defer resp.Body.Close()

		pngBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data avatar"})
			return
		}

		c.Header("Content-Type", "image/png")
		c.Data(http.StatusOK, "image/png", pngBytes)
	})

	// Route health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "db": db.Ping() == nil})
	})

	// Jalankan server di port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server berjalan di port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}

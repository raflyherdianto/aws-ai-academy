package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	DB *sql.DB
}

func NewGameHandler(db *sql.DB) *GameHandler {
	return &GameHandler{DB: db}
}

func (h *GameHandler) GetQuiz(c *gin.Context) {
	// Pilih 1 peserta secara acak yang sudah terregistrasi (is_registered = 1)
	var correct models.Participant
	var namaPanggilan, gender, kota, provinsi, pekerjaan, bio, rpgClass, rpgStats, imagePath, linkedin, backgroundIT, motivasi sql.NullString

	err := h.DB.QueryRow(`
		SELECT id, email, whatsapp, nama_lengkap, nama_panggilan, gender, kota, provinsi, pekerjaan, bio, rpg_class, rpg_stats, image_path, linkedin, background_it, motivasi, is_registered
		FROM participants
		WHERE is_registered = 1
		ORDER BY RANDOM()
		LIMIT 1
	`).Scan(
		&correct.ID, &correct.Email, &correct.Whatsapp, &correct.NamaLengkap,
		&namaPanggilan, &gender, &kota, &provinsi, &pekerjaan, &bio, &rpgClass, &rpgStats, &imagePath, &linkedin, &backgroundIT, &motivasi,
		&correct.IsRegistered,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Belum ada peserta yang mengisi profil. Game belum bisa dimainkan!"})
		return
	} else if err != nil {
		fmt.Printf("[Security Log] Kesalahan database GetQuiz: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan internal server"})
		return
	}

	if namaPanggilan.Valid { correct.NamaPanggilan = &namaPanggilan.String }
	if gender.Valid { correct.Gender = &gender.String }
	if kota.Valid { correct.Kota = &kota.String }
	if provinsi.Valid { correct.Provinsi = &provinsi.String }
	if pekerjaan.Valid { correct.Pekerjaan = &pekerjaan.String }
	if bio.Valid { correct.Bio = &bio.String }
	if rpgClass.Valid { correct.RPGClass = &rpgClass.String }
	if rpgStats.Valid { correct.RPGStats = &rpgStats.String }
	if imagePath.Valid { correct.ImagePath = &imagePath.String }
	if linkedin.Valid { correct.Linkedin = &linkedin.String }
	if backgroundIT.Valid { correct.BackgroundIT = &backgroundIT.String }
	if motivasi.Valid { correct.Motivasi = &motivasi.String }

	// Ambil 3 nama acak lain sebagai pengecoh
	rows, err := h.DB.Query("SELECT nama_lengkap FROM participants WHERE id != ? ORDER BY RANDOM() LIMIT 3", correct.ID)
	var choices []string
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err == nil {
				choices = append(choices, name)
			}
		}
	}

	// Tambahkan nama yang benar ke pilihan
	choices = append(choices, correct.NamaLengkap)

	// Acak urutan pilihan jawaban
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	// Bangun petunjuk (clues)
	var clues []string
	if correct.Pekerjaan != nil && *correct.Pekerjaan != "" {
		clues = append(clues, "Pekerjaan saya saat ini: "+*correct.Pekerjaan)
	}
	if correct.Kota != nil && correct.Provinsi != nil && *correct.Kota != "" {
		clues = append(clues, fmt.Sprintf("Saya tinggal di Kota %s, Provinsi %s", *correct.Kota, *correct.Provinsi))
	}
	if correct.RPGClass != nil && *correct.RPGClass != "" {
		clues = append(clues, "RPG Class saya: "+*correct.RPGClass)
	}
	if correct.Bio != nil && *correct.Bio != "" {
		clues = append(clues, "Bio singkat saya: \""+*correct.Bio+"\"")
	}
	if correct.Motivasi != nil && *correct.Motivasi != "" {
		clues = append(clues, "Motivasi saya mengikuti AWS AI Academy: \""+*correct.Motivasi+"\"")
	}

	// Parsir status RPG jika ada
	if correct.RPGStats != nil && *correct.RPGStats != "" {
		var stats map[string]int
		if err := json.Unmarshal([]byte(*correct.RPGStats), &stats); err == nil {
			for statName, val := range stats {
				clues = append(clues, fmt.Sprintf("Stat RPG %s saya bernilai %d", statName, val))
			}
		}
	}

	// Acak urutan petunjuk agar tidak terlalu mudah ditebak polanya
	rand.Shuffle(len(clues), func(i, j int) {
		clues[i], clues[j] = clues[j], clues[i]
	})

	c.JSON(http.StatusOK, models.QuizQuestion{
		CorrectID: correct.ID,
		Clues:     clues,
		Choices:   choices,
		Correct:   correct.NamaLengkap,
	})
}

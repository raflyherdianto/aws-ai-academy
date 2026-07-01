package models

import "time"

type Participant struct {
	ID            int64     `json:"id"`
	Email         string    `json:"email"`
	Whatsapp      string    `json:"whatsapp"`
	NamaLengkap   string    `json:"nama_lengkap"`
	NamaPanggilan *string   `json:"nama_panggilan"`
	Gender        *string   `json:"gender"`
	Kota          *string   `json:"kota"`
	Provinsi      *string   `json:"provinsi"`
	Pekerjaan     *string   `json:"pekerjaan"`
	Bio           *string   `json:"bio"`
	RPGClass      *string   `json:"rpg_class"`
	RPGStats      *string   `json:"rpg_stats"` // JSON string
	ImagePath     *string   `json:"image_path"`
	Linkedin      *string   `json:"linkedin"`
	BackgroundIT  *string   `json:"background_it"`
	Motivasi      *string   `json:"motivasi"`
	Slug          *string   `json:"slug"`
	IsRegistered  int       `json:"is_registered"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Email         string         `json:"email" binding:"required,email"`
	NamaPanggilan string         `json:"nama_panggilan" binding:"required"`
	Gender        string         `json:"gender" binding:"required"`
	Kota          string         `json:"kota" binding:"required"`
	Provinsi      string         `json:"provinsi" binding:"required"`
	Pekerjaan     string         `json:"pekerjaan" binding:"required"`
	Bio           string         `json:"bio" binding:"required"`
	RPGClass      string         `json:"rpg_class" binding:"required"`
	RPGStats      map[string]int `json:"rpg_stats" binding:"required"`
	Linkedin      string         `json:"linkedin"`
	BackgroundIT  string         `json:"background_it" binding:"required"`
	Motivasi      string         `json:"motivasi" binding:"required"`
}

type QuizQuestion struct {
	CorrectID int64    `json:"correct_id"`
	Clues     []string `json:"clues"`   // e.g., ["Domisili: Jakarta", "Pekerjaan: Student", "Bio: Hello guys!"]
	Choices   []string `json:"choices"` // 4 full names
	Correct   string   `json:"correct"` // Correct name
}

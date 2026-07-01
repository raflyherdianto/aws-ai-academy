package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WilayahHandler struct {
	DB *sql.DB
}

func NewWilayahHandler(db *sql.DB) *WilayahHandler {
	return &WilayahHandler{DB: db}
}

type ProvinceResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type RegencyResponse struct {
	Code         string `json:"code"`
	ProvinceCode string `json:"province_code"`
	Name         string `json:"name"`
}

// GET /api/wilayah/provinces
func (h *WilayahHandler) GetProvinces(c *gin.Context) {
	rows, err := h.DB.Query("SELECT code, name FROM provinces ORDER BY name ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data provinsi"})
		return
	}
	defer rows.Close()

	var provinces []ProvinceResponse
	for rows.Next() {
		var p ProvinceResponse
		if err := rows.Scan(&p.Code, &p.Name); err != nil {
			continue
		}
		provinces = append(provinces, p)
	}
	c.JSON(http.StatusOK, gin.H{"data": provinces})
}

// GET /api/wilayah/regencies/:province_code
func (h *WilayahHandler) GetRegencies(c *gin.Context) {
	provinceCode := c.Param("province_code")
	rows, err := h.DB.Query(
		"SELECT code, province_code, name FROM regencies WHERE province_code = ? ORDER BY name ASC",
		provinceCode,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kota/kabupaten"})
		return
	}
	defer rows.Close()

	var regencies []RegencyResponse
	for rows.Next() {
		var r RegencyResponse
		if err := rows.Scan(&r.Code, &r.ProvinceCode, &r.Name); err != nil {
			continue
		}
		regencies = append(regencies, r)
	}
	c.JSON(http.StatusOK, gin.H{"data": regencies})
}

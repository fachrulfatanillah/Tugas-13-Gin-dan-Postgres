package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Tugas-13-Gin-dan-Postgres/database"
	"Tugas-13-Gin-dan-Postgres/model"
)

type InputBioskop struct {
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}

func CreateBioskop(c *gin.Context) {
	var input InputBioskop

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input tidak valid",
		})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama dan Lokasi wajib diisi",
		})
		return
	}

	bioskop := model.Bioskop{
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	database.DB.Create(&bioskop)

	c.JSON(http.StatusCreated, bioskop)
}

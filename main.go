package main

import (
	"Tugas-13-Gin-dan-Postgres/database"
	"Tugas-13-Gin-dan-Postgres/handler"
	"Tugas-13-Gin-dan-Postgres/model"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(&model.Bioskop{})

	r := gin.Default()

	r.POST("/bioskop", handler.CreateBioskop)

	log.Println("Server berjalan di http://localhost:8080")
	r.Run(":8080")
}

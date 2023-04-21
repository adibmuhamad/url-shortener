package main

import (
	"id/projects/url-shortener/shorter"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	repository := shorter.NewRepository(db)
	service := shorter.NewService(repository)
	handler := shorter.NewHandler(service)

	router := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST"}
	config.AllowHeaders = []string{"Content-Type"}
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	api := router.Group("/api/v1")

	api.POST("/generateShorterUrl", handler.GenerateShorterUrl)
	api.POST("/saveShorterUrl", handler.SaveShorterUrl)

	router.GET("/:backHalf", handler.RedirectShorterUrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}

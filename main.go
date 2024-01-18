package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kankanya45/go-gorm-db/db"
	"github.com/Kankanya45/go-gorm-db/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	database, err := db.ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = database.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&models.Subject{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	itemRepo := models.NewItemRepository(database)
	subjectRepo := models.NewSubjectRepository(database)

	r := gin.Default()

	// Middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// Routes setup
	setupRoutes(r, itemRepo, subjectRepo)

	// No route handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Run the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Server is not running: %v", err)
	}
}

func setupRoutes(r *gin.Engine, itemRepo *models.ItemRepository, subjectRepo *models.SubjectRepository) {
	// API routes for items
	r.GET("/items", itemRepo.GetItems)
	r.POST("/items", itemRepo.PostItem)
	r.GET("/items/:id", itemRepo.GetItem)
	r.PUT("/items/:id", itemRepo.UpdateItem)
	r.DELETE("/items/:id", itemRepo.DeleteItem)

	// API routes for subjects
	r.GET("/subjects", subjectRepo.GetSubjects)
	r.POST("/subjects", subjectRepo.PostSubject)
	r.GET("/subjects/:id", subjectRepo.GetSubject)
	r.PUT("/subjects/:id", subjectRepo.UpdateSubject)
	r.DELETE("/subjects/:id", subjectRepo.DeleteSubject)

	// Add other routes as needed
}

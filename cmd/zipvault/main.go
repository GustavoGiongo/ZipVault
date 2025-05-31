package main

import (
	"ZipVault/internal/zipvault/api/http"
	"ZipVault/internal/zipvault/app"
	"ZipVault/internal/zipvault/infra/mysql/mysql"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	host, port, user, password, name := initDb()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	log.Println("Conectando ao banco com:", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("MySQL didn't respond", err)
	}

	log.Println("MySql connection is open!")
	defer db.Close()
	repository := mysql.NewMySQLRepository(db)
	listingService := app.NewListingService(repository)
	listHandler := http.NewListingHandler(listingService)

	router := gin.Default()
	router.GET("/health", http.HealthHandler)
	router.GET("/list", listHandler.HandleListing)
	router.Run(":8080")
}

func initDb() (string, string, string, string, string) {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	return host, port, user, password, name
}

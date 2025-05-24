package main

import (
	"ZipVault/internal/zipvault/api/http"
	"ZipVault/internal/zipvault/app"
	"ZipVault/internal/zipvault/infra/mysql/mysql"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/folders")
	if err != nil {
		log.Println("Error connecting to database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("MySQL didn't respond", err)
	}

	log.Println("MySql connection is open!")
	defer db.Close()
	downloadRepo := mysql.NewMySQLDownloadRepository(db)
	downLoadService := app.NewDownloadService(downloadRepo)
	downloadHandler := http.NewDownloadHandler(downLoadService)

	router := gin.Default()
	router.GET("/health", http.HealthHandler)
	router.GET("/download", downloadHandler.HandleDownload)
	router.Run(":8080")
}

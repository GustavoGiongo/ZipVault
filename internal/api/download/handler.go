package download

import (
	"ZipVault/internal/dto"
	"ZipVault/internal/repository"
	"ZipVault/internal/service"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DownloadHandler(c *gin.Context) {
	var req dto.DownloadRequest
	fmt.Println(req)
	db, err := sql.Open("mysql", "usuario:senha@tcp(localhost:3306)/database")
	if err != nil {
		log.Println("Error connecting to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
	}
	repo := repository.NewMySQLDownloadRepository(db)
	service := service.NewDownloadService(repo)
	foundFile := service.HandleDownload(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}
	if foundFile {
		c.JSON(http.StatusOK, gin.H{"message:": "File was Found"})
	}
	c.JSON(http.StatusNoContent, gin.H{"message:": "File was not Found"})
}

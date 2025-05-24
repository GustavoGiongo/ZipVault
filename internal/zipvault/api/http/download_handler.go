package http

import (
	"ZipVault/internal/zipvault/app"
	"ZipVault/internal/zipvault/dto/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DownloadHandler struct {
	Service *app.DownloadService
}

func NewDownloadHandler(svc *app.DownloadService) *DownloadHandler {
	return &DownloadHandler{Service: svc}
}

func (dh *DownloadHandler) HandleDownload(c *gin.Context) {
	var req dto.DownloadRequest
	fmt.Println(req)
	FolderName, err := dh.Service.SearchFolderName(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}
	if FolderName {
		c.JSON(http.StatusOK, gin.H{"message:": "File was Found"})
	}
	c.JSON(http.StatusNoContent, gin.H{"message:": "File was not Found"})
}

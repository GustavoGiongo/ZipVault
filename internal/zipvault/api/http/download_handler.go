package http

import (
	"ZipVault/internal/zipvault/app"
	"ZipVault/internal/zipvault/dto/dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type DownloadHandler struct {
	Service *app.DownloadService
}

func NewDownloadHandler(svc *app.DownloadService) *DownloadHandler {
	return &DownloadHandler{Service: svc}
}

type FilteredFileResponse struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Folder   string `json:"folder"`
	FullPath string `json:"fullPath"`
}

func (dh *DownloadHandler) HandleListing(c *gin.Context) {
	var req dto.ListingRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters!"})
		return
	}

	if req.Name == "" && req.StartDate == "" && req.EndDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one parameter should be informed!"})
		return
	}

	fileRecords, err := dh.Service.SearchFileRecords(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filteredRecordsResponse := filterRecords(fileRecords)

	c.JSON(http.StatusOK, gin.H{
		"files": filteredRecordsResponse,
	})

	log.Println("Filtered file list returned")
}

func filterRecords(fileRecords []dto.FileRecord) []FilteredFileResponse {
	base := os.Getenv("BASE_PATH")
	var (
		wg                      sync.WaitGroup
		filteredRecordsResponse []FilteredFileResponse
	)

	resultChan := make(chan FilteredFileResponse)

	for _, fileRecord := range fileRecords {
		wg.Add(1)
		go func(record dto.FileRecord) {
			defer wg.Done()

			fullPath := filepath.Join(base, record.Folder, record.Name)

			if _, err := os.Stat(fullPath); err == nil {
				resultChan <- FilteredFileResponse{
					Name:     record.Name,
					Date:     record.Date,
					Folder:   record.Folder,
					FullPath: fullPath,
				}
			}
		}(fileRecord)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		filteredRecordsResponse = append(filteredRecordsResponse, result)
	}
	return filteredRecordsResponse
}

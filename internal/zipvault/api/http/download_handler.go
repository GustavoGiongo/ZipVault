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
	FullPath string `json:"full_path"`
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

	base := os.Getenv("BASE_PATH")

	var (
		mu                      sync.Mutex
		wg                      sync.WaitGroup
		filteredRecordsResponse []FilteredFileResponse
	)

	//TODO check why it's returning only two files and not three
	for _, fileRecord := range fileRecords {
		wg.Add(1)
		go func(record dto.FileRecord) {
			defer wg.Done()

			fullPath := filepath.Join(base, record.Folder, record.Name)

			if _, err := os.Stat(fullPath); err == nil {
				filtered := FilteredFileResponse{
					Name:     record.Name,
					Date:     record.Date,
					Folder:   record.Folder,
					FullPath: fullPath,
				}
				mu.Lock()
				filteredRecordsResponse = append(filteredRecordsResponse, filtered)
				mu.Unlock()
			}
		}(fileRecord)
	}

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": "Files found",
		"files":   filteredRecordsResponse,
	})

	log.Println("Filtered file list returned")
}

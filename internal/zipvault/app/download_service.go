package app

import (
	download "ZipVault/internal/zipvault/domain/download"
	"ZipVault/internal/zipvault/dto/dto"
	"strings"
)

type DownloadService struct {
	repo download.DownloadRepository
}

func (s DownloadService) SearchFileLocation(d *dto.DownloadRequest) (string, error) {
	d.Name = Sanitize(d.Name)
	d.Date = Sanitize(d.Date)
	var fileLocation, err = s.repo.FindFileLocationByNameAndDate(d.Name, d.Date)
	if err != nil {
		return fileLocation, err
	}
	return fileLocation, nil
}

func NewDownloadService(repo download.DownloadRepository) *DownloadService {
	return &DownloadService{repo: repo}
}

func Sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "")

}

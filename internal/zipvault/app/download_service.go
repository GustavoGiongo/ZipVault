package app

import (
	download "ZipVault/internal/zipvault/domain/download"
	"ZipVault/internal/zipvault/dto/dto"
	"strings"
)

type DownloadService struct {
	repo download.DownloadRepository
}

func (s DownloadService) SearchFolderName(d *dto.DownloadRequest) (bool, error) {
	d.Name = Sanitize(d.Name)
	d.Date = Sanitize(d.Date)
	var b, err = s.repo.FindByNameAndDate(d.Name, d.Date)
	if err != nil {
		return b, err
	}
	return b, nil
}

func NewDownloadService(repo download.DownloadRepository) *DownloadService {
	return &DownloadService{repo: repo}
}

func Sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "")

}

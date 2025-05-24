package service

import (
	"ZipVault/internal/domain"
	"ZipVault/internal/dto"
	"ZipVault/internal/repository"
	"strings"
)

type DownloadService struct {
	repo domain.DownloadRepository
}

func (s DownloadService) HandleDownload(d *dto.DownloadRequest) bool {
	d.Name = Sanitize(d.Name)
	d.Date = Sanitize(d.Date)
	_, err := s.repo.FindByNameAndDate(d.Name, d.Date)
	if err != nil {
		return false
	}
	return true
}

func NewDownloadService(repo *repository.DownloadRepositoryImpl) *DownloadService {
	return &DownloadService{repo: repo}
}

func Sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "")

}

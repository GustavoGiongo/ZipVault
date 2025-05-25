package downloadInterface

import "ZipVault/internal/zipvault/dto/dto"

type DownloadRepository interface {
	FindFileRecordsByNameOrDate(name, startDate, endDate string) ([]dto.FileRecord, error)
}

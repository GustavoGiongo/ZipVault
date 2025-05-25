package mysql

import (
	"ZipVault/internal/zipvault/dto/dto"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type DownloadRepositoryImpl struct {
	Db *sql.DB
}

func NewMySQLDownloadRepository(db *sql.DB) *DownloadRepositoryImpl {
	return &DownloadRepositoryImpl{Db: db}
}

func (r *DownloadRepositoryImpl) FindFileRecordsByNameOrDate(name, startDate, endDate string) ([]dto.FileRecord, error) {
	var (
		query      = "SELECT name,date,folder FROM files"
		conditions []string
		args       = []interface{}{}
	)

	if name != "" {
		conditions = append(conditions, "LOWER(name) LIKE ?")
		args = append(args, "%"+strings.ToLower(name)+"%")
	}

	if startDate != "" && endDate != "" {
		conditions = append(conditions, "date BETWEEN ? AND ?")
		args = append(args, startDate, endDate)
	} else if startDate != "" {
		conditions = append(conditions, "date = ?")
		args = append(args, startDate)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	stmt, err := r.Db.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fileRecords []dto.FileRecord

	for rows.Next() {
		var record dto.FileRecord
		if err := rows.Scan(&record.Name, &record.Date, &record.Folder); err != nil {
			return nil, err
		}
		fileRecords = append(fileRecords, record)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return fileRecords, nil
}

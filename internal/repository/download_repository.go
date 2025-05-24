package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DownloadRepositoryImpl struct {
	Db *sql.DB
}

func NewMySQLDownloadRepository(db *sql.DB) *DownloadRepositoryImpl {
	return &DownloadRepositoryImpl{Db: db}
}

func (r *DownloadRepositoryImpl) FindByNameAndDate(name string, date string) (bool, error) {

	return true, nil
}

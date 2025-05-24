package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DownloadRepositoryImpl struct {
	Db *sql.DB
}

func NewMySQLDownloadRepository(db *sql.DB) *DownloadRepositoryImpl {
	return &DownloadRepositoryImpl{Db: db}
}

func (r *DownloadRepositoryImpl) FindFileLocationByNameAndDate(name, date string) (string, error) {
	log.Println("name", name)
	log.Println("date", date)
	query := "SELECT location FROM files WHERE name = ? AND date = ?"
	stmt, err := r.Db.Prepare(query)

	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var fileLocation string
	err = stmt.QueryRow(name, date).Scan(&fileLocation)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

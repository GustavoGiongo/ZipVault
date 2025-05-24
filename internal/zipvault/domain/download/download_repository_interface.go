package downloadusuario

type DownloadRepository interface {
	FindFileLocationByNameAndDate(name, date string) (string, error)
}

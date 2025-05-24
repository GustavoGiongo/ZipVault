package downloadusuario

type DownloadRepository interface {
	FindByNameAndDate(name, date string) (bool, error)
}

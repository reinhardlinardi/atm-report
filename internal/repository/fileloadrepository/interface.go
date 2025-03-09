package fileloadrepository

type Repository interface {
	IsExist(atmId, date string) (bool, error)
}

package atmrepository

type Repository interface {
	IsExist(atmId string) (bool, error)
}

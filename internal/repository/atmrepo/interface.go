package atmrepo

type Repository interface {
	IsExist(atmId string) (bool, error)
}

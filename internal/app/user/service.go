package user

type Service struct {
	repo *RepositoryImpl
}

func NewService(repo *RepositoryImpl) Service {
	return Service{
		repo: repo,
	}
}

package user

type Service struct {
	repo *RepositoryImpl
}

func NewService(repo *RepositoryImpl) Service {
	return Service{
		repo: repo,
	}
}

// Create a user.
func (s Service) Create(user User) (err error) {
	err = s.repo.Save(user)
	return
}

// Find a user for given email.
func (s Service) Find(email string) (user User, err error) {
	user, err = s.repo.Find(email)
	return
}

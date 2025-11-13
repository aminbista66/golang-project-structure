package user


type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUser(id int64) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateUser(u *User) error {
	return s.repo.Create(u)
}

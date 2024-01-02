package user

func NewService(repository Repository) Service {
	return &service{}
}

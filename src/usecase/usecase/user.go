package usecase

import (
	e "airland-server/src/domain/entities"
	"airland-server/src/usecase/repository"
	"errors"
)

type userUsecase struct {
	userRepository repository.IUserRepository
	dBRepository   repository.DBRepository
}

type User interface {
	List(u []*e.User) ([]*e.User, error)
	Create(u *e.User) (*e.User, error)
}

func NewUserUsecase(r repository.IUserRepository, d repository.DBRepository) User {
	return &userUsecase{r, d}
}

func (uu *userUsecase) List(u []*e.User) ([]*e.User, error) {
	u, err := uu.userRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Create(u *e.User) (*e.User, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.userRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	user, ok := data.(*e.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

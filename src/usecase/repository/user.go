package repository

import (
	e "airland-server/src/domain/entities"
)

type IUserRepository interface {
	FindAll(u []*e.User) ([]*e.User, error)
	Create(u *e.User) (*e.User, error)
}

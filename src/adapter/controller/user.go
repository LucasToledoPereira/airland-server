package controller

import (
	"airland-server/src/usecase/usecase"
	"net/http"

	e "airland-server/src/domain/entities"
)

type userController struct {
	userUsecase usecase.User
}

type IUser interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

func NewUserController(us usecase.User) IUser {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx Context) error {
	var u []*e.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(ctx Context) error {
	var params e.User

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.userUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}

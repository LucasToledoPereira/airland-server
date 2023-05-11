package registry

import (
	"airland-server/src/adapter/controller"
	"airland-server/src/adapter/repository"
	"airland-server/src/usecase/usecase"
)

func (r *registry) NewUserController() controller.IUser {
	u := usecase.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUserController(u)
}

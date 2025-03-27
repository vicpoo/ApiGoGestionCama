// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/application"
)

func InitRolDependencies() (
	*CreateRolController,
	*GetRolByIdController,
	*UpdateRolController,
	*DeleteRolController,
	*GetAllRolesController,
) {
	repo := NewMySQLRolRepository()

	createUseCase := application.NewCreateRolUseCase(repo)
	getByIdUseCase := application.NewGetRolByIdUseCase(repo)
	updateUseCase := application.NewUpdateRolUseCase(repo)
	deleteUseCase := application.NewDeleteRolUseCase(repo)
	getAllUseCase := application.NewGetAllRolesUseCase(repo)

	createController := NewCreateRolController(createUseCase)
	getByIdController := NewGetRolByIdController(getByIdUseCase)
	updateController := NewUpdateRolController(updateUseCase)
	deleteController := NewDeleteRolController(deleteUseCase)
	getAllController := NewGetAllRolesController(getAllUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController
}

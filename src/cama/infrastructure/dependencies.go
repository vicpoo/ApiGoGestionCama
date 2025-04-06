// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
)

func InitCamaDependencies() (
	*CreateCamaController,
	*GetCamaByIdController,
	*UpdateCamaController,
	*DeleteCamaController,
	*GetAllCamasController,
	*GetCamasByUsuarioIDController,
	*GetCamasByTipoIDController,
) {
	repo := NewMySQLCamaRepository()

	// Inicialización de casos de uso
	createUseCase := application.NewCreateCamaUseCase(repo)
	getByIdUseCase := application.NewGetCamaByIdUseCase(repo)
	updateUseCase := application.NewUpdateCamaUseCase(repo)
	deleteUseCase := application.NewDeleteCamaUseCase(repo)
	getAllUseCase := application.NewGetAllCamasUseCase(repo)
	getByUsuarioUseCase := application.NewGetCamasByUsuarioIDUseCase(repo)
	getByTipoUseCase := application.NewGetCamasByTipoIDUseCase(repo)

	// Inicialización de controladores
	createController := NewCreateCamaController(createUseCase)
	getByIdController := NewGetCamaByIdController(getByIdUseCase)
	updateController := NewUpdateCamaController(updateUseCase)
	deleteController := NewDeleteCamaController(deleteUseCase)
	getAllController := NewGetAllCamasController(getAllUseCase)
	getByUsuarioController := NewGetCamasByUsuarioIDController(getByUsuarioUseCase)
	getByTipoController := NewGetCamasByTipoIDController(getByTipoUseCase)

	return createController, getByIdController, updateController,
		deleteController, getAllController, getByUsuarioController, getByTipoController
}

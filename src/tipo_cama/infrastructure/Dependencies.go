package infrastructure

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/application"
)

func InitTipoCamaDependencies() (
	*CreateTipoCamaController,
	*GetTipoCamaByIdController,
	*UpdateTipoCamaController,
	*DeleteTipoCamaController,
	*GetAllTiposCamaController,
) {
	repo := NewMySQLTipoCamaRepository()

	createUseCase := application.NewCreateTipoCamaUseCase(repo)
	getByIdUseCase := application.NewGetTipoCamaByIdUseCase(repo)
	updateUseCase := application.NewUpdateTipoCamaUseCase(repo)
	deleteUseCase := application.NewDeleteTipoCamaUseCase(repo)
	getAllUseCase := application.NewGetAllTiposCamaUseCase(repo)

	createController := NewCreateTipoCamaController(createUseCase)
	getByIdController := NewGetTipoCamaByIdController(getByIdUseCase)
	updateController := NewUpdateTipoCamaController(updateUseCase)
	deleteController := NewDeleteTipoCamaController(deleteUseCase)
	getAllController := NewGetAllTiposCamaController(getAllUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController
}

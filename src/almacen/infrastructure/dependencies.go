// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
)

func InitAlmacenDependencies() (
	*CreateAlmacenController,
	*GetAlmacenByIdController,
	*UpdateAlmacenController,
	*DeleteAlmacenController,
	*GetAllAlmacenesController,
	*IncrementarCantidadController,
) {
	repo := NewMySQLAlmacenRepository()

	createUseCase := application.NewCreateAlmacenUseCase(repo)
	getByIdUseCase := application.NewGetAlmacenByIdUseCase(repo)
	updateUseCase := application.NewUpdateAlmacenUseCase(repo)
	deleteUseCase := application.NewDeleteAlmacenUseCase(repo)
	getAllUseCase := application.NewGetAllAlmacenesUseCase(repo)
	incrementarUseCase := application.NewIncrementarCantidadUseCase(repo)

	createController := NewCreateAlmacenController(createUseCase)
	getByIdController := NewGetAlmacenByIdController(getByIdUseCase)
	updateController := NewUpdateAlmacenController(updateUseCase)
	deleteController := NewDeleteAlmacenController(deleteUseCase)
	getAllController := NewGetAllAlmacenesController(getAllUseCase)
	incrementarController := NewIncrementarCantidadController(incrementarUseCase)

	return createController, getByIdController, updateController,
		deleteController, getAllController, incrementarController
}

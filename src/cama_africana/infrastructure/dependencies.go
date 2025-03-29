package infrastructure

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

func InitCamaAfricanaDependencies() (
	*CreateCamaAfricanaController,
	*GetCamaAfricanaByIdController,
	*UpdateCamaAfricanaController,
	*DeleteCamaAfricanaController,
	*GetAllCamasAfricanasController,
	*GetCamaAfricanaByCamaIDController,
	*GetCamasAfricanasByUsuarioIDController,
	*AssignCamaToUsuarioController,
) {
	repo := NewMySQLCamaAfricanaRepository()

	// Casos de uso básicos
	createUseCase := application.NewCreateCamaAfricanaUseCase(repo)
	getByIdUseCase := application.NewGetCamaAfricanaByIdUseCase(repo)
	updateUseCase := application.NewUpdateCamaAfricanaUseCase(repo)
	deleteUseCase := application.NewDeleteCamaAfricanaUseCase(repo)
	getAllUseCase := application.NewGetAllCamasAfricanasUseCase(repo)

	// Casos de uso adicionales
	getByCamaIDUseCase := application.NewGetCamaAfricanaByCamaIDUseCase(repo)
	getByUsuarioIDUseCase := application.NewGetCamasAfricanasByUsuarioIDUseCase(repo)
	assignUseCase := application.NewAssignCamaToUsuarioUseCase(repo)

	// Controladores
	createController := NewCreateCamaAfricanaController(createUseCase)
	getByIdController := NewGetCamaAfricanaByIdController(getByIdUseCase)
	updateController := NewUpdateCamaAfricanaController(updateUseCase)
	deleteController := NewDeleteCamaAfricanaController(deleteUseCase)
	getAllController := NewGetAllCamasAfricanasController(getAllUseCase)
	getByCamaIDController := NewGetCamaAfricanaByCamaIDController(getByCamaIDUseCase)
	getByUsuarioIDController := NewGetCamasAfricanasByUsuarioIDController(getByUsuarioIDUseCase)
	assignController := NewAssignCamaToUsuarioController(assignUseCase)

	return createController, getByIdController, updateController,
		deleteController, getAllController, getByCamaIDController,
		getByUsuarioIDController, assignController
}

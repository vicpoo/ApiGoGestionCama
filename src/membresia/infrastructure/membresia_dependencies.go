// membresia_dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

func InitMembresiaDependencies() (
	*CreateMembresiaController,
	*GetMembresiaByIdController,
	*UpdateMembresiaController,
	*DeleteMembresiaController,
	*GetAllMembresiasController,
	*GetMembresiasByUsuarioIDController,
	*GetActiveMembresiaByUsuarioIDController,
	*GetExpiringMembresiasController,
	*RenewMembresiaController,
) {
	// Inicializar repositorio con la conexión a la base de datos
	repo := NewMySQLMembresiaRepository()

	// Inicializar casos de uso
	createUseCase := application.NewCreateMembresiaUseCase(repo)
	getByIdUseCase := application.NewGetMembresiaByIdUseCase(repo)
	updateUseCase := application.NewUpdateMembresiaUseCase(repo)
	deleteUseCase := application.NewDeleteMembresiaUseCase(repo)
	getAllUseCase := application.NewGetAllMembresiasUseCase(repo)
	getByUsuarioIDUseCase := application.NewGetMembresiasByUsuarioIDUseCase(repo)
	getActiveUseCase := application.NewGetActiveMembresiaByUsuarioIDUseCase(repo)
	getExpiringUseCase := application.NewGetExpiringMembresiasUseCase(repo)
	renewUseCase := application.NewRenewMembresiaUseCase(repo)

	// Inicializar controladores
	createController := NewCreateMembresiaController(createUseCase)
	getByIdController := NewGetMembresiaByIdController(getByIdUseCase)
	updateController := NewUpdateMembresiaController(updateUseCase)
	deleteController := NewDeleteMembresiaController(deleteUseCase)
	getAllController := NewGetAllMembresiasController(getAllUseCase)
	getByUsuarioIDController := NewGetMembresiasByUsuarioIDController(getByUsuarioIDUseCase)
	getActiveController := NewGetActiveMembresiaByUsuarioIDController(getActiveUseCase)
	getExpiringController := NewGetExpiringMembresiasController(getExpiringUseCase)
	renewController := NewRenewMembresiaController(renewUseCase)

	return createController, getByIdController, updateController,
		deleteController, getAllController, getByUsuarioIDController,
		getActiveController, getExpiringController, renewController
}

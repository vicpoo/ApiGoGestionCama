package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type AssignCamaToUsuarioUseCase struct {
	repo repositories.ICamaAfricana
}

func NewAssignCamaToUsuarioUseCase(repo repositories.ICamaAfricana) *AssignCamaToUsuarioUseCase {
	return &AssignCamaToUsuarioUseCase{repo: repo}
}

func (uc *AssignCamaToUsuarioUseCase) Run(camaID int32, usuarioID *int32) (*entities.CamaAfricana, error) {
	// Primero obtenemos la cama africana por su CamaID
	cama, err := uc.repo.GetByCamaID(camaID)
	if err != nil {
		return nil, err
	}

	// Actualizamos el usuarioID
	cama.UsuarioID = usuarioID

	// Guardamos los cambios
	err = uc.repo.Update(cama)
	if err != nil {
		return nil, err
	}

	return cama, nil
}

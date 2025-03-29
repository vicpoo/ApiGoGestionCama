package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"

type DeleteCamaAfricanaUseCase struct {
	repo repositories.ICamaAfricana
}

func NewDeleteCamaAfricanaUseCase(repo repositories.ICamaAfricana) *DeleteCamaAfricanaUseCase {
	return &DeleteCamaAfricanaUseCase{repo: repo}
}

func (uc *DeleteCamaAfricanaUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

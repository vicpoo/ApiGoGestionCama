package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"

type DeleteCamaUseCase struct {
	repo repositories.ICamaRepository
}

func NewDeleteCamaUseCase(repo repositories.ICamaRepository) *DeleteCamaUseCase {
	return &DeleteCamaUseCase{repo: repo}
}

func (uc *DeleteCamaUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

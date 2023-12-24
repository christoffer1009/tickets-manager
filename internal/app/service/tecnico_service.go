package service

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
)

type TecnicoService struct {
	TecnicoRepository *repository.TecnicoRepository
}

func NovoTecnicoService(tecnicoRepository *repository.TecnicoRepository) *TecnicoService {
	return &TecnicoService{
		TecnicoRepository: tecnicoRepository,
	}
}

func (s *TecnicoService) CriarTecnico(tecnicoDTO *models.TecnicoDTO) (*models.Tecnico, error) {
	novoTecnico := models.NovoTecnico(

		tecnicoDTO.Nome,
		tecnicoDTO.Email,
		models.NivelPrivilegio(tecnicoDTO.Nivel),
	)

	return s.TecnicoRepository.CriarTecnico(novoTecnico)
}

func (s *TecnicoService) ListarTodosTecnicos() ([]*models.Tecnico, error) {
	return s.TecnicoRepository.ListarTodosTecnicos()
}

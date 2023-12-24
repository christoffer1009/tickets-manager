package service

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/google/uuid"
)

type TecnicoService struct {
	TecnicoRepository *repository.TecnicoRepository
}

func NovoTecnicoService(tecnicoRepository *repository.TecnicoRepository) *TecnicoService {
	return &TecnicoService{
		TecnicoRepository: tecnicoRepository,
	}
}

func (s *TecnicoService) Criar(tecnicoDTO *models.TecnicoDTO) (*models.Tecnico, error) {
	novoTecnico := models.NovoTecnico(

		tecnicoDTO.Nome,
		tecnicoDTO.Email,
		tecnicoDTO.Nivel,
	)

	return s.TecnicoRepository.Criar(novoTecnico)
}

func (s *TecnicoService) EncontrarTodos() ([]*models.Tecnico, error) {
	return s.TecnicoRepository.EncontrarTodos()
}

func (s *TecnicoService) EncontrarPorID(tecnicoID uuid.UUID) (*models.Tecnico, error) {
	return s.TecnicoRepository.EncontrarPorID(tecnicoID)
}

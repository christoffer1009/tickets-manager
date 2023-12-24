package service

import (
	"fmt"

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

func (s *TecnicoService) Atualizar(tecnicoDTO *models.AtualizarTecnicoDTO) error {
	// Verifica se o tecnico existe antes de atualizar
	if !s.TecnicoRepository.Existe(tecnicoDTO.ID) {
		return fmt.Errorf("tecnico com ID %v não encontrado", tecnicoDTO.ID)
	}
	// Converte o DTO para o modelo
	tecnico := &models.Tecnico{
		ID:                  tecnicoDTO.ID,
		Nome:                tecnicoDTO.Nome,
		Email:               tecnicoDTO.Email,
		TicketsSolucionados: tecnicoDTO.TicketsSolucionados,
		SetorLotacao:        tecnicoDTO.SetorLotacao,
		Nivel:               tecnicoDTO.Nivel,
	}

	return s.TecnicoRepository.Atualizar(tecnico)
}

func (s *TecnicoService) Excluir(id uuid.UUID) error {
	// Verifica se o tecnico existe antes de deletar
	if !s.TecnicoRepository.Existe(id) {
		return fmt.Errorf("tecnico com ID %v não encontrado", id)
	}

	return s.TecnicoRepository.Excluir(id)
}

package service

import (
	"fmt"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/christoffer1009/tickets-manager/utils"
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
	hashedSenha, err := utils.GerarHashSenha(tecnicoDTO.Senha)
	if err != nil {
		return nil, err
	}

	novoTecnico := models.NovoTecnico(

		tecnicoDTO.Nome,
		tecnicoDTO.Email,
		hashedSenha,
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

func (s *TecnicoService) EncontrarPorEmail(email string) (*models.Tecnico, error) {
	return s.TecnicoRepository.EncontrarPorEmail(email)
}

func (s *TecnicoService) Atualizar(id uuid.UUID, tecnicoDTO *models.AtualizarTecnicoDTO) error {
	// Verifica se o tecnico existe antes de atualizar
	if !s.TecnicoRepository.Existe(id) {
		return fmt.Errorf("tecnico com ID %v não encontrado", id)
	}

	hashedSenha, err := utils.GerarHashSenha(tecnicoDTO.Senha)
	if err != nil {
		return err
	}
	tecnicoDTO.Senha = hashedSenha

	return s.TecnicoRepository.Atualizar(id, tecnicoDTO)
}

func (s *TecnicoService) Excluir(id uuid.UUID) error {
	// Verifica se o tecnico existe antes de deletar
	if !s.TecnicoRepository.Existe(id) {
		return fmt.Errorf("tecnico com ID %v não encontrado", id)
	}

	return s.TecnicoRepository.Excluir(id)
}

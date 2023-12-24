package service

import (
	"fmt"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/google/uuid"
)

type ClienteService struct {
	ClienteRepository *repository.ClienteRepository
}

func NovoClienteService(clienteRepository *repository.ClienteRepository) *ClienteService {
	return &ClienteService{
		ClienteRepository: clienteRepository,
	}
}

func (s *ClienteService) Criar(clienteDTO *models.ClienteDTO) (*models.Cliente, error) {

	novoCliente := models.NovoCliente(
		clienteDTO.Nome,
		clienteDTO.Email,
		clienteDTO.SetorLotacao,
	)

	return s.ClienteRepository.Criar(novoCliente)
}

func (s *ClienteService) EncontrarTodos() ([]*models.Cliente, error) {
	return s.ClienteRepository.EncontrarTodos()
}

func (s *ClienteService) EncontrarPorID(clienteID uuid.UUID) (*models.Cliente, error) {
	return s.ClienteRepository.EncontrarPorID(clienteID)
}

func (s *ClienteService) Atualizar(id uuid.UUID, clienteDTO *models.AtualizarClienteDTO) error {
	// Verifica se o cliente existe antes de atualizar
	if !s.ClienteRepository.Existe(id) {
		return fmt.Errorf("cliente com ID %v não encontrado", id)
	}
	// Converte o DTO para o modelo
	cliente := &models.AtualizarClienteDTO{
		// ID:           clienteDTO.ID,
		Nome:         clienteDTO.Nome,
		Email:        clienteDTO.Email,
		SetorLotacao: clienteDTO.SetorLotacao,
		TotalTickets: clienteDTO.TotalTickets,
	}

	return s.ClienteRepository.Atualizar(id, cliente)
}

func (s *ClienteService) Excluir(id uuid.UUID) error {
	// Verifica se o cliente existe antes de deletar
	if !s.ClienteRepository.Existe(id) {
		return fmt.Errorf("cliente com ID %v não encontrado", id)
	}

	return s.ClienteRepository.Excluir(id)
}

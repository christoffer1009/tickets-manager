package service

import (
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

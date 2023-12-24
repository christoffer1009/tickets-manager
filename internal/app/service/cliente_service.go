package service

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
)

type ClienteService struct {
	ClienteRepository *repository.ClienteRepository
}

func NovoClienteService(clienteRepository *repository.ClienteRepository) *ClienteService {
	return &ClienteService{
		ClienteRepository: clienteRepository,
	}
}

func (s *ClienteService) CriarCliente(clienteDTO *models.ClienteDTO) (*models.Cliente, error) {

	novoCliente := models.NovoCliente(
		clienteDTO.Nome,
		clienteDTO.Email,
		clienteDTO.SetorLotacao,
	)

	return s.ClienteRepository.CriarCliente(novoCliente)
}

func (s *ClienteService) ListarTodosClientes() ([]*models.Cliente, error) {
	return s.ClienteRepository.ListarTodosClientes()
}

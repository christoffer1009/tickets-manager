package service

import (
	"fmt"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/christoffer1009/tickets-manager/utils"
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
	hashedSenha, err := utils.GerarHashSenha(clienteDTO.Senha)
	if err != nil {
		return nil, err
	}

	novoCliente := models.NovoCliente(
		clienteDTO.Nome,
		clienteDTO.Email,
		hashedSenha,
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

	hashedSenha, err := utils.GerarHashSenha(clienteDTO.Senha)
	if err != nil {
		return err
	}

	clienteDTO.Senha = hashedSenha
	return s.ClienteRepository.Atualizar(id, clienteDTO)
}

func (s *ClienteService) Excluir(id uuid.UUID) error {
	// Verifica se o cliente existe antes de deletar
	if !s.ClienteRepository.Existe(id) {
		return fmt.Errorf("cliente com ID %v não encontrado", id)
	}

	return s.ClienteRepository.Excluir(id)
}

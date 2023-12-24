package service

import (
	"fmt"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/google/uuid"
)

type TicketService struct {
	TicketRepository  *repository.TicketRepository
	TecnicoRepository *repository.TecnicoRepository
	ClienteRepository *repository.ClienteRepository
}

func NovoTicketService(ticketRepository *repository.TicketRepository,
	tecnicoRepository *repository.TecnicoRepository,
	clienteRepository *repository.ClienteRepository) *TicketService {

	return &TicketService{
		TicketRepository:  ticketRepository,
		TecnicoRepository: tecnicoRepository,
		ClienteRepository: clienteRepository,
	}
}

func (s *TicketService) Criar(ticketDTO *models.TicketDTO) (*models.Ticket, error) {

	cliente, err := s.ClienteRepository.EncontrarPorID(ticketDTO.ClienteID)
	if err != nil {
		return nil, err
	}

	novoTicket := models.NovoTicket(
		ticketDTO.Titulo,
		ticketDTO.Descricao,
		nil,
		cliente,
	)

	return s.TicketRepository.Criar(novoTicket)
}

func (s *TicketService) EncontrarTodos() ([]*models.Ticket, error) {
	return s.TicketRepository.EncontrarTodos()
}

func (s *TicketService) EncontrarPorID(ticketID uuid.UUID) (*models.Ticket, error) {
	return s.TicketRepository.EncontrarPorID(ticketID)
}

// AtribuirTecnico atribui um técnico a um ticket
func (s *TicketService) AtribuirTecnico(ticketID uuid.UUID, tecnicoID uuid.UUID) error {
	return s.TicketRepository.AtribuirTecnico(ticketID, tecnicoID)
}

func (s *TicketService) Atualizar(ticketDTO *models.AtualizarTicketDTO) error {
	// Verifica se o ticket existe antes de atualizar
	if !s.TicketRepository.Existe(ticketDTO.ID) {
		return fmt.Errorf("ticket com ID %v não encontrado", ticketDTO.ID)
	}
	// Converte o DTO para o modelo
	ticket := &models.Ticket{
		ID:             ticketDTO.ID,
		Titulo:         ticketDTO.Titulo,
		Descricao:      ticketDTO.Descricao,
		Status:         ticketDTO.Status,
		DataAbertura:   ticketDTO.DataAbertura,
		DataFechamento: ticketDTO.DataFechamento,
		TecnicoID:      &ticketDTO.TecnicoID,
	}

	return s.TicketRepository.Atualizar(ticket)
}

func (s *TicketService) Excluir(id uuid.UUID) error {
	// Verifica se o ticket existe antes de deletar
	if !s.TicketRepository.Existe(id) {
		return fmt.Errorf("ticket com ID %v não encontrado", id)
	}

	return s.TicketRepository.Excluir(id)
}

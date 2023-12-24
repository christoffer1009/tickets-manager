package service

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
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

func (s *TicketService) CriarTicket(ticketDTO *models.TicketDTO) (*models.Ticket, error) {

	cliente, err := s.ClienteRepository.EncontraClientePorID(ticketDTO.ClienteID)
	if err != nil {
		return nil, err
	}

	novoTicket := models.NovoTicket(
		ticketDTO.Titulo,
		ticketDTO.Descricao,
		nil,
		cliente,
	)

	return s.TicketRepository.CriarTicket(novoTicket)
}

func (s *TicketService) ListarTodosTickets() ([]*models.Ticket, error) {
	return s.TicketRepository.ListarTodosTickets()
}

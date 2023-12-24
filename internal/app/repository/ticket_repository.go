package repository

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func NovoTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		DB: db,
	}
}

func (r *TicketRepository) Criar(ticket *models.Ticket) (*models.Ticket, error) {
	err := r.DB.Create(ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, err
}

func (r *TicketRepository) EncontrarTodos() ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	err := r.DB.Find(&tickets).Error
	return tickets, err
}

func (r *TicketRepository) EncontrarPorID(ticketID uuid.UUID) (*models.Ticket, error) {
	var ticket models.Ticket
	if err := r.DB.First(&ticket, ticketID).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

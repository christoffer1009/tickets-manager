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
	if err := r.DB.Preload("Cliente").Preload("Tecnico").First(&ticket, ticketID).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepository) AtribuirTecnico(ticketID uuid.UUID, tecnicoID uuid.UUID) error {
	var ticket models.Ticket
	if err := r.DB.First(&ticket, ticketID).Error; err != nil {
		return err
	}

	var tecnico models.Tecnico
	if err := r.DB.First(&tecnico, tecnicoID).Error; err != nil {
		return err
	}

	ticket.Tecnico = &tecnico
	ticket.TecnicoID = &tecnico.ID

	if err := r.DB.Save(&ticket).Error; err != nil {
		return err
	}

	return nil
}

func (r *TicketRepository) Existe(id uuid.UUID) bool {
	var count int64
	r.DB.Model(&models.Ticket{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func (r *TicketRepository) Atualizar(ticket *models.Ticket) error {
	// Verifica se o ticket existe antes de atualizar
	if !r.Existe(ticket.ID) {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Save(ticket).Error
}

func (r *TicketRepository) Excluir(id uuid.UUID) error {
	// Verifica se o ticket existe antes de deletar
	if !r.Existe(id) {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Delete(&models.Ticket{}, id).Error
}

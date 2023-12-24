package repository

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClienteRepository struct {
	DB *gorm.DB
}

func NovoClienteRepository(db *gorm.DB) *ClienteRepository {
	return &ClienteRepository{
		DB: db,
	}
}

func (r *ClienteRepository) Criar(cliente *models.Cliente) (*models.Cliente, error) {
	if err := r.DB.Create(cliente).Error; err != nil {
		return nil, err
	}

	return cliente, nil
}

func (r *ClienteRepository) EncontrarTodos() ([]*models.Cliente, error) {
	var clientes []*models.Cliente
	err := r.DB.Find(&clientes).Error
	return clientes, err
}

func (r *ClienteRepository) EncontrarPorID(clienteID uuid.UUID) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.DB.First(&cliente, "id = ?", clienteID).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

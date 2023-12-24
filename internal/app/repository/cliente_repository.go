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

func (r *ClienteRepository) Existe(id uuid.UUID) bool {
	var count int64
	r.DB.Model(&models.Cliente{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func (r *ClienteRepository) Atualizar(cliente *models.Cliente) error {
	// Verifica se o cliente existe antes de atualizar
	if !r.Existe(cliente.ID) {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Save(cliente).Error
}

func (r *ClienteRepository) Excluir(id uuid.UUID) error {
	// Verifica se o cliente existe antes de deletar
	if !r.Existe(id) {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Delete(&models.Cliente{}, id).Error
}

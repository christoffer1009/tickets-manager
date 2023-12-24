package repository

import (
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TecnicoRepository struct {
	DB *gorm.DB
}

func NovoTecnicoRepository(db *gorm.DB) *TecnicoRepository {
	return &TecnicoRepository{
		DB: db,
	}
}

func (r *TecnicoRepository) Criar(tecnico *models.Tecnico) (*models.Tecnico, error) {
	if err := r.DB.Create(tecnico).Error; err != nil {
		return nil, err
	}

	return tecnico, nil
}

func (r *TecnicoRepository) EncontrarTodos() ([]*models.Tecnico, error) {
	var tecnicos []*models.Tecnico
	err := r.DB.Find(&tecnicos).Error
	return tecnicos, err
}

func (r *TecnicoRepository) EncontrarPorID(tecnicoID uuid.UUID) (*models.Tecnico, error) {
	var tecnico models.Tecnico
	err := r.DB.First(&tecnico, "id = ?", tecnicoID).Error
	if err != nil {
		return nil, err
	}
	return &tecnico, nil
}

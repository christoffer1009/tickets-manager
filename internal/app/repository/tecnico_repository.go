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

func (r *TecnicoRepository) EncontrarPorEmail(email string) (*models.Tecnico, error) {
	var tecnico models.Tecnico
	err := r.DB.First(&tecnico, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &tecnico, nil
}

func (r *TecnicoRepository) Existe(id uuid.UUID) bool {
	var count int64
	r.DB.Model(&models.Tecnico{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func (r *TecnicoRepository) Atualizar(id uuid.UUID, tecnicoDTO *models.AtualizarTecnicoDTO) error {

	tecnico, err := r.EncontrarPorID(id)
	if err != nil {
		return err
	}

	// Atualiza apenas os campos fornecidos no DTO
	if tecnicoDTO.Nome != "" {
		tecnico.Nome = tecnicoDTO.Nome
	}

	if tecnicoDTO.Email != "" {
		tecnico.Email = tecnicoDTO.Email
	}

	if tecnicoDTO.Nivel != "" {
		tecnico.Nivel = tecnicoDTO.Nivel
	}

	if tecnicoDTO.TicketsSolucionados != 0 {
		tecnico.TicketsSolucionados = tecnicoDTO.TicketsSolucionados
	}

	if tecnicoDTO.SetorLotacao != "" {
		tecnico.SetorLotacao = tecnicoDTO.SetorLotacao
	}

	if tecnicoDTO.Senha != "" {
		tecnico.Senha = tecnicoDTO.Senha
	}

	return r.DB.Save(tecnico).Error
}

func (r *TecnicoRepository) Excluir(id uuid.UUID) error {
	// Verifica se o tecnico existe antes de deletar
	if !r.Existe(id) {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Delete(&models.Tecnico{}, id).Error
}

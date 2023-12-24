package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NivelPrivilegio representa os níveis de privilégio para um técnico.
type NivelPrivilegio string

const (
	Administrador NivelPrivilegio = "Administrador"
	Normal        NivelPrivilegio = "Normal"
)

// Tecnico representa um técnico responsável por resolver os tickets.
type Tecnico struct {
	ID                  uuid.UUID       `json:"id" gorm:"primaryKey"`
	Nome                string          `json:"nome"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel"`
	SetorLotacao        string          `json:"setor_lotacao"`
	Email               string          `json:"email" gorm:"unique"`
	Tickets             []*Ticket       `json:"tickets,omitempty" gorm:"foreignKey:TecnicoID"`
}

// NovoTecnico é um construtor para criar uma nova instância de Tecnico.
func NovoTecnico(nome, email string, nivel NivelPrivilegio) *Tecnico {
	return &Tecnico{
		ID:           uuid.New(),
		Nome:         nome,
		Nivel:        nivel,
		SetorLotacao: "TI",
		Email:        email,
	}
}

func (t *Tecnico) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

func (t *Tecnico) AdicionarTicketsSolucionados() int {
	t.TicketsSolucionados += 1
	return t.TicketsSolucionados
}

// ToString retorna uma string formatada com informações sobre o técnico.
func (t *Tecnico) ToString() string {
	return fmt.Sprintf("ID: %s\nNome: %s\nEmail: %s\nSetor de Locação: %s\nNivel : %s\nTickets Solucionados: %d",
		t.ID, t.Nome, t.Email, t.SetorLotacao, t.Nivel, t.TicketsSolucionados)
}

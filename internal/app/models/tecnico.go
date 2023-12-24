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
	TicketsSolucionados int             `json:"ticketsSolucionados"`
	Nivel               NivelPrivilegio `json:"nivel"`
	SetorLocacao        string          `json:"setorLocacao"`
	Email               string          `json:"email" gorm:"unique"`
	Tickets             []*Ticket       `json:"tickets,omitempty" gorm:"foreignKey:ResponsavelID"`
}

// NovoTecnico é um construtor para criar uma nova instância de Tecnico.
func NovoTecnico(nome, email string, nivel NivelPrivilegio) *Tecnico {
	return &Tecnico{
		ID:           uuid.New(),
		Nome:         nome,
		Nivel:        nivel,
		SetorLocacao: "TI",
		Email:        email,
	}
}

func (t *Tecnico) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

// SetID é um setter para o campo ID.
// func (t *Tecnico) SetID(id string) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.id = id
// }

// GetID é um getter para o campo ID.
// func (t *Tecnico) GetID() uuid.UUID {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.id
// }

// // SetNome é um setter para o campo Nome.
// func (t *Tecnico) SetNome(nome string) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.nome = nome
// }

// // GetNome é um getter para o campo Nome.
// func (t *Tecnico) GetNome() string {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.nome
// }

// func (t *Tecnico) SetEmail(email string) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.email = email
// }

// func (t *Tecnico) GetEmail() string {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.email
// }

// // SetNivel é um setter para o campo Nivel.
// func (t *Tecnico) SetNivel(nivel NivelPrivilegio) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.nivel = nivel
// }

// // GetNivel é um getter para o campo Nivel.
// func (t *Tecnico) GetNivel() NivelPrivilegio {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.nivel
// }

// // SetTicketsSolucionados é um setter para o campo ticketsSolucionados.
// func (t *Tecnico) setTicketsSolucionados(ticketsSolucionados int) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.ticketsSolucionados = ticketsSolucionados
// }

// // GetTicketsSolucionados é um getter para o campo ticketsSolucionados.
// func (t *Tecnico) GetTicketsSolucionados() int {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.ticketsSolucionados
// }

// func (t *Tecnico) GetSetorLocacao() string {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.setorLocacao
// }

func (t *Tecnico) AdicionarTicketsSolucionados() int {
	t.TicketsSolucionados += 1
	return t.TicketsSolucionados
}

// ToString retorna uma string formatada com informações sobre o técnico.
func (t *Tecnico) ToString() string {
	return fmt.Sprintf("ID: %s\nNome: %s\nEmail: %s\nSetor de Locação: %s\nNivel : %s\nTickets Solucionados: %d",
		t.ID, t.Nome, t.Email, t.SetorLocacao, t.Nivel, t.TicketsSolucionados)
}

package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// *Cliente representa um *cliente usuário do sistema.
type Cliente struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email" gorm:"unique"`
	SetorLocacao string    `json:"setorLocacao"`
	TotalTickets int       `json:"totalTickets"`
	Tickets      []*Ticket `json:"tickets,omitempty" gorm:"foreignKey:SolicitanteID"`
}

// NovoCliente cria uma nova instância de Cliente.
func NovoCliente(nome, email, setorLocacao string) *Cliente {
	return &Cliente{
		ID:           uuid.New(),
		Nome:         nome,
		Email:        email,
		SetorLocacao: setorLocacao,
		TotalTickets: 0,
	}
}

func (c *Cliente) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

// Implementação dos métodos da interface Usuario em Cliente

// Setter e Getters
// func (c *Cliente) SetID(id uuid.UUID) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.id = id
// }

// func (c *Cliente) GetID() uuid.UUID {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.id
// }

// func (c *Cliente) SetNome(nome string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.nome = nome
// }

// func (c *Cliente) GetNome() string {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.nome
// }

// func (c *Cliente) SetEmail(email string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.email = email
// }

// func (c *Cliente) GetEmail() string {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.email
// }

// func (c *Cliente) SetSetorLocacao(setorLocacao string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.setorLocacao = setorLocacao
// }

// func (c *Cliente) GetSetorLocacao() string {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.setorLocacao
// }

// func (c *Cliente) SetTotalTickets(totalTickets int) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.totalTickets = totalTickets
// }

// func (c *Cliente) GetTotalTickets() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.totalTickets
// }

func (c *Cliente) AdicionarTotalTickets() int {
	c.TotalTickets += 1
	return c.TotalTickets
}

func (c *Cliente) ToString() string {
	return fmt.Sprintf("ID: %s\nNome: %s\nEmail: %s\nSetor de Locação: %s\nTotal de Tickets: %d",
		c.ID, c.Nome, c.Email, c.SetorLocacao, c.TotalTickets)
}

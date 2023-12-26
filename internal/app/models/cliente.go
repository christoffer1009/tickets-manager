package models

import (
	"fmt"

	"github.com/google/uuid"
)

// *Cliente representa um *cliente usuário do sistema.
type Cliente struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email" gorm:"unique"`
	Senha        string    `json:"-"`
	SetorLotacao string    `json:"setor_lotacao"`
	TotalTickets int       `json:"total_tickets"`
	Tickets      []*Ticket `json:"tickets,omitempty" gorm:"foreignKey:ClienteID"`
}

// NovoCliente cria uma nova instância de Cliente.
func NovoCliente(nome, email, senha, setorLotacao string) *Cliente {
	return &Cliente{
		ID:           uuid.New(),
		Nome:         nome,
		Email:        email,
		Senha:        senha,
		SetorLotacao: setorLotacao,
		TotalTickets: 0,
	}
}

func (c *Cliente) AdicionarTotalTickets() int {
	c.TotalTickets += 1
	return c.TotalTickets
}

func (c *Cliente) ToString() string {
	return fmt.Sprintf("ID: %s\nNome: %s\nEmail: %s\nSetor de Locação: %s\nTotal de Tickets: %d",
		c.ID, c.Nome, c.Email, c.SetorLotacao, c.TotalTickets)
}

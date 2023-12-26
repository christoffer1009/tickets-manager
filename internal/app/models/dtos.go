package models

import (
	"time"

	"github.com/google/uuid"
)

// TicketDTO representa a estrutura de dados para transferência de dados relacionados a tickets.
type TicketDTO struct {
	ID             uuid.UUID  `json:"id"`
	Titulo         string     `json:"titulo" validate:"required,min=2,max=50"`
	Descricao      string     `json:"descricao" validate:"required,min=6,max=1000"`
	Status         Status     `json:"status" validade:"required,oneof=aberto atribuido solucionado fechado"`
	DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id" validate:"uuid"`
	ClienteID      uuid.UUID  `json:"cliente_id"  validate:"uuid"`
}

type AtualizarTicketDTO struct {
	// ID             uuid.UUID  `json:"id"`
	Titulo    string `json:"titulo" validate:"omitempty,min=2,max=50"`
	Descricao string `json:"descricao" validate:"omitempty,min=6,max=1000"`
	Status    Status `json:"status" validate:"omitempty,oneof=aberto atribuido solucionado fechado"`
	// DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id" validate:"omitempty,uuid"`
}

// TecnicoDTO representa a estrutura de dados para transferência de dados relacionados a técnicos.
type TecnicoDTO struct {
	ID                  uuid.UUID       `json:"id"`
	Nome                string          `json:"nome" validate:"required,min=3,max=50"`
	Email               string          `json:"email" validate:"required,email"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel" validate:"required,oneof=administrador normal"`
	SetorLotacao        string          `json:"setor_lotacao"  validate:"required,min=1,max=50"`
	Senha               string          `json:"senha" validate:"required,min=6,max=32"`
}

type AtualizarTecnicoDTO struct {
	// ID                  uuid.UUID       `json:"id"`
	Nome                string          `json:"nome" validate:"omitempty,min=3,max=50"`
	Email               string          `json:"email" validate:"omitempty,email"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel" validate:"omitempty,oneof=administrador normal"`
	SetorLotacao        string          `json:"setor_lotacao"  validate:"omitempty,min=1,max=50"`
	Senha               string          `json:"senha" validate:"omitempty,min=6,max=32"`
}

// ClienteDTO representa a estrutura de dados para transferência de dados relacionados a clientes.
type ClienteDTO struct {
	ID           uuid.UUID `json:"id"`
	Nome         string    `json:"nome" validate:"required,min=3,max=50"`
	Email        string    `json:"email" validate:"required,email"`
	SetorLotacao string    `json:"setor_lotacao"  validate:"omitempty,min=1,max=50"`
	TotalTickets int       `json:"total_tickets"`
	Senha        string    `json:"senha" validate:"required,min=6,max=32"`
}

type AtualizarClienteDTO struct {
	// ID           uuid.UUID `json:"id"`
	Nome         string `json:"nome" validate:"omitempty,min=3,max=50"`
	Email        string `json:"email" validate:"omitempty,email"`
	SetorLotacao string `json:"setor_lotacao"  validate:"omitempty,min=1,max=50"`
	TotalTickets int    `json:"total_tickets"`
	Senha        string `json:"senha" validate:"omitempty,min=6,max=32"`
}

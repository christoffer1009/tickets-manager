package models

import (
	"time"

	"github.com/google/uuid"
)

// TicketDTO representa a estrutura de dados para transferência de dados relacionados a tickets.
type TicketDTO struct {
	ID             uuid.UUID  `json:"id"`
	Titulo         string     `json:"titulo" validate:"required,min=6,max=50"`
	Descricao      string     `json:"descricao" validate:"required,min=6,max=1000"`
	Status         Status     `json:"status"`
	DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id"`
	ClienteID      uuid.UUID  `json:"cliente_id"`
}

type AtualizarTicketDTO struct {
	// ID             uuid.UUID  `json:"id"`
	Titulo    string `json:"titulo" validate:"required,min=6,max=50"`
	Descricao string `json:"descricao" validate:"required,min=6,max=1000"`
	Status    Status `json:"status"`
	// DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id"`
}

// TecnicoDTO representa a estrutura de dados para transferência de dados relacionados a técnicos.
type TecnicoDTO struct {
	ID                  uuid.UUID       `json:"id"`
	Nome                string          `json:"nome" validate:"required,min=6,max=50"`
	Email               string          `json:"email" validate:"required,email"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel" validate:"required,oneof=administrador normal"`
	SetorLotacao        string          `json:"setor_lotacao"`
}

type AtualizarTecnicoDTO struct {
	// ID                  uuid.UUID       `json:"id"`
	Nome                string          `json:"nome" validate:"required,min=6,max=50"`
	Email               string          `json:"email" validate:"required,email"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel" validate:"required,oneof=administrador normal"`
	SetorLotacao        string          `json:"setor_lotacao"`
}

// ClienteDTO representa a estrutura de dados para transferência de dados relacionados a clientes.
type ClienteDTO struct {
	ID           uuid.UUID `json:"id"`
	Nome         string    `json:"nome" validate:"required,min=6,max=50"`
	Email        string    `json:"email" validate:"required,email"`
	SetorLotacao string    `json:"setor_lotacao" validate:"required,max=50"`
	TotalTickets int       `json:"total_tickets"`
}

type AtualizarClienteDTO struct {
	// ID           uuid.UUID `json:"id"`
	Nome         string `json:"nome" validate:"required,min=6,max=50"`
	Email        string `json:"email" validate:"required,email"`
	SetorLotacao string `json:"setor_lotacao" validate:"required,max=50"`
	TotalTickets int    `json:"total_tickets"`
}

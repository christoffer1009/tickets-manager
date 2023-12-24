package models

import (
	"time"

	"github.com/google/uuid"
)

// TicketDTO representa a estrutura de dados para transferência de dados relacionados a tickets.
type TicketDTO struct {
	ID             uuid.UUID  `json:"id"`
	Titulo         string     `json:"titulo"`
	Descricao      string     `json:"descricao"`
	Status         Status     `json:"status"`
	DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id"`
	ClienteID      uuid.UUID  `json:"cliente_id"`
}

type AtualizarTicketDTO struct {
	// ID             uuid.UUID  `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    Status `json:"status"`
	// DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id"`
}

// TecnicoDTO representa a estrutura de dados para transferência de dados relacionados a técnicos.
type TecnicoDTO struct {
	ID                  uuid.UUID       `json:"id"`
	Nome                string          `json:"nome"`
	Email               string          `json:"email"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel"`
	SetorLotacao        string          `json:"setor_lotacao"`
}

type AtualizarTecnicoDTO struct {
	// ID                  uuid.UUID       `json:"id"`
	Nome                string          `json:"nome"`
	Email               string          `json:"email"`
	TicketsSolucionados int             `json:"tickets_solucionados"`
	Nivel               NivelPrivilegio `json:"nivel"`
	SetorLotacao        string          `json:"setor_lotacao"`
}

// ClienteDTO representa a estrutura de dados para transferência de dados relacionados a clientes.
type ClienteDTO struct {
	ID           uuid.UUID `json:"id"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email"`
	SetorLotacao string    `json:"setor_lotacao"`
	TotalTickets int       `json:"total_tickets"`
}

type AtualizarClienteDTO struct {
	// ID           uuid.UUID `json:"id"`
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	SetorLotacao string `json:"setor_lotacao"`
	TotalTickets int    `json:"total_tickets"`
}

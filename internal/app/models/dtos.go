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
	Status         string     `json:"status"`
	DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento"`
	TecnicoID      uuid.UUID  `json:"tecnico_id"`
	ClienteID      uuid.UUID  `json:"cliente_id"`
}

// TecnicoDTO representa a estrutura de dados para transferência de dados relacionados a técnicos.
type TecnicoDTO struct {
	ID                  uuid.UUID `json:"id"`
	Nome                string    `json:"nome"`
	Email               string    `json:"email"`
	TicketsSolucionados int       `json:"tickets_solucionados"`
	Nivel               string    `json:"nivel"`
}

// ClienteDTO representa a estrutura de dados para transferência de dados relacionados a clientes.
type ClienteDTO struct {
	ID           uuid.UUID `json:"id"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email"`
	SetorLocacao string    `json:"setor_locacao"`
	TotalTickets int       `json:"total_tickets"`
}

package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Status representa os estados de um ticket.
type Status string

const (
	aberto      Status = "aberto"
	atribuido   Status = "atribuído"
	solucionado Status = "solucionado"
	fechado     Status = "fechado"
)

// Ticket representa um ticket de suporte.
type Ticket struct {
	ID             uuid.UUID  `json:"id" gorm:"primaryKey"`
	Titulo         string     `json:"titulo"`
	Descricao      string     `json:"descricao"`
	Status         Status     `json:"status"`
	DataAbertura   time.Time  `json:"data_abertura"`
	DataFechamento *time.Time `json:"data_fechamento,omitempty"`
	Tecnico        *Tecnico   `json:"tecnico,omitempty" gorm:"foreignKey:TecnicoID"`
	Cliente        *Cliente   `json:"cliente,omitempty" gorm:"foreignKey:ClienteID"`
	TecnicoID      *uuid.UUID `json:"-" gorm:"type:uuid;index:idx_tecnico_id,tecnico_id"`
	ClienteID      uuid.UUID  `json:"-" gorm:"type:uuid;index:idx_cliente_id,cliente_id"`
}

// NovoTicket é um construtor para criar uma nova instância de Ticket.
func NovoTicket(titulo, descricao string, responsavel *Tecnico, cliente *Cliente) *Ticket {
	return &Ticket{
		ID:             uuid.New(),
		Titulo:         titulo,
		Descricao:      descricao,
		Status:         Status(aberto),
		DataAbertura:   time.Now(),
		DataFechamento: nil,
		Tecnico:        nil,
		Cliente:        cliente,
	}
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

func (t *Ticket) AtribuirTecnico(tecnico *Tecnico) {
	t.Tecnico = tecnico
	t.TecnicoID = &tecnico.ID
}

func (t *Ticket) ToString() string {

	dataFechamentoStr := "N/A"
	if t.DataFechamento != nil {
		dataFechamentoStr = t.DataFechamento.String()
	}

	return fmt.Sprintf("ID: %s\nTitulo: %s\nDescricao: %s\nStatus: %s\nData Abertura: %s\nData Fechamento: %s\nResponsavel: %s\nSolicitante: %s",
		t.ID, t.Titulo, t.Descricao, t.Status,
		t.DataAbertura.String(), dataFechamentoStr,
		t.Tecnico.Nome, t.Cliente.Nome)
}

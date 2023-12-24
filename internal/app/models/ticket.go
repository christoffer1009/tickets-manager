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
	Aberto      Status = "Aberto"
	Atribuido   Status = "Atribuído"
	Solucionado Status = "Solucionado"
	Fechado     Status = "Fechado"
)

// Ticket representa um ticket de suporte.
type Ticket struct {
	ID             uuid.UUID  `json:"id" gorm:"primaryKey"`
	Titulo         string     `json:"titulo"`
	Descricao      string     `json:"descricao"`
	Status         Status     `json:"status"`
	DataAbertura   time.Time  `json:"dataAbertura"`
	DataFechamento *time.Time `json:"dataFechamento,omitempty"`
	Responsavel    *Tecnico   `json:"responsavel,omitempty" gorm:"foreignKey:ResponsavelID"`
	Solicitante    *Cliente   `json:"solicitante,omitempty" gorm:"foreignKey:SolicitanteID"`
	ResponsavelID  *uuid.UUID `json:"-" gorm:"type:uuid;index:idx_responsavel_id,responsavel_id"`
	SolicitanteID  uuid.UUID  `json:"-" gorm:"type:uuid;index:idx_solicitante_id,solicitante_id"`
}

// NovoTicket é um construtor para criar uma nova instância de Ticket.
func NovoTicket(titulo, descricao string, responsavel *Tecnico, solicitante *Cliente) *Ticket {
	return &Ticket{
		ID:             uuid.New(),
		Titulo:         titulo,
		Descricao:      descricao,
		Status:         Status(Aberto),
		DataAbertura:   time.Now(),
		DataFechamento: nil,
		Responsavel:    nil,
		Solicitante:    solicitante,
	}
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

func (t *Ticket) AtribuirTecnico(tecnico *Tecnico) {
	t.Responsavel = tecnico
	t.ResponsavelID = &tecnico.ID
}

func (t *Ticket) ToString() string {

	dataFechamentoStr := "N/A"
	if t.DataFechamento != nil {
		dataFechamentoStr = t.DataFechamento.String()
	}

	return fmt.Sprintf("ID: %s\nTitulo: %s\nDescricao: %s\nStatus: %s\nData Abertura: %s\nData Fechamento: %s\nResponsavel: %s\nSolicitante: %s",
		t.ID, t.Titulo, t.Descricao, t.Status,
		t.DataAbertura.String(), dataFechamentoStr,
		t.Responsavel.Nome, t.Solicitante.Nome)
}

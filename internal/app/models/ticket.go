package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Ticket representa um ticket de suporte.
type Ticket struct {
	ID             uuid.UUID  `json:"id" gorm:"primaryKey"`
	Titulo         string     `json:"titulo"`
	Descricao      string     `json:"descricao"`
	Status         string     `json:"status"`
	DataAbertura   time.Time  `json:"dataAbertura"`
	DataFechamento *time.Time `json:"dataFechamento,omitempty"`
	Responsavel    *Tecnico   `json:"responsavel,omitempty" gorm:"foreignKey:ResponsavelID"`
	Solicitante    *Cliente   `json:"solicitante,omitempty" gorm:"foreignKey:SolicitanteID"`
	ResponsavelID  uuid.UUID  `json:"-" gorm:"type:uuid;index:idx_responsavel_id,responsavel_id"`
	SolicitanteID  uuid.UUID  `json:"-" gorm:"type:uuid;index:idx_solicitante_id,solicitante_id"`
}

// NovoTicket é um construtor para criar uma nova instância de Ticket.
func NovoTicket(titulo, descricao, status string, responsavel *Tecnico, solicitante *Cliente) *Ticket {
	return &Ticket{
		ID:             uuid.New(),
		Titulo:         titulo,
		Descricao:      descricao,
		Status:         status,
		DataAbertura:   time.Now(),
		DataFechamento: nil,
		Responsavel:    responsavel,
		Solicitante:    solicitante,
	}
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

// GetID é um getter para o campo ID.
// func (t *Ticket) GetID() uuid.UUID {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.id
// }

// // SetTitulo é um setter para o campo Titulo.
// func (t *Ticket) SetTitulo(titulo string) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.titulo = titulo
// }

// // GetTitulo é um getter para o campo Titulo.
// func (t *Ticket) GetTitulo() string {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.titulo
// }

// // SetDescricao é um setter para o campo Descricao.
// func (t *Ticket) SetDescricao(descricao string) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.descricao = descricao
// }

// // GetDescricao é um getter para o campo Descricao.
// func (t *Ticket) GetDescricao() string {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.descricao
// }

// // SetStatus é um setter para o campo Status.
// func (t *Ticket) SetStatus(status string) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.status = status
// }

// // GetStatus é um getter para o campo Status.
// func (t *Ticket) GetStatus() string {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.status
// }

// // GetDataAbertura é um getter para o campo DataFechamento.
// func (t *Ticket) GetDataAbertura() *time.Time {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return &t.dataAbertura
// }

// // SetDataFechamento é um setter para o campo DataFechamento.
// func (t *Ticket) SetDataFechamento(dataFechamento *time.Time) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.dataFechamento = dataFechamento
// }

// // GetDataFechamento é um getter para o campo DataFechamento.
// func (t *Ticket) GetDataFechamento() *time.Time {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.dataFechamento
// }

// // SetResponsavel é um setter para o campo Responsavel.
// func (t *Ticket) SetResponsavel(responsavel *Tecnico) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.responsavel = responsavel
// }

// // GetResponsavel é um getter para o campo Responsavel.
// func (t *Ticket) GetResponsavel() *Tecnico {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.responsavel
// }

// // SetSolicitante é um setter para o campo Solicitante.
// func (t *Ticket) SetSolicitante(solicitante *Cliente) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	t.solicitante = solicitante
// }

// // GetSolicitante é um getter para o campo Solicitante.
// func (t *Ticket) GetSolicitante() *Cliente {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.solicitante
// }

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

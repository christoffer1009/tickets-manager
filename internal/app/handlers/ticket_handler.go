package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	TicketService *service.TicketService
}

func NovoTicketHandler(ticketService *service.TicketService) *TicketHandler {
	return &TicketHandler{
		TicketService: ticketService,
	}
}

func (h *TicketHandler) CriarTicket(c *gin.Context) {
	var novoTicketDTO models.TicketDTO
	if err := c.BindJSON(&novoTicketDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novoTicket, err := h.TicketService.CriarTicket(&novoTicketDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	c.JSON(http.StatusCreated, novoTicket)

}

func (h *TicketHandler) ListarTodosTickets(c *gin.Context) {
	tickets, err := h.TicketService.ListarTodosTickets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao listar os tickets"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

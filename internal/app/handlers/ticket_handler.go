package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketHandler struct {
	TicketService *service.TicketService
}

func NovoTicketHandler(ticketService *service.TicketService) *TicketHandler {
	return &TicketHandler{
		TicketService: ticketService,
	}
}

func (h *TicketHandler) Criar(c *gin.Context) {
	var novoTicketDTO models.TicketDTO
	if err := c.BindJSON(&novoTicketDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novoTicket, err := h.TicketService.Criar(&novoTicketDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	c.JSON(http.StatusCreated, novoTicket)

}

func (h *TicketHandler) EncontrarTodos(c *gin.Context) {
	tickets, err := h.TicketService.EncontrarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao listar os tickets"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) EncontrarPorID(c *gin.Context) {
	// Obter o ID do ticket dos parâmetros da rota
	ticketIDStr := c.Param("id")

	// Validar o formato do ID
	ticketID, err := uuid.Parse(ticketIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	ticket, err := h.TicketService.EncontrarPorID(ticketID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Ticket não encontrado"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

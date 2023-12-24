package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClienteHandler struct {
	ClienteService *service.ClienteService
}

func NovoClienteHandler(clienteService *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{
		ClienteService: clienteService,
	}
}

func (h *ClienteHandler) Criar(c *gin.Context) {
	var novoClienteDTO models.ClienteDTO
	if err := c.BindJSON(&novoClienteDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novoCliente, err := h.ClienteService.Criar(&novoClienteDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	c.JSON(http.StatusCreated, novoCliente)
}

func (h *ClienteHandler) EncontrarTodos(c *gin.Context) {
	clientes, err := h.ClienteService.EncontrarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao listar os clientes"})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

func (h *ClienteHandler) EncontrarPorID(c *gin.Context) {
	// Obter o ID do cliente dos parâmetros da rota
	clienteIDStr := c.Param("id")

	// Validar o formato do ID
	clienteID, err := uuid.Parse(clienteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	cliente, err := h.ClienteService.EncontrarPorID(clienteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

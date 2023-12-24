package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	ClienteService *service.ClienteService
}

func NovoClienteHandler(clienteService *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{
		ClienteService: clienteService,
	}
}

func (h *ClienteHandler) CriarCliente(c *gin.Context) {
	var novoClienteDTO models.ClienteDTO
	if err := c.BindJSON(&novoClienteDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novoCliente, err := h.ClienteService.CriarCliente(&novoClienteDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	c.JSON(http.StatusCreated, novoCliente)
}

func (h *ClienteHandler) ListarTodosClientes(c *gin.Context) {
	clientes, err := h.ClienteService.ListarTodosClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao listar os clientes"})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

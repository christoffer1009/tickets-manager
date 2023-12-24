package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	novoClienteDTO.ID = novoCliente.ID
	novoClienteDTO.TotalTickets = novoCliente.TotalTickets

	c.JSON(http.StatusCreated, novoClienteDTO)
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
	clienteID, err := uuid.Parse(c.Param("id"))
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

func (h *ClienteHandler) Atualizar(c *gin.Context) {
	clienteID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	var clienteDTO models.AtualizarClienteDTO

	if err := c.ShouldBindJSON(&clienteDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Passa o ID para o DTO
	clienteDTO.ID = clienteID

	if err := h.ClienteService.Atualizar(&clienteDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente atualizado com sucesso"})
}

func (h *ClienteHandler) Excluir(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	if err := h.ClienteService.Excluir(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente deletado com sucesso"})
}

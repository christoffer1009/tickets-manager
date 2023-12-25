package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/custom_errors"
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var clienteDTO models.ClienteDTO

	if err := c.BindJSON(&clienteDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	validate := validator.New()

	if err := validate.Struct(clienteDTO); err != nil {
		var errosValidacao []custom_errors.ErroValidacao

		for _, fieldError := range err.(validator.ValidationErrors) {
			erroValidacao := custom_errors.ErroValidacao{
				Campo:    fieldError.Field(),
				Mensagem: fieldError.Tag(),
			}
			errosValidacao = append(errosValidacao, erroValidacao)
		}

		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro de validação", "detalhes": errosValidacao})
		return
	}

	novoCliente, err := h.ClienteService.Criar(&clienteDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	clienteDTO.ID = novoCliente.ID
	clienteDTO.TotalTickets = novoCliente.TotalTickets

	c.JSON(http.StatusCreated, clienteDTO)
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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	var clienteDTO models.AtualizarClienteDTO

	if err := c.ShouldBindJSON(&clienteDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ClienteService.Atualizar(id, &clienteDTO); err != nil {
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

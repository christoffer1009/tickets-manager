package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TecnicoHandler struct {
	TecnicoService *service.TecnicoService
}

func NovoTecnicoHandler(tecnicoService *service.TecnicoService) *TecnicoHandler {
	return &TecnicoHandler{
		TecnicoService: tecnicoService,
	}
}

func (h *TecnicoHandler) Criar(c *gin.Context) {
	var novoTecnicoDTO models.TecnicoDTO
	if err := c.BindJSON(&novoTecnicoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novoTecnico, err := h.TecnicoService.Criar(&novoTecnicoDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	c.JSON(http.StatusCreated, novoTecnico)
}

func (h *TecnicoHandler) EncontrarTodos(c *gin.Context) {
	tecnicos, err := h.TecnicoService.EncontrarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao listar os técnicos"})
		return
	}

	c.JSON(http.StatusOK, tecnicos)
}

func (h *TecnicoHandler) EncontrarPorID(c *gin.Context) {
	// Obter o ID do tecnico dos parâmetros da rota
	tecnicoIDStr := c.Param("id")

	// Validar o formato do ID
	tecnicoID, err := uuid.Parse(tecnicoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	tecnico, err := h.TecnicoService.EncontrarPorID(tecnicoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Tecnico não encontrado"})
		return
	}

	c.JSON(http.StatusOK, tecnico)
}

func (h *TecnicoHandler) Atualizar(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	var tecnicoDTO models.AtualizarTecnicoDTO

	if err := c.ShouldBindJSON(&tecnicoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.TecnicoService.Atualizar(id, &tecnicoDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tecnico atualizado com sucesso"})
}

func (h *TecnicoHandler) Excluir(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido de ID"})
		return
	}

	if err := h.TecnicoService.Excluir(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Técnico não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Técnico deletado com sucesso"})
}

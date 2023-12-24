package handlers

import (
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/gin-gonic/gin"
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

func (h *TecnicoHandler) ListarTodos(c *gin.Context) {
	tecnicos, err := h.TecnicoService.EncontrarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao listar os técnicos"})
		return
	}

	c.JSON(http.StatusOK, tecnicos)
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/christoffer1009/tickets-manager/internal/app/validators"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	var tecnicoDTO models.TecnicoDTO
	if err := c.BindJSON(&tecnicoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	errValidacao := validators.ValidarCriarTecnico(tecnicoDTO)
	if errValidacao != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro de validação", "detalhes": errValidacao})
		return
	}

	tecnico, err := h.TecnicoService.Criar(&tecnicoDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar o ticket"})
		return
	}

	c.JSON(http.StatusCreated, tecnico)
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

	errValidacao := validators.ValidarAtualizarTecnico(tecnicoDTO)
	if errValidacao != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro de validação", "detalhes": errValidacao})
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

func (h *TecnicoHandler) Protegido(c *gin.Context) {
	// Extrai informações do token, se necessário
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(401, gin.H{"error": "Token claims not found"})
		return
	}

	id := claims.(jwt.MapClaims)["id"].(string)
	nome := claims.(jwt.MapClaims)["nome"].(string)
	email := claims.(jwt.MapClaims)["email"].(string)

	c.JSON(200, gin.H{"message": fmt.Sprintf("Olá, %s! email: %s e id: %s.", email, nome, id)})
}

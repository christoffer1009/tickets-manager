package handlers

import (
	"net/http"
	"time"

	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/christoffer1009/tickets-manager/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	ClienteService *service.ClienteService
	TecnicoService *service.TecnicoService
}

func NovoLoginHandler(clienteService *service.ClienteService, tecnicoService *service.TecnicoService) *LoginHandler {
	return &LoginHandler{
		ClienteService: clienteService,
		TecnicoService: tecnicoService,
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	type loginInput struct {
		Email string
		Senha string
	}

	var input loginInput
	var tipo string

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	cliente, _ := h.ClienteService.EncontrarPorEmail(input.Email)
	if cliente != nil {
		tipo = "cliente"
	}

	tecnico, _ := h.TecnicoService.EncontrarPorEmail(input.Email)
	if tecnico != nil {
		tipo = "tecnico"
	}

	if cliente == nil && tecnico == nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuario não encontrado"})
	}

	secretkey, err := utils.GetSecretkey()
	if err != nil {
		return
	}

	switch tipo {
	case "cliente":

		if bcrypt.CompareHashAndPassword([]byte(cliente.Senha), []byte(input.Senha)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "credenciais erradas"})
		} else {
			// Cria o token JWT
			token, err := h.gerarToken(secretkey, cliente.ID, cliente.Nome, cliente.Email)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
		}

	case "tecnico":
		if bcrypt.CompareHashAndPassword([]byte(tecnico.Senha), []byte(input.Senha)) != nil {
			c.JSON(http.StatusNotFound, gin.H{"erro": "erro de hash"})
		} else {
			// Cria o token JWT
			token, err := h.gerarToken(secretkey, tecnico.ID, tecnico.Nome, tecnico.Email)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
		}

	default:
		return
	}

}

func (h *LoginHandler) gerarToken(secretkey []byte, id uuid.UUID, nome, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"nome":  nome,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

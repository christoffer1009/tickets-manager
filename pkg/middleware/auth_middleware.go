package middleware

import (
	"github.com/christoffer1009/tickets-manager/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware é um middleware que verifica a validade do token
func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	secretkey, err := utils.GetSecretkey()
	if err != nil {
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Converte as reivindicações para jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{"error": "Invalid token claims type"})
		c.Abort()
		return
	}

	c.Set("claims", claims)

	// Se o token for válido, permita o acesso à rota protegida
	c.Next()
}

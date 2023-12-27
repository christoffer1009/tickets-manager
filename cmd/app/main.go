package main

import (
	"github.com/christoffer1009/tickets-manager/internal/app/handlers"
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
	"github.com/christoffer1009/tickets-manager/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configurar conexão com o banco de dados (substitua esta parte conforme necessário)
	db, err := models.SetupDB()
	if err != nil {
		panic("Erro ao conectar ao banco de dados")
	}

	// Configurar repositórios
	ticketRepository := repository.NovoTicketRepository(db)
	tecnicoRepository := repository.NovoTecnicoRepository(db)
	clienteRepository := repository.NovoClienteRepository(db)

	// Configurar serviços
	ticketService := service.NovoTicketService(ticketRepository, tecnicoRepository, clienteRepository)
	tecnicoService := service.NovoTecnicoService(tecnicoRepository)
	clienteService := service.NovoClienteService(clienteRepository)

	// Configurar handlers
	ticketHandler := handlers.NovoTicketHandler(ticketService)
	tecnicoHandler := handlers.NovoTecnicoHandler(tecnicoService)
	clienteHandler := handlers.NovoClienteHandler(clienteService)
	loginHandler := handlers.NovoLoginHandler(clienteService, tecnicoService)

	// Configurar rotas
	r.POST("/tickets", ticketHandler.Criar)
	r.GET("/tickets", ticketHandler.EncontrarTodos)
	r.GET("/tickets/:id", ticketHandler.EncontrarPorID)
	r.PUT("/tickets/:id/atribuir/:tecnicoID", ticketHandler.AtribuirTecnico)
	r.PUT("/tickets/:id", ticketHandler.Atualizar)
	r.DELETE("/tickets/:id", ticketHandler.Excluir)

	r.POST("/tecnicos", tecnicoHandler.Criar)
	r.GET("/tecnicos", tecnicoHandler.EncontrarTodos)
	r.GET("/tecnicos/:id", tecnicoHandler.EncontrarPorID)
	r.PUT("/tecnicos/:id", tecnicoHandler.Atualizar)
	r.DELETE("/tecnicos/:id", tecnicoHandler.Excluir)

	r.POST("/clientes", clienteHandler.Criar)
	r.GET("/clientes", clienteHandler.EncontrarTodos)
	r.GET("/clientes/:id", clienteHandler.EncontrarPorID)
	r.PUT("/clientes/:id", clienteHandler.Atualizar)
	r.DELETE("/clientes/:id", clienteHandler.Excluir)

	r.POST("/login", loginHandler.Login)

	r.GET("/protegida", middleware.AuthMiddleware, tecnicoHandler.Protegido)

	// Iniciar servidor
	r.Run(":8080")
}

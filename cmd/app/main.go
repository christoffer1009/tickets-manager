package main

import (
	"github.com/christoffer1009/tickets-manager/internal/app/handlers"
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/christoffer1009/tickets-manager/internal/app/repository"
	"github.com/christoffer1009/tickets-manager/internal/app/service"
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

	// Configurar rotas
	r.POST("/tickets", ticketHandler.CriarTicket)
	r.GET("/tickets", ticketHandler.ListarTodosTickets)

	r.POST("/tecnicos", tecnicoHandler.CriarTecnico)
	r.GET("/tecnicos", tecnicoHandler.ListarTodosTecnicos)

	r.POST("/clientes", clienteHandler.CriarCliente)
	r.GET("/clientes", clienteHandler.ListarTodosClientes)

	// Iniciar servidor
	r.Run(":8080")
}
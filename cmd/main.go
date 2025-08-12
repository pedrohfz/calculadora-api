package main

import (
	"calculadora-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// A função 'main' é o ponto de entrada da aplicação
// Aqui iniciamos o servidor HTTP e registramos as rotas disponíveis.
func main() {
	server := gin.Default()
	routes.InitRoutes(&server.RouterGroup)

	if err := server.Run(":5050"); err != nil {
		log.Fatal(err)
	}
}

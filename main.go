package main

import (
	"calculadora-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// A função 'main' é o ponto de entrada da aplicação
// Aqui iniciamos o servidor HTTP e registramos as rotas disponíveis.
func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":5050"); err != nil {
		log.Fatal(err)
	}
}

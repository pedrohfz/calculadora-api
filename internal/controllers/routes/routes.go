package routes

import (
	"calculadora-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// InitRoutes define e registra todas as rotas disponíveis da API.
// Cada rota está associada a uma operação matemática específica.
func InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/adicao", handlers.Somar)
	rg.POST("/subtracao", handlers.Subtrair)
	rg.POST("/multiplicacao", handlers.Multiplicar)
	rg.POST("/divisao", handlers.Dividir)
	rg.POST("/modulo", handlers.Modulo)
}

package routes

import (
	"calculadora-api/handlers"

	"github.com/gin-gonic/gin"
)

// InitRouter define e registra todas as rotas disponíveis da API.
// Cada rota está associada a uma operação matemática específica.
func InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/adicao", handlers.Adicao)
	rg.POST("/subtracao", handlers.Subtracao)
	rg.POST("/multiplicacao", handlers.Multiplicacao)
	rg.POST("/divisao", handlers.Divisao)
	rg.POST("/modulo", handlers.Modulo)
}

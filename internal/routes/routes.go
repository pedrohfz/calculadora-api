package routes

import (
	"calculadora-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// InitRoutes define e registra todas as rotas disponíveis da API.
// Cada rota está associada a uma operação matemática específica.
func InitRoutes(r *gin.RouterGroup) {
	r.POST("/adicao", handlers.Somar)
	r.POST("/subtracao", handlers.Subtrair)
	r.POST("/multiplicacao", handlers.Multiplicar)
	r.POST("/divisao", handlers.Dividir)
	r.POST("/modulo", handlers.Modulo)
	r.GET("/operacoes", handlers.Operacoes)
}

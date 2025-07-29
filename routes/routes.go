package routes

import (
	"calculadora-api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/adicao", handlers.Adicao)
	rg.POST("/subtracao", handlers.Subtracao)
	rg.POST("/multiplicacao", handlers.Multiplicacao)
	rg.POST("/divisao", handlers.Divisao)
	rg.POST("/modulo", handlers.Modulo)
}

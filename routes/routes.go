package routes

import (
	"calculadora-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/adicao", controller.Adicao)
	rg.POST("/subtracao", controller.Subtracao)
	rg.POST("/multiplicacao", controller.Multiplicacao)
	rg.POST("/divisao", controller.Divisao)
	rg.POST("/modulo", controller.Modulo)
}

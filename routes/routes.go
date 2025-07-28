package routes

import (
	"calculadora-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/calcular", controller.Calcular)
}

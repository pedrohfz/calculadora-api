package handlers

import (
	"calculadora-api/controllers"
	"calculadora-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Adicao(c *gin.Context) {
	HandleOperation(c, "+")
}

func Subtracao(c *gin.Context) {
	HandleOperation(c, "-")
}

func Multiplicacao(c *gin.Context) {
	HandleOperation(c, "*")
}

func Divisao(c *gin.Context) {
	HandleOperation(c, "/")
}

func Modulo(c *gin.Context) {
	HandleOperation(c, "%")
}

func HandleOperation(c *gin.Context, operador string) {
	var req models.OperationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.OperationResponse{
			Mensagem: "Dados inválidos: " + err.Error(),
			Erro:     true,
		})
		return
	}

	resp, erro := controllers.CalcularOperacao(req.Num1, req.Num2, operador)
	if erro {
		c.JSON(http.StatusBadRequest, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

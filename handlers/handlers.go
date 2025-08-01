package handlers

import (
	"calculadora-api/controllers"
	"calculadora-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Adicao lida com a rota POST /adicao
func Adicao(c *gin.Context) {
	HandleOperation(c, "+")
}

// Subtracao lida com a rota POST /subtracao
func Subtracao(c *gin.Context) {
	HandleOperation(c, "-")
}

// Multiplicacao lida com a rota POST /multiplicacao
func Multiplicacao(c *gin.Context) {
	HandleOperation(c, "*")
}

// Divisao lida com a rota POST /divisao
func Divisao(c *gin.Context) {
	HandleOperation(c, "/")
}

// Modulo lida com a rota POST /modulo
func Modulo(c *gin.Context) {
	HandleOperation(c, "%")
}

// HandleOperation centraliza o tratamento da requisição e resposta HTTP.
// Faz o bind dos dados, chama o controller e retorna no JSON apropriado.
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

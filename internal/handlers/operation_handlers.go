package handlers

import (
	"calculadora-api/internal/models"
	"calculadora-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleMathOperation centraliza o tratamento da requisição e resposta HTTP.
// Faz o bind dos dados, chama o controller e retorna no JSON apropriado.
func HandleMathOperation(c *gin.Context, operador string) {
	var req models.OperationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.OperationResponse{
			Mensagem: "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CalcularOperacao(req.Num1, req.Num2, operador)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// HandleListOperations... 
func HandleListOperations(c *gin.Context) {
	// TODO: Função para retornar todas as operações feitas.
}
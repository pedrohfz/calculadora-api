package handlers

import (
	"calculadora-api/internal/services"
	"calculadora-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Somar lida com a rota POST /adicao
func Somar(c *gin.Context) {
	HandleOperation(c, "+")
}

// Subtrair lida com a rota POST /subtracao
func Subtrair(c *gin.Context) {
	HandleOperation(c, "-")
}

// Multiplicar lida com a rota POST /multiplicacao
func Multiplicar(c *gin.Context) {
	HandleOperation(c, "*")
}

// Dividir lida com a rota POST /divisao
func Dividir(c *gin.Context) {
	HandleOperation(c, "/")
}

// Modulo lida com a rota POST /modulo
func Modulo(c *gin.Context) {
	HandleOperation(c, "%")
}

func Operacoes(c *gin.Context) {
	// TODO: Função para listar todas as operações já feitas na execução da API.
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

	resp, err := controllers.CalcularOperacao(req.Num1, req.Num2, operador)
	if err {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	
	c.JSON(http.StatusOK, resp)
}

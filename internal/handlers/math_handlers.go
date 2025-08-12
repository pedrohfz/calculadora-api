package handlers

import "github.com/gin-gonic/gin"

// Somar lida com a rota POST /adicao
func Somar(c *gin.Context) {
	HandleMathOperation(c, "+")
}

// Subtrair lida com a rota POST /subtracao
func Subtrair(c *gin.Context) {
	HandleMathOperation(c, "-")
}

// Multiplicar lida com a rota POST /multiplicacao
func Multiplicar(c *gin.Context) {
	HandleMathOperation(c, "*")
}

// Dividir lida com a rota POST /divisao
func Dividir(c *gin.Context) {
	HandleMathOperation(c, "/")
}

// Modulo lida com a rota POST /modulo
func Modulo(c *gin.Context) {
	HandleMathOperation(c, "%")
}

// Operacoes lida com a rota GET /operacoes
func Operacoes(c *gin.Context) {
	// TODO: Função para listar todas as operações já feitas na execução da API.
}

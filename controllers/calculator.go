package controllers

import (
	"calculadora-api/models"
	"fmt"
	"math"
)

var (
	resultado float64
	mensagem  string
	erro      bool
)

// CalcularOperacao executa a lógica da operação matemática com base no operador fornecido.
// Retorna uma OperationResponse contendo o resultado ou mensagem de erro, além de um bool indicando falha ou não.
func CalcularOperacao(num1, num2 float64, operador string) (models.OperationResponse, bool) {
	switch operador {
	case "+":
		resultado = num1 + num2
		mensagem = fmt.Sprintf("A soma de %.2f e %.2f é %.2f", num1, num2, resultado)
	case "-":
		resultado = num1 - num2
		mensagem = fmt.Sprintf("A subtração de %.2f e %.2f é %.2f", num1, num2, resultado)
	case "*":
		resultado = num1 * num2
		mensagem = fmt.Sprintf("A multiplicação de %.2f e %.2f é %.2f", num1, num2, resultado)
	case "/":
		if num2 == 0 {
			mensagem = "Não é possível dividir um número por zero."
			erro = true
		} else {
			resultado = num1 / num2
			mensagem = fmt.Sprintf("A divisão de %.2f e %.2f é %.2f", num1, num2, resultado)
		}
	case "%":
		if num2 == 0 {
			mensagem = "Não é possível calcular o módulo com divisor zero."
			erro = true
		} else {
			resultado = math.Mod(num1, num2)
			mensagem = fmt.Sprintf("O módulo de %.2f e %.2f é %.2f", num1, num2, resultado)
		}
	default:
		mensagem = "Operação desconhecida."
		erro = true
	}

	return models.OperationResponse{
		Resultado: resultado,
		Mensagem:  mensagem,
		Erro:      erro,
	}, erro
}

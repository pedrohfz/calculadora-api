package services

import (
	"calculadora-api/internal/models"
	"errors"
	"fmt"
	"math"
)

// CalcularOperacao executa a lógica da operação matemática com base no operador fornecido.
// Retorna uma OperationResponse contendo o resultado ou mensagem de erro, além de um bool indicando falha ou não.
func CalcularOperacao(num1, num2 float64, operador string) (models.OperationResponse, error) {

	// TODO: Armazenar todos os resultados numa slice e listar ao chamar a rota GET /operacoes.

	var resultado float64
	var mensagem string

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
			return models.OperationResponse{}, errors.New("Não é possível dividir um número por zero!")
		}
		resultado = num1 / num2
		mensagem = fmt.Sprintf("A divisão de %.2f e %.2f é %.2f", num1, num2, resultado)
	case "%":
		if num2 == 0 {
			return models.OperationResponse{}, errors.New("Não é possível dividir um número por zero!")
		}
		resultado = math.Mod(num1, num2)
		mensagem = fmt.Sprintf("O módulo de %.2f e %.2f é %.2f", num1, num2, resultado)
	default:
		return models.OperationResponse{}, errors.New("Operação inválida!")
	}

	return models.OperationResponse{Resultado: resultado, Mensagem: mensagem}, nil
}

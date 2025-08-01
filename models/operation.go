package models

// OperationRequest representa o corpo da requisição enviada pelo cliente.
// Contém dois números sobre os quais será realizada a operação matemática.
type OperationRequest struct {
	Num1 float64 `json:"num1" binding:"required"`
	Num2 float64 `json:"num2" binding:"required"`
}

// OperationResponse representa a resposta retornada pela API.
// Contém o resultado de uma operação, uma mensagem explicativa e o status de erro.
type OperationResponse struct {
	Resultado float64 `json:"resultado,omitempty"`
	Mensagem  string  `json:"mensagem,omitempty"`
	Erro      bool    `json:"erro,omitempty"`
}

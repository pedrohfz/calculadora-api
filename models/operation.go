package models

type OperationRequest struct {
	Num1 float64 `json:"num1" binding:"required"`
	Num2 float64 `json:"num2" binding:"required"`
}

type OperationResponse struct {
	Resultado float64 `json:"resultado,omitempty"`
	Mensagem  string  `json:"mensagem,omitempty"`
	Erro      bool    `json:"erro"`
}

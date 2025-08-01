# 🧮 Calculadora API em Go

    Esta é uma API RESTful desenvolvida em Golang com o framework Gin, que permite realizar operações matemáticas básicas por meio de requisições HTTP. A API está estruturada de forma modular, com separação clara entre rotas, handlers, controllers e modelos.

## 🚀 Funcionalidades

    ✅ Adição (POST /adicao)
    ✅ Subtração (POST /subtracao)
    ✅ Multiplicação (POST /multiplicacao)
    ✅ Divisão (POST /divisao)
    ✅ Módulo (POST /modulo)

## 📦 Tecnologias Utilizadas

    - Go 1.24.5
    - Framework Gin Gonic

## 📂 Estrutura do Projeto

    calculadora-api/        :
    ├── go.mod              : Arquivo de gerenciamento de dependências e módulos do Go
    ├── main.go             : Ponto de entrada da aplicação. Inicializa o servidor HTTP
    ├── controllers/        : Contém a lógica de cálculo para cada operação matemática
    │   └── calculator.go   : Implementa as funções de soma, subtração, multiplicação, divisão e módulo
    ├── handlers/           : Define os manipuladores das requisições HTTP
    │   └── calculator.go   : Trata as entradas do usuário e responde com o resultado
    ├── models/             : Define os modelos utilizados nas requisições e respostas
    │   └── operation.go    : Contém as structs de request e response
    ├── routes/             : Define e registra todas as rotas da aplicação
    │   └── routes.go       : Mapeia os endpoints para os handlers correspondentes

## 📥 Exemplo de Requisição

    POST /divisao

    # Request:

        {
            "num1": 10,
            "num2": 2
        }

    # Response:

        {
            "resultado": 5,
            "mensagem": "A divisão de 10.00 por 2.00 é 5.00",
            "erro": false
        }

## ▶️ Como executar localmente

    1. Clone o repositório:
        git clone https://github.com/pedrohfz/calculadora-api.git
        cd calculadora-api
    
    2. Instale as dependências:
        go mod tidy

    3. Rode a aplicação:
        go run main.go

    4. Acesse a API (Postman/Insonmia):
        http://localhost:5050/

## 📎 Autor

    Desenvolvido por Pedro Henrique Leite como uma forma de estudar Go.

## 📚 Licença

    Este projeto está licenciado sob a MIT License.
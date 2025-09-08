package main

import (
    "fmt"
    "log"
)

func main() {
    // Simulação de dados sensíveis
    username := "usuario123"
    password := "senhaSuperSecreta"
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    cpf := "123.456.789-00"
    cnpj := "12.345.678/0001-99"
    creditCard := "4111 1111 1111 1111"
    email := "usuario@email.com"
    apiKey := "AKIAIOSFODNN7EXAMPLE"
    secret := "mySuperSecret"
    refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...refresh"
    phone := "+55 11 91234-5678"
    address := "Rua Exemplo, 123, São Paulo"
    sessionID := "sess-1234567890"
    pin := "1234"

    // Exemplo de logs inseguros: NÃO FAÇA ISSO!
    log.Printf("Usuário: %s fez login com a senha: %s", username, password)
    log.Printf("Token de autenticação: %s", token)
    log.Printf("CPF do usuário: %s", cpf)
    log.Printf("CNPJ da empresa: %s", cnpj)
    log.Printf("Cartão de crédito informado: %s", creditCard)
    log.Printf("E-mail cadastrado: %s", email)
    log.Printf("API Key utilizada: %s", apiKey)
    log.Printf("Segredo de autenticação: %s", secret)
    log.Printf("Refresh Token: %s", refreshToken)
    log.Printf("Telefone do usuário: %s", phone)
    log.Printf("Endereço do usuário: %s", address)
    log.Printf("Session ID: %s", sessionID)
    log.Printf("PIN informado: %s", pin)

    // Exemplo de log seguro (NÃO exponha dados sensíveis!)
    fmt.Println("Usuário fez login com sucesso.")
}

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

    // Exemplo de log inseguro: expondo dados sensíveis
    log.Printf("Usuário: %s fez login com a senha: %s", username, password)
    log.Printf("Token de autenticação: %s", token)

    // Exemplo de log seguro (NÃO exponha dados sensíveis!)
    fmt.Println("Usuário fez login com sucesso.")
}

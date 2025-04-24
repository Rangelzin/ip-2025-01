package main

import (
    "fmt"
)

func main() {
    var numero int

    // Lê o número na base 10
    fmt.Print("Digite um valor na base 10: ")
    fmt.Scan(&numero)

    // Verifica se o número é positivo
    if numero <= 0 {
        fmt.Println("Por favor, digite um número inteiro positivo.")
        return
    }

    // Calcula o equivalente binário
    binario := ""
    for numero > 0 {
        binario = fmt.Sprintf("%d", numero%2) + binario
        numero = numero / 2
    }

    // Imprime o número na base 2
    fmt.Println("O número na base 2:", binario)
}
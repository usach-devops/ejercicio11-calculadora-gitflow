package main

import (
	"fmt"
	"os"
	"strings"
)

//USO go run calculadora.go 2/3
//USO go run calculadora.go 2*3
//USO go run calculadora.go 2+3
//USO go run calculadora.go 2-3
//test CMP
func main() {
	fmt.Println("Calculadora 10%")

	inputValue := os.Args[1:]

	str1 := strings.Join(inputValue, " ")

	if strings.Contains(str1, "+") {
		fmt.Println("Implementar Suma")
	}
	if strings.Contains(str1, "-") {
		fmt.Println("Implementar Resta")
	}
	if strings.Contains(str1, "*") {
		fmt.Println("Implementar Multiplicación")
	}
	if strings.Contains(str1, "/") {
		fmt.Println("Implementar División")
	}
}

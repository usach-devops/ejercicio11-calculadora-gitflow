package main

import (
	"fmt"
	"os"
	"strings"
)

//USO go run main.go 2/3
//USO go run main.go 2*3
//USO go run main.go 2+3
//USO go run maim.go 2-3
//test CMP
//test CMP2
func main() {
	fmt.Println("Calculadora 10%")

	inputValue := os.Args[1:]

	str1 := strings.Join(inputValue, " ")

	fmt.Println(str1)

	if strings.Contains(str1, "+") {
		fmt.Println("Implementar Suma")
	}
	if strings.Contains(str1, "-") {
		fmt.Println("Implementar Resta")
	}
	if strings.Contains(str1, "*") {
		res := strings.Split(str1, "*")
		fmt.Println(res[0])
		fmt.Println(getMultiplicar(3, 3))
	}
	if strings.Contains(str1, "/") {
		fmt.Println("Implementar División")
	}
}

func getMultiplicar(value1 int, value2 int) int {
	return value1 * value2
}

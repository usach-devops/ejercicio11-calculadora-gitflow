package main

import (
	"fmt"
	"os"
	"strconv"
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

	if strings.Contains(str1, "+") {
		res := strings.Split(str1, "+")

		value1, err := strconv.Atoi(res[0])
		value2, err1 := strconv.Atoi(res[1])

		if err != nil && err1 != nil {
			fmt.Println(getSumar(value1, value2))
		}
		fmt.Println(getSumar(value1, value2))
	}
	if strings.Contains(str1, "-") {
		fmt.Println("Implementar Resta")
	}
	if strings.Contains(str1, "*") {
		res := strings.Split(str1, "*")

		value1, err := strconv.Atoi(res[0])
		value2, err1 := strconv.Atoi(res[1])

		if err != nil && err1 != nil {
			fmt.Println(getMultiplicar(value1, value2))
		}
		fmt.Println(getMultiplicar(value1, value2))
	}

	if strings.Contains(str1, "/") {
		res := strings.Split(str1, "/")

		value1, err := strconv.Atoi(res[0])
		value2, err1 := strconv.Atoi(res[1])

		if err != nil && err1 != nil {
			fmt.Println(getDivsion(value1, value2))
		}
		fmt.Println(getDivsion(value1, value2))

	}
}

func getMultiplicar(value1 int, value2 int) int {
	return value1 * value2
}

func getSumar(value1 int, value2 int) int {
	return value1 + value2
}

func getDivsion(value1 int, value2 int) int {
	return value1 / value2

}

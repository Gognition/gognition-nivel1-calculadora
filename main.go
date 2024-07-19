package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: go run main.go <número1> <operación> <número2>")
		fmt.Println("Operaciones disponibles: +, -, *, /")
		os.Exit(1)
	}
	n1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error: El primer número no es válido")
		os.Exit(1)
	}

	operacion := os.Args[2]

	n2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error: El segundo número no es válido")
		os.Exit(1)
	}

	var resultado float64
	switch operacion {
	case "+":
		resultado = sumar(n1, n2)
	case "-":
		resultado = restar(n1, n2)
	case "*":
		resultado = multiplicar(n1, n2)
	case "/":
		resultado = dividir(n1, n2)
	default:
		fmt.Println("Operación no válida. ¿Estás inventando matemáticas nuevas?")
		os.Exit(1)
	}

	fmt.Printf("El resultado de %.2f %s %.2f es: %.2f\n", n1, operacion, n2, resultado)

}

func sumar(n1, n2 float64) float64 {
	return n1 + n2
}

func restar(n1, n2 float64) float64 {
	return n1 - n2
}

func multiplicar(n1, n2 float64) float64 {
	return n1 * n2
}

func dividir(n1, n2 float64) float64 {
	if n2 == 0 {
		fmt.Println("dividir zero")
		return 0
	}
	return n1 / n2
}

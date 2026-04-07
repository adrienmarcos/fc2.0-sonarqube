package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Soma:", sum(1, 2))
	fmt.Println("Resto:", mod(10, 3))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// --- CÓDIGO NOVO ABAIXO (Feito para reprovar no Gate) ---

func mod(a int, b int) int {
	return a % b
}

func power(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func isOdd(a int) bool {
	if a%2 != 0 {
		return true
	}
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func sayHello(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name
}
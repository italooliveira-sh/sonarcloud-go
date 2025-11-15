package main

import "fmt"

func main() {
	fmt.Println(soma(2, 2))
}

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func multiplicacao(a int, b int) int {
	return a * b
}

func divisao(a int, b int) int {
	return a / b
}

func isPar(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

package main

import "testing"

func TestSoma(t *testing.T) {
	result := soma(2, 3)
	resultEspected := 5

	if result != resultEspected {
		t.Errorf("Resultado da soma é inválido! Valor esperado é %d e o resultado foi %d.", resultEspected, result)
	}

}

func TestSubtracao(t *testing.T) {
	result := subtracao(5, 3)
	resultEspected := 2

	if result != resultEspected {
		t.Errorf("Resultado da subtração é inválido! Valor esperado é %d e o resultado foi %d.", resultEspected, result)
	}

}

func TestSubtracaoWhenParamALessThanB(t *testing.T) {
	result := subtracao(3, 5)
	resultEspected := 2

	if result != resultEspected {
		t.Errorf("Resultado da subtração é inválido! Valor esperado é %d e o resultado foi %d.", resultEspected, result)
	}

}

func TestMultiplicacao(t *testing.T) {
	result := multiplicacao(5, 3)
	resultEspected := 15

	if result != resultEspected {
		t.Errorf("Resultado da multiplicação é inválido! Valor esperado é %d e o resultado foi %d.", resultEspected, result)
	}

}

func TestDivisao(t *testing.T) {
	result := divisao(6, 3)
	resultEspected := 2

	if result != resultEspected {
		t.Errorf("Resultado da divisão é inválido! Valor esperado é %d e o resultado foi %d.", resultEspected, result)
	}

}

func TestIsPar(t *testing.T) {
	result := isPar(4)
	resultEspected := true

	if result != resultEspected {
		t.Errorf("Resultado é inválido! Valor esperaro é %t e o resultado foi %t", resultEspected, result)
	}

}

func TestIsNotPar(t *testing.T) {
	result := isPar(5)
	resultEspected := false

	if result != resultEspected {
		t.Errorf("Resultado é inválido! Valor esperaro é %t e o resultado foi %t", resultEspected, result)
	}

}

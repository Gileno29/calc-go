package main

import "testing"

func DeveSumarCorrectamente(t *testing.T) {
	teste := soma(3, 2)
	resultado := 5

	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

func DeveSumarIncorrectamente(t *testing.T) {

	teste := soma(3.0, 2.0)
	resultado := 7

	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

func DeveSubtrairCorretamente(t *testing.T) {

	teste := subtracao(3.0, 2.0)
	resultado := 1
	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}
}

func DeveSubtrairIncorretamente(t *testing.T) {

	teste := subtracao(3.0, 2.0)
	resultado := 10
	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

func DeveMultiplicarCorretamente(t *testing.T) {

	teste := multiplicacao(3.0, 2.0)
	resultado := 6
	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

func DeveMultiplicarIncorretamente(t *testing.T) {

	teste := multiplicacao(3.0, 2.0)
	resultado := 10
	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

func DeveDividirCorretamente(t *testing.T) {

	teste := multiplicacao(10.0, 2.0)
	resultado := 5
	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

func DeveDividirIncorretamente(t *testing.T) {

	teste := multiplicacao(10.0, 2.0)
	resultado := 20
	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}

}

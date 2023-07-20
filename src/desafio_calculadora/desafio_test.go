package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	test := sum(4, 5, 6)

	result := 15

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}
func TestShouldSumIncorrect(t *testing.T) {
	test := sum(4, 5, 6)

	result := 14

	if test == result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	test := mult(4, 5, 6)

	result := 120

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}

}
func TestShouldMultIncorrect(t *testing.T) {
	test := mult(4, 5, 6)

	result := 60

	if test == result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	test := sub(5, 3)

	result := 2

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}
func TestShouldSubIncorrect(t *testing.T) {
	test := sub(5, 3)

	result := 4

	if test == result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	test := div(4, 2)

	result := 2.0

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}
func TestShouldDivIncorrect(t *testing.T) {
	test := div(4, 2)

	result := 3.0

	if test == result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

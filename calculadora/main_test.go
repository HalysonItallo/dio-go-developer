package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1)

	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	teste := multiplica(6, 2, 3)

	resultado := 36

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultiplyIncorrect(t *testing.T) {
	teste := multiplica(6, 2, 3)
	resultado := 24
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	teste := subtrai(6, 2)

	resultado := 4

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubtractIncorrect(t *testing.T) {
	teste := multiplica(6, 2, 3)
	resultado := 2
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	teste := divide(6, 2)

	resultado := 3

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	teste := divide(6, 2, 3)
	resultado := 2
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}

	totalMultiply := multiply(3, 3)

	if totalMultiply != 9 {
		t.Errorf("Resultado da multiplicacao é inválido: Resultado %d. Esperado: %d", totalMultiply, 9)
	}

}

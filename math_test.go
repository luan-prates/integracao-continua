package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(10, 23)

	if total != 33 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
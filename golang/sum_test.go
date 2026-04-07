package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)
	if result != 5 {
		t.Error("The result should be 5")
	}
}

// Testando apenas uma das funções novas para ter uma cobertura mínima,
// mas que vai ficar muito abaixo dos 90% exigidos.
func TestMod(t *testing.T) {
	result := mod(10, 3)
	if result != 1 {
		t.Error("The result should be 1")
	}
}
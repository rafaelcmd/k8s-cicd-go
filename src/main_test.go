package main

import "testing"

func TestRaiz(t *testing.T) {
	raiz := raiz(0.001)
	if raiz != "Code.education Rocks!" {
		t.Error("O texto não é: ", raiz)
	}
}

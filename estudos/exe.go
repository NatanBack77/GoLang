package main

import (
	"fmt"
)

func main() {

	// --- EXERCÍCIO: SOMA DE TODOS OS ELEMENTOS DE UM SLICE ---

	lista := []int{1, 2, 3, 4}

	// Acumulador: começa em 0 e vai somando cada elemento
	soma := 0

	// Percorre o slice pelo índice e acumula os valores
	// soma += lista[i] equivale a: soma = soma + lista[i]
	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}

	// Resultado esperado: 1 + 2 + 3 + 4 = 10
	fmt.Println("Soma:", soma)

	// ALTERNATIVA IDIOMÁTICA com "range":
	// for _, v := range lista { soma += v }
	// O "_" descarta o índice quando não precisamos dele.
	// "range" é preferido em Go quando não é necessário manipular o índice.
}

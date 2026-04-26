package main

import "fmt"

func main() {

	// --- SLICES: CRIAÇÃO E ITERAÇÃO ---
	// Slice é a estrutura de lista/array dinâmica mais usada em Go.
	// Diferente de arrays (tamanho fixo), slices crescem dinamicamente.

	// FORMA 1: literal — cria slice já com valores iniciais
	// Sintaxe: []Tipo{val1, val2, ...}
	lista2 := []int{1, 2, 3, 4}

	// FORMA 2 (comentada): make([]Tipo, tamanho)
	// make cria um slice com tamanho definido, todos os valores = zero do tipo
	// lista := make([]int, 1000) → slice de 1000 ints, todos zero

	// Iteração com for clássico usando índice
	for i := 0; i < len(lista2); i++ {
		fmt.Printf("%v\n", lista2[i]) // %v imprime o "valor" padrão do tipo
	}

	// DIFERENÇA ENTRE ARRAY E SLICE em Go:
	//   [4]int{1,2,3,4}  → array: tamanho FIXO, faz parte do tipo
	//   []int{1,2,3,4}   → slice: tamanho DINÂMICO, pode crescer com append()

	// append() adiciona elementos ao slice (comentado acima):
	// lista = append(lista, 6) → adiciona o valor 6 ao final

	// listaString := []string{"golang", "javacript", "php"}
	// listaString = append(listaString, "python") → slice de strings também funciona
}

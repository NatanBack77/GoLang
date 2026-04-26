package main

import "fmt"

func main() {

	// --- SLICES: FATIAMENTO (SLICING) ---
	// Go permite criar "sub-slices" de um slice existente usando a notação [inicio:fim].
	// IMPORTANTE: o sub-slice NÃO copia os dados — ele aponta para o mesmo array
	// subjacente. Modificar um sub-slice pode alterar o slice original!

	listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3}
	//                  0   1  2  3  4  5  6  7   ← índices

	// [:3] → do início até o índice 3 (EXCLUSIVO): posições 0, 1, 2
	// Resultado: [2, 10, 9]
	segundaLista := listaToda[:3]

	// [3:] → do índice 3 (INCLUSIVO) até o fim
	// Resultado: [4, 5, 8, 1, 3]
	terceiraLista := listaToda[3:]

	// [len-1:] → pega apenas o último elemento
	// len(listaToda) = 8 → listaToda[7:] = [3]
	ultimoItem := listaToda[len(listaToda)-1:]

	fmt.Println(ultimoItem)     // → [3]
	fmt.Println(terceiraLista)  // → [4 5 8 1 3]
	fmt.Println(segundaLista)   // → [2 10 9]

	// REGRA GERAL: slice[inicio:fim]
	//   inicio → índice inclusivo (padrão: 0)
	//   fim    → índice exclusivo (padrão: len(slice))
	//
	// Exemplos:
	//   [2:5]  → posições 2, 3, 4
	//   [:]    → cópia superficial do slice inteiro (mesmo array base)
	//   [:len] → equivalente a [:]
}

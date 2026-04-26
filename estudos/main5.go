package main

import "fmt"

func main() {

	// --- SLICES: MAKE, APPEND E FILTRAGEM ---

	listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3}

	// make([]int, 0): cria um slice vazio de int, com comprimento 0.
	// Diferença entre make e declaração nil:
	//   make([]int, 0) → slice vazio, mas inicializado (não é nil)
	//   var i []int    → slice nil (valor zero do tipo slice)
	// Ambos aceitam append(), mas nil slice imprime "[]" enquanto make imprime "[]" também.
	listaMenor := make([]int, 0)

	var i []int        // slice nil — valor zero, sem alocação de memória
	fmt.Println(i)     // → []
	fmt.Println(listaMenor) // → []

	// FILTRAGEM: percorre listaToda e copia apenas os valores menores que 5
	for i := 0; i < len(listaToda); i++ {
		if listaToda[i] < 5 {
			// append(slice, elemento) → retorna um NOVO slice com o elemento adicionado.
			// Em Go, append pode realocar memória internamente se necessário.
			// IMPORTANTE: o retorno deve ser reatribuído ao mesmo slice.
			listaMenor = append(listaMenor, listaToda[i])
		}
	}

	fmt.Println(listaMenor) // → [2 4 1 3] (elementos < 5, em ordem de aparição)

	// ATENÇÃO: a variável "i" do for interno (i := 0) é diferente da variável
	// "var i []int" declarada acima. O i do for cria um novo escopo local e
	// "shadowing" acontece — o i do for oculta o i de slice dentro do loop.
}

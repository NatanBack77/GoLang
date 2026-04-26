// VERSÃO COMENTADA (abordagem com fatiamento):
// Divide o slice em duas metades e soma cada metade separadamente.
// Esta abordagem usa slicing para separar os dados antes de processar.

// package main
// import "fmt"
// func main(){
// 	slice := []int{2,8,3,10,5,4,7,9,1}
// 	somaPrimeira := 0
// 	somaSegunda := 0
// 	slice1:= slice[:5]    // primeiros 5 elementos: [2,8,3,10,5]
// 	slice2:= slice[5:]    // do índice 5 em diante: [4,7,9,1]
// 	fmt.Println(len(slice1)) // → 5
// 	fmt.Println(len(slice2)) // → 4
// 	for i := 0; i < len(slice1); i++ { somaPrimeira += slice1[i] }
// 	for i := 0; i < len(slice2); i++ { somaSegunda += slice2[i] }
// 	fmt.Println("soma1", somaPrimeira) // → 28
// 	fmt.Println("soma2", somaSegunda)  // → 21
// }

package main

import "fmt"

func main() {

	// --- EXERCÍCIO: SOMA SEPARADA POR CONDIÇÃO ---
	// Percorre o slice UMA VEZ e distribui cada elemento em uma das duas somas
	// baseado em uma condição. Mais eficiente que a versão com fatiamento acima
	// porque faz apenas 1 passagem pelo slice (O(n)) e não aloca memória extra.

	slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	somaPrimeira := 0 // acumula elementos <= 5
	somaSegunda := 0  // acumula elementos > 5

	for i := 0; i < len(slice); i++ {
		if slice[i] <= 5 {
			somaPrimeira += slice[i] // 2 + 3 + 5 + 4 + 1 = 15
		} else {
			somaSegunda += slice[i] // 8 + 10 + 7 + 9 = 34
		}
	}

	fmt.Println("soma1", somaPrimeira) // → 15
	fmt.Println("soma2", somaSegunda)  // → 34

	// COMPARAÇÃO DAS DUAS ABORDAGENS:
	// Versão fatiada:    divide primeiro, soma depois → 2 loops, sub-slices na memória
	// Versão condicional: soma enquanto percorre → 1 loop, sem alocação extra
}

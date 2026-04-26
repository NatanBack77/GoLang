package main

import "fmt"

func main() {

	// --- FOR CLÁSSICO COM 3 PARTES + LOOPS ANINHADOS ---
	// Sintaxe: for inicialização; condição; pós-execução { ... }
	// Equivalente direto ao "for (int i=0; i<10; i++)" de C/Java.

	// Loop externo: controla o número base (1 a 10)
	for numbase := 1; numbase <= 10; numbase++ {

		// Loop interno: controla o multiplicador (1 a 10)
		// A cada iteração do loop externo, o loop interno roda 10 vezes completo.
		// Total de iterações: 10 × 10 = 100
		for multiplicador := 1; multiplicador <= 10; multiplicador++ {
			fmt.Println(numbase, " x ", multiplicador, " = ", numbase*multiplicador)
		}
	}

	// LOOPS ANINHADOS: o loop de dentro tem acesso às variáveis do de fora.
	// Cuidado com a performance: a complexidade cresce multiplicativamente.
	// 2 loops aninhados com N iterações cada = O(N²) operações.
}

package main

import "fmt"

func main() {

	// --- EXERCÍCIO: ACESSAR O ÚLTIMO ELEMENTO DE UM SLICE ---

	var listaToda = []int{12, 2, 3, 3, 3}

	// Técnica para pegar o último elemento usando slicing:
	// len(listaToda) retorna o tamanho (5 neste caso)
	// listaToda[len-1:] → do índice 4 até o fim → retorna um SLICE com o último elemento
	// Resultado: [3]
	segundaLIsta := listaToda[len(listaToda)-1:]

	fmt.Print(segundaLIsta) // → [3]

	// ALTERNATIVA para acessar o último elemento como VALOR (não slice):
	// ultimoValor := listaToda[len(listaToda)-1]
	// Isso retorna int (3), não []int ([3])
	//
	// Em Go 1.21+ existe a sintaxe: listaToda[len(listaToda)-1]
	// Futuramente virá suporte a índice negativo como Python (listaToda[-1]),
	// mas ainda não existe na linguagem padrão.
}

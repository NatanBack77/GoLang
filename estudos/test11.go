package main

import "fmt"

func main() {

	// --- MAP VAZIO COM MAKE ---
	// Demonstra a criação de um map sem valores iniciais.
	// make(map[TipoChave]TipoValor) inicializa o map e o deixa pronto para uso.

	m := make(map[string]int)

	// Um map recém-criado com make está vazio, mas NÃO é nil.
	// Diferença crucial:
	//   var m map[string]int → nil map (PANIC se tentar escrever)
	//   m := make(map[string]int) → map vazio e seguro para leitura e escrita

	fmt.Println(m) // → map[]

	// Após criar com make, podemos adicionar entradas normalmente:
	// m["chave"] = valor
}

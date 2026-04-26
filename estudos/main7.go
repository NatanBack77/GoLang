package main

import "fmt"

func main() {

	// --- MAPS: DICIONÁRIOS EM GO ---
	// Map é a estrutura chave-valor do Go (equivalente a dict em Python, objeto em JS).
	// Sintaxe: map[TipoChave]TipoValor

	// FORMA 1: literal — cria map já com valores
	m1 := map[string]int{"sp": 100, "rj": 200, "mg": 300}
	fmt.Println(m1) // A ordem de impressão NÃO é garantida em maps

	// FORMA 2: make — cria map vazio e adiciona depois
	// make(map[TipoChave]TipoValor) é necessário para mapas criados sem valores iniciais.
	// SEM make, o map seria nil e qualquer escrita causaria panic em runtime.
	m2 := make(map[string]int)
	m2["sp"] = 100
	m2["rj"] = 200
	m2["mg"] = 300
	fmt.Println(m2)

	// LEITURA SIMPLES: retorna o valor ou o "zero value" do tipo se a chave não existir
	// Para int, zero value = 0. Por isso não dá para distinguir "chave inexistente" de "valor 0".
	valor, _ := m2["sp"] // o "_" descarta o segundo retorno (bool ok)
	fmt.Printf("Valor: %v\n", valor)

	// LEITURA COM VERIFICAÇÃO DE EXISTÊNCIA (padrão idiomático em Go):
	// Map retorna dois valores: (valor, ok)
	// ok == true  → chave existe
	// ok == false → chave não existe (valor será o zero value do tipo)
	if valor, ok := m2["sp"]; ok {
		fmt.Printf("Valor: %v\n", valor)
	}

	// DELETE: remove uma chave do map.
	// Se a chave não existir, não gera erro (operação é segura).
	delete(m2, "sp")

	// ITERAÇÃO COM RANGE:
	// "range" sobre map retorna (chave, valor) em cada iteração.
	// A ORDEM DAS ITERAÇÕES É ALEATÓRIA — Go randomiza propositalmente para
	// evitar que código dependa de uma ordem específica.
	for chave, valor := range m2 {
		fmt.Println("cidade", chave, "habitantes", valor)
	}
}

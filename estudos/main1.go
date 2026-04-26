package main

import (
	"fmt"
)

func main() {

	// --- ESTRUTURA CONDICIONAL: if / else if / else ---
	// Go usa if/else igual a outras linguagens, mas SEM parênteses na condição.
	// As chaves {} são OBRIGATÓRIAS, mesmo que o bloco tenha uma única linha.

	salario := 1000.00
	tipo := "CLT"

	// Operadores lógicos em Go:
	//   &&  → AND (ambas as condições devem ser verdadeiras)
	//   ||  → OR  (pelo menos uma condição deve ser verdadeira)
	//   !   → NOT (inverte o valor booleano)

	if salario < 1000.00 && tipo == "CLT" {
		// Desconto de 10% para CLT com salário abaixo de 1000
		// salario -= (salario * 0.1) equivale a: salario = salario - (salario * 0.1)
		salario -= (salario * 0.1)
	} else if salario <= 2000.00 {
		// Se o primeiro if for falso, testa essa condição:
		// salário entre 1000 e 2000 → desconto de 15%
		salario -= (salario * 0.15)
	} else {
		// Nenhuma condição anterior foi satisfeita → salário > 2000
		// Desconto de 20%
		salario -= (salario * 0.2)
	}

	// ATENÇÃO: neste exemplo, salario == 1000.00 e tipo == "CLT"
	// A condição "salario < 1000.00" é FALSA (1000 não é MENOR que 1000)
	// Então cai no else if (salario <= 2000.00) → desconto de 15%
	// Resultado: 1000 - 150 = 850
	fmt.Println("Salário líquido:", salario)
}

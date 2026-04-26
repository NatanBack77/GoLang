package main

import (
	"fmt"
)

func main() {

	// --- LOOP FOR COM CONTINUE + ÍNDICE DE STRING ---
	// Go tem apenas UMA estrutura de loop: o "for".
	// Ele substitui while, do-while e for clássico de outras linguagens.

	texto := "palavra"

	// len() retorna o número de BYTES da string (não de caracteres Unicode).
	// Para strings ASCII (letras latinas sem acento), bytes == caracteres.
	fmt.Println("quantidade de caracteres:", len(texto)) // → 7

	tamanho := len(texto)
	i := 0

	// FOR estilo "while": só tem a condição, sem inicialização nem pós-execução.
	// Equivalente em outras linguagens: while (i < tamanho) { ... }
	for i < tamanho {

		// texto[i] retorna o BYTE (uint8) na posição i.
		// string(texto[i]) converte esse byte de volta para string legível.
		if string(texto[i]) == "r" {
			// CONTINUE: pula para a próxima iteração do loop SEM executar
			// o restante do bloco. O problema aqui é que "i++" vem DEPOIS,
			// então se entrar no continue sem incrementar i, vira loop infinito!
			// Isso é um BUG clássico de estudo: o continue impede o i++ de rodar.
			continue
		}

		fmt.Println(string(texto[i]))
		i++ // incrementa só se NÃO caiu no continue
	}

	// RESUMO DO BUG: quando texto[i] == "r", o continue é executado,
	// i nunca é incrementado e o loop trava infinitamente na letra "r".
	// A correção seria incrementar i ANTES do continue.
}

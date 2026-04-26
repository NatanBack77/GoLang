package main

// Pacote "strconv" fornece funções para converter tipos primitivos de/para string.
// É necessário quando precisamos converter entre tipos incompatíveis como string <-> int.
import (
	"fmt"
	"strconv"
)

func main() {

	// --- CONVERSÃO EXPLÍCITA DE TIPOS (Type Casting) ---
	// Em Go, conversões de tipo NUNCA são implícitas. Você sempre precisa
	// converter explicitamente usando o nome do tipo como função: TipoDestino(valor)

	a := 10 // Go infere o tipo como "int"

	// Convertendo int -> float32
	// Sem o float32(...), o compilador recusa e retorna erro.
	// Diferença de outras linguagens (JS, Python) onde isso acontece automaticamente.
	var numero2 float32
	numero2 = float32(a)
	fmt.Printf("%T\n", numero2) // %T imprime o TIPO da variável → "float32"

	// --- CONVERSÃO DE STRING PARA NÚMERO (strconv.ParseInt) ---
	b := "20" // b é do tipo string

	// ParseInt(string, base, bitSize) → converte string para int64
	// base 10 = decimal, bitSize 64 = int64
	// Retorna (int64, error). O "_" descarta o erro (não recomendado em produção)
	i, _ := strconv.ParseInt(b, 10, 64)
	fmt.Printf("%T\n", &i) // &i = ponteiro para i → imprime "*int64"

	// --- CONVERSÃO DE NÚMERO PARA STRING (strconv.Itoa) ---
	c := 20

	// Itoa = "Integer to ASCII" — converte int para string de forma simples
	// Alternativa: strconv.FormatInt(int64(c), 10) para mais controle
	s := strconv.Itoa(c)
	fmt.Printf("%T\n", s) // → "string"

	// --- DECLARAÇÃO EXPLÍCITA DE TIPO ---
	// Aqui declaramos "d" com tipo explícito "int" em vez de deixar o Go inferir.
	// Ambas as formas abaixo são equivalentes:
	//   var d int = 10   (declaração explícita)
	//   d := 10          (inferência de tipo — shorthand)
	var d int = 10
	fmt.Printf("%T\n", d) // → "int"
}

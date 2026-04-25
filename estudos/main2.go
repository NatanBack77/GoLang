package main

import (
	"fmt"
)
func main() {
	texto := "palavra"
	fmt.Println("quantidade de caracteres:", len(texto))
	tamanho := len(texto)
    i := 0
	for i < tamanho {
		if string(texto[i]) == "r" {	
			continue
		}
		fmt.Println(string(texto[i]))
		i ++
		
	}
}
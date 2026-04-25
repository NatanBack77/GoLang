package main

import "fmt"

func main(){
	m1 := map[string]int{"sp": 100, "rj": 200, "mg": 300}
	fmt.Println(m1)
	
	m2 := make(map[string]int)
	m2["sp"] = 100
	m2["rj"] = 200
	m2["mg"] = 300
	fmt.Println(m2)

	valor, _ := m2["sp"]
	fmt.Printf("Valor: %v\n", valor)


	if valor, ok := m2["sp"]; ok {
		fmt.Printf("Valor: %v\n", valor)
	}
	delete(m2, "sp")

	for chave, valor := range m2 {
		fmt.Println("cidade", chave, "habitantes", valor)
	}




}
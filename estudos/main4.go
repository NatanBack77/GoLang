package main

import "fmt"

func main(){
	lista2 := []int{1,2,3,4}
	// lista := make([]int, 1000)

    for i := 0; i < len(lista2); i++ {
	fmt.Printf("%v\n", lista2[i])
    // lista = append(lista, i)
	
}

	// fmt.Printf("%v\n", lista)
	// lista = append(lista, 6)
	// listaString := []string{"golang", "javacript", "php"}
	// listaString = append(listaString, "python")
	// fmt.Printf("%t\n", lista)
	// fmt.Println("lista pos1\n", lista[4])
	// fmt.Println("tamanho da lista", len(lista))
	// fmt.Printf("%t\n", listaString)
	// fmt.Println("lista pos1\n", listaString[3])
	// fmt.Println("tamanho da lista", len(listaString))
}
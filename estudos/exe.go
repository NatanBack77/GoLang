package main

import (
	"fmt"
)

func main() {
	lista := []int{1,2,3,4}
	soma := 0

	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}

	fmt.Println("Soma:", soma)
}
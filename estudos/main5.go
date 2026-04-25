package main

import "fmt"

func main(){
	listaToda := []int{2,10,9,4,5,8,1,3}
	listaMenor := make([]int, 0)
	var i []int
	fmt.Println(i)
	fmt.Println(listaMenor)
	for i := 0; i < len(listaToda); i++ {
		if listaToda[i] < 5 {
			listaMenor = append(listaMenor, listaToda[i])
		}
	}
	fmt.Println(listaMenor)
}
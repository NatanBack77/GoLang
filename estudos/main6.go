package main

import "fmt"

func main(){
	listaToda := []int{2,10,9,4,5,8,1,3}
	segundaLista := listaToda[:3]
	terceiraLista := listaToda[3:]
	ultimoItem := listaToda[len(listaToda)-1:]
	fmt.Println(ultimoItem)
	// fmt.Println(listaToda)
	fmt.Println(terceiraLista)
	fmt.Println(segundaLista)
}
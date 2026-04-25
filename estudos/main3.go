package main

import "fmt"

func main(){
	for numbase := 1; numbase <= 10; numbase++ {
		for multiplicador := 1; multiplicador <= 10; multiplicador++ {
			fmt.Println(numbase, " x ", multiplicador, " = ", numbase*multiplicador)
		}
	}
}
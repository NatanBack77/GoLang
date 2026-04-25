package main

import(
	"fmt"
)

func main(){
   salario := 1000.00
   tipo := "CLT"

   if salario < 1000.00 && tipo == "CLT" {
	  salario -=(salario * 0.1)
   } else if salario <= 2000.00 {
	  salario -=(salario * 0.15)
   } else {
	  salario -=(salario * 0.2)
   }

   fmt.Println("Salário líquido:", salario)
}  
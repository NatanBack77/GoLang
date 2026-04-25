// package main

// import "fmt"

// func main(){
// 	slice := []int{2,8,3,10,5,4,7,9,1}

// 	somaPrimeira := 0
// 	somaSegunda := 0

// 	slice1:= slice[:5]
// 	slice2:= slice[5:]
// 	fmt.Println(len(slice1))
// 	fmt.Println(len(slice2))
// 	for i := 0; i < len(slice1); i++ {
// 		somaPrimeira += slice1[i]
// 	}

// 	for i := 0; i < len(slice2); i++ {
// 		somaSegunda += slice2[i]
// 	}

// 	fmt.Println("soma1", somaPrimeira)
// 	fmt.Println("soma2", somaSegunda)
// }

package main 
 
import "fmt"

func main(){
	slice := []int{2,8,3,10,5,4,7,9,1}
	
	somaPrimeira := 0
	somaSegunda := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] <= 5 {
			somaPrimeira += slice[i]
		} else {
			somaSegunda += slice[i]
		}
}
fmt.Println("soma1", somaPrimeira)
fmt.Println("soma2", somaSegunda)
}
package main

import(
	"fmt"
	"strconv"
)

func main() {

   a := 10
   var numero2 float32
   numero2 = float32(a)
   fmt.Printf("%T\n", numero2)

   b := "20"
   i, _ := strconv.ParseInt(b,10, 64)
   fmt.Printf("%T\n", &i)

   c := 20
   s := strconv.Itoa(c)
   fmt.Printf("%T\n", s)
   
  var d int =  10
   fmt.Printf("%T\n", d)

}



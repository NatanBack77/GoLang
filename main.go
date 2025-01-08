package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d recebied %d \n", workerId, x)
		time.Sleep(time.Second)
	}
}
func main() {
	channel := make(chan int)
  for i:=0;i<1000000;i++{
	go worker(i, channel)
  }

	for i := 0; i < 1000000; i++ {
		channel <- i
	}
}

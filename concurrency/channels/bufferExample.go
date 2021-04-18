package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 8
	fmt.Println(<-ch)
}

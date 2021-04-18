package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Arch:\t\t", runtime.GOARCH)
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("Num CPU:\t", runtime.NumCPU())
	fmt.Println("Num Go Routines:", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Num CPU:\t", runtime.NumCPU())
	fmt.Println("Num Go Routines:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo prints: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar prints: ", i)
	}
}

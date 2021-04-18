package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Arch:\t", runtime.GOARCH)
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("Num CPU:\t", runtime.NumCPU())
	fmt.Println("Num Go Routines:\t", runtime.NumGoroutine())
	foo()
	bar()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo prints: ", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar prints: ", i)
	}
}

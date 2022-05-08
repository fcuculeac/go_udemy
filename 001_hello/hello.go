package main

import "fmt"

func main() {
	fmt.Println("This is my first program")
	foo()
	for i := 0; i < 10; i++ {
		fmt.Println("index", i)
	}
}

func foo() {
	fmt.Println("inside foo")
}

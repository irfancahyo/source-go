package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {

	for i := 0; i < 3; i++ {
		// use Function sayHello
		sayHello()
	}
}

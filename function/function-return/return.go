package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro/Sist"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Cahyo")
	fmt.Println(result)
	fmt.Println(getHello(""))
}

package main

import "fmt"

func main() {
	fruits := map[string]string{
		"name":  "apple",
		"stock": "1",
	}

	fmt.Println(fruits)

	// function map
	// len(map), map[key], map[key] = value, make(map[TypeKey]TypeValue), delete(map,key)
	fmt.Println(len(fruits))

	// map[key]
	fmt.Println(fruits["name"])
	fmt.Println(fruits["stock"])

	// map[key] = value
	fruits["name"] = "orange"
	fruits["stock"] = "2"
	fmt.Println(fruits)

	// make(map[Typekey]TypeValue)
	book := make(map[string]string)
	book["title"] = "Go language"
	book["author"] = "Irfan Cahyo"
	book["status"] = "Denied"

	fmt.Println(book)

	// delete(map,key)
	delete(book, "status")

	fmt.Println(book)
}

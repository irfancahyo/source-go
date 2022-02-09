package main

import "fmt"

func main() {
	// Perulangan dengan for
	counter := 1

	for counter <= 2 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	// for with statement
	for counter := 1; counter <= 2; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	// for range slice
	slice := []string{"Irfan", "Cahyo", "Ariawan"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range map
	names := map[string]string{
		"a": "Irfan",
		"b": "Cahyo",
		"c": "Ariawan",
	}
	for n, v := range names {
		fmt.Printf("%s -> %s\n", n, v)
	}
}

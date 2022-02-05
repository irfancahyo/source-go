package main

import "fmt"

func main() {
	// Operasi perbandingan data yang bernilai bool (benar/salah)
	name1 := "Cahyo"
	name2 := "Cahyo"

	result1 := name1 == name2

	fmt.Println(result1)

	value1 := 20
	value2 := 90

	fmt.Println(value1 < value2)  //true 20 < 90
	fmt.Println(value1 > value2)  //false 20 > 90
	fmt.Println(value1 == value2) //false 20 = 90
	fmt.Println(value1 != value2) //false 20 bukan = 90

}

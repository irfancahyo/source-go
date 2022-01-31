// Variable hanya bisa menyimpan tipe data yang sama
// membutuhkan beberapa Variable untuk menyimpan data yang berbeda jenis

package main

import "fmt"

func main() {
	var name string

	name = "Irfan Cahyo"
	fmt.Println(name)

	//  Variable dapat diubah value datanya
	name = "Cahyo Ariawan"
	fmt.Println(name)
	fmt.Println(len(name))

	// Variable tanpa harus menuliskan tipe data
	var age = 22
	fmt.Println(age)

	// Variable tanpa menggunakan var
	nasionality := "Indonesia"
	fmt.Println(nasionality)

	// Multiple Variable = readable code
	var (
		firstName = "Irfan Cahyo"
		lastName  = "Ariawan"
	)
	fmt.Println(firstName, lastName)
}

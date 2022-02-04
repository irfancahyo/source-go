// Constant adalah Variable yang tidak dapat diubah valuenya setelah dideklarasikan
// ciri Constant tidak menggunakan var melainkan menggunakan kata kunci const

package main

import "fmt"

func main() {

	const firstName string = "Irfan Cahyo"
	const lastName = "Ariawan"
	const value = 100000

	// walaupun dideklarasikan namun tidak digunakan, dia tidak akan error seperti variable
	fmt.Println(firstName, lastName, "memiliki uang sebanyak Rp.", value)

	// Multiple Constant
	const (
		herFirstName string = "Gita"
		herLastName         = "Candra"
	)

	fmt.Println(herFirstName, herLastName)
}

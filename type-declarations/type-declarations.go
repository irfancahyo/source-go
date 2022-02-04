package main

import "fmt"

func main() {
	// Berfungsi untuk membuat alias pada tipe data yang sudah ada
	type noKtp string
	type wakeUp bool

	var ktpCahyo noKtp = "1234567890"
	fmt.Println(ktpCahyo)

	var status wakeUp = true
	fmt.Println(status)
}

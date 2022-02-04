package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 30000
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
	)
	// Konversi Data int
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi Data string ke byte
	name := "Cahyo"
	e := name[0]
	// konversi byte ke string
	eString := string(e)

	fmt.Println(name)
	fmt.Println(eString)

}

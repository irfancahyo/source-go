package main

import "fmt"

func main() {
	// deklarasi array
	// [3] maximal kapasitas data yang dapat ditampung
	var names [3]string
	names[0] = "irfan"
	names[1] = "Cahyo"
	names[2] = "Ariawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	// Array secara langsung
	var ages = [3]int{
		20, 21, 22,
	}

	fmt.Println(ages)

	// function array
	// len(array) , menampilkan panjang array
	fmt.Println(len(names))
	fmt.Println(len(ages))
	// array[index] , menampilkan posisi index
	fmt.Println(ages[0])
	fmt.Println(ages[1])
	fmt.Println(ages[2])
	// array[index] = value , mengubah data pada posisi index
	ages[2] = 23
	fmt.Println(ages[2])

	// array tanpa maximal index
	var barangs = []string{
		"sabun",
		"sampo",
		"odol",
		"masker",
		"sendal",
	}

	fmt.Println(barangs)
}

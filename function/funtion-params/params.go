package main

import "fmt"

// Function Parameter

func addMahasiswa(firstNames string, lastNames string, prodi string, smt int) {
	fmt.Println("Hello", firstNames, lastNames, "Kamu telah terdaftar di prodi", prodi, "silahkan untuk mendaftar ulang untuk memulai perkuliahan semester", smt)
}

func main() {
	firstNames := "Irfan"
	addMahasiswa(firstNames, "Cahyo", "Rekayasa Perangkat Lunak", 1)

}

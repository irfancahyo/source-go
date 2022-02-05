package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 70

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus1 = lulusUjian && lulusAbsensi
	var lulus2 = lulusUjian || lulusAbsensi

	fmt.Println(lulus1) // and
	fmt.Println(lulus2) // or

	// real code here
	fmt.Println(ujian >= 80 && absensi >= 80)

}

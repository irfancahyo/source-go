package main

import "fmt"

func main() {
	// Membuat Slice dari array
	/*
	   array[low:high] = membuat slice dari array mulai dari low index to high index
	   array[low:] = membuat slide dari array mulai dari low index to end index
	   array[:high] = membuat slice dari array mulai dari index 0 to index sebelum high
	   array[:] = membuat slice dari array mulai dari index 0 to end index
	*/

	var months = [12]string{
		"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Sep", "Okt", "Nov", "Des",
	}

	slice1 := months[0:5]

	fmt.Println(slice1)

	/*
		Function Slice :
		len(slice) , menampilkan panjang slice
		cap(slice) , menmpilkan kapasitas slice
		append(slice, data) , membuat slice baru dengan menambahkan data e posisi terakhir slice / membuat array baru bisa slice sudah penuh
		make([]TypeData, lenght, capacity) , membuat slice baru
		copy(destination, source) , enyalin slice dari source destination
	*/

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Append Slice
	days := [...]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	daySlice := days[1:]
	fmt.Println(days)
	fmt.Println(daySlice)

	daySlice2 := append(daySlice, "Sun")
	fmt.Println(daySlice2)

	// Make Slice
	newSlice := make([]string, 3, 5)
	newSlice[0] = "Irfan"
	newSlice[1] = "Cahyo"
	newSlice[2] = "Ariawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// *perbedaan array & slice
	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// example

	e := []int{1, 2, 3}
	e = append(e, 4)
	fmt.Println(e, len(e))

	f := make(map[string]int)
	f["one"] = 1
	f["two"] = 2
	fmt.Println(f, len(f), f["one"], f["three"])

}

package main

import "fmt"

func main() {
	// Continue menghentikan perulangan saat ini dan melanjutkan perulangan berikutnya
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}

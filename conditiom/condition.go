package main

import "fmt"

func main() {
	var value = 3

	// if else condition
	if value == 10 {
		fmt.Println("Lulus Sempurna")
	} else if value > 5 {
		fmt.Println("Lulus")
	} else if value == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("Tidak Lulus nilai anda : %d\n", value)
	}

	// Variable Temporary
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Switch Case
	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Multiple Case Output
	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// default
	var point4 = 0

	switch point4 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Switch dengan if else statement
	var point5 = 6

	switch {
	case point5 == 8:
		fmt.Println("perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Bercabang
	var point6 = 7

	if point6 > 7 {
		switch point6 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point6 == 5 {
			fmt.Println("not bad")
		} else if point6 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point6 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}

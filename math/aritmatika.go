package main

import "fmt"

func main() {
	p := 20
	q := 4

	var (
		// Additional
		result1 = p + q
		// Subtraction
		result2 = p - q
		// Multiplication
		result3 = p * q
		// Division
		result4 = p / q
	)

	fmt.Println("Result of p + q =", result1)
	fmt.Println("Result of p - q =", result2)
	fmt.Println("Result of p * q =", result3)
	fmt.Println("Result of p / q =", result4)

}

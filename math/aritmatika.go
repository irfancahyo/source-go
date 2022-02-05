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

	// Augmented Assigment
	// a = a + 10 ~ a += 10
	// a = a - 10 ~ a -= 10
	// a = a * 10 ~ a *= 10
	// a = a + 10 ~ a /= 10
	// a = a % 10 ~ a %= 10 sisa bagi
	var i = 10
	i += 10
	fmt.Println("i += 10 hasilnya", i)

	// Unary Operations
	// ++, --, +, -, !

	i++ //i = i + 1
	fmt.Println("i++ =", i)

}

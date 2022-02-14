package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var p Point
	p.X = 1
	p.Y = 2

	var q *Point
	q.X = 3
	q.Y = 4

	fmt.Println(p, q)
}

package main

import "fmt"

func main() {
	var X, d1, d2 int
	fmt.Scanln(&X)

	d1 = X % 10 //4
	X = X / 10
	konsekutif := true
	for X > 0 && konsekutif {
		d2 = X % 10 //3
		konsekutif = d1-d2 == 1 || d2-d1 == 1
		d1 = d2
		X = X / 10
	}
	fmt.Println(konsekutif)
}

package main

import (
	"fmt"
)

func main() {
	f := func(a, b int) int {
		return a + b
	}

	x := f(10, 20)
	y := subtrair(30, 50)
	z := sum(10, 2, 5)
	fmt.Println(x, y, z)
}

func sum(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}

	return out
}

func subtrair(a, b int) int {
	return a - b
}

func somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

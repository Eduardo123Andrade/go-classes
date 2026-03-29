package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello, world!")
	// a, b := swap(10, 20)
	// fmt.Println(a, b)
	res, rem := dividir(10,3)
	fmt.Println(res, rem)
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (int, int) {
	res := a / b
	rem := a % b
	return res, rem
}
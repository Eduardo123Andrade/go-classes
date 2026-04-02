package main

import (
	"fmt"
)

func main() {
	arr1 := [3]int{}
	arr2 := [3]int{1, 2, 3}
	arr3 := [10]int{5: 400, 7: 300}

	const x = 10
	arry := [x]int{}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arry)
}

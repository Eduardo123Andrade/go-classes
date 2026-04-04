package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [4]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}

	naoMudaArray(arr1)
	mudaArray(&arr2)
	mudaSlice(slice)

	fmt.Println(arr1)  // [1, 2, 3, 4]
	fmt.Println(arr2)  // [123, 2, 3, 4]
	fmt.Println(slice) // [123, 2, 3, 4]
}

func mudaArray(arr *[4]int) {
	arr[0] = 123
}

func naoMudaArray(arr [4]int) {
	arr[0] = 123
}

func mudaSlice(arr []int) {
	arr[0] = 123
}

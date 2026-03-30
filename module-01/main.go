package main

import (
	"fmt"
)

func main() {
	const x = 3
	takeInt32(x)
	takeInt64(x)
}

func takeFloat32(x float32) {
	fmt.Println(x)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

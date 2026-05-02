package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.OpenFile("files/file.txt", os.O_WRONLY, 0644)
	file, err := os.OpenFile("files/file.txt", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// stats, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// stats

	// err = os.WriteFile(file.Name(), []byte("hello, world2"), 0644)

	// result, err := file.Re("testes")

	r, err := file.Write([]byte("hello, world"))
	r, err = file.Write([]byte("\n"))
	r, err = file.Write([]byte("hello, world2"))
	r, err = file.Write([]byte("\n"))

	if err != nil {
		panic(err)	
	}

	// if err == nil {
	// 	fmt.Println("File written successfully")
	// }

	fmt.Println(r)


}

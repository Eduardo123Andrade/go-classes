package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["Pedro"] = "Pessoa"
	m["foo"] = "bar"

	for k, v := range m {
		fmt.Println(k, v)
	}
}

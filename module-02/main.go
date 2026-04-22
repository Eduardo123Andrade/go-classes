package main

import (
	"fmt"
)

type Animal interface {
	Sound() string
}

type Dog struct{
	Name string
}

type Cat struct {}

func (d *Dog) Sound() string {
	return "Au! Au!"
}

func (c *Cat) Sound() string {
	return "Miau!"
}

func takeAnimal(a Animal) (x string) {
	x = "<nil>"
	switch t := a.(type) {
	case *Dog:
		 x = t.Sound()
	case *Cat:
		x = t.Sound()
	}

	return
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func takeString(a any) {
	str:= a.(string)
	fmt.Println(str)
}

func main() {
	c := Cat{}
	sound := takeAnimal(&c)
	fmt.Println(sound)
	// var a Animal
	// fmt.Println(a.Sound())
}

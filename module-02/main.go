package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))

}

var (
	ErrQualquer  = errors.New("Erro")
	ErrQualquer2 = errors.New("Erro 2")
)

func a() error { return ErrQualquer }
func b() error { return ErrQualquer2 }

func foo() error {
	var errorRetults error

	if err := a(); err != nil {
		errorRetults = errors.Join(errorRetults, err)
	}

	if err := b(); err != nil {
		errorRetults = errors.Join(errorRetults, err)

	}

	return errorRetults

}

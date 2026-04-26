package main

func main() {
	foo(Foo{})
	foo(Bar{})
}

type Foo struct{}

func (Foo) Fazer() {}

type Bar struct{}

func (Bar) Fazer() {}

type MyInterface interface {
	Foo | Bar
	Fazer()
}

// func foo[T MyInterface](arg T) {
func foo[T MyInterface](arg T) {
	arg.Fazer()
}

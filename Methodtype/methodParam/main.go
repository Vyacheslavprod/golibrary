//Practice method
package main

import "fmt"

type MyType string

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n) + otherNumber) //Преобразование n(T) к Number 
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n) - otherNumber) //Преобразование n(T) к Number 
}

func main() {
	value := MyType("MyType value")
	value.MethodWithParameters(4, true)
	fmt.Println(value.WithReturn())

	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
}
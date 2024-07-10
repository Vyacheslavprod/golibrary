//Вызов методов интерфейса
package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoizeMaker interface {
	MakeSound()
}

func play(n NoizeMaker) {
	n.MakeSound()
}

func main()  {
	var toy NoizeMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
}
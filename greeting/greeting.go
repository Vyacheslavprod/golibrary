package greeting

import "fmt"

func Hello()  {
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi")
}

func AllGreetings() {
	Hello()
	Hi()
}

var S = "Hello"
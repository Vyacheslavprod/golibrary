package main

import (
	"github.com/vyacheslavprod/golibrary/greeting"
	"fmt"
	"github.com/vyacheslavprod/golibrary/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println("This is Module")
	greeting.AllGreetings()
	greeting.S = "374"
	fmt.Println(greeting.S)
	deutsch.Hello()
	deutsch.GutenTag()
}



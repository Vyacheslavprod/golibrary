package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func doMath(passedFunction func(int, int) float64) {
	result := passedFunction(10, 2)
	fmt.Println(result)
}

func main() {
	//var myFunction func()
	myFunction := sayHi
	myFunction()
	fmt.Println()
	twice(sayHi)
	twice(sayBye)
	fmt.Println()
	//var mathFunction func(int, int) float64
	mathFunction := divide
	fmt.Println(mathFunction(5, 2))
	doMath(divide)

}
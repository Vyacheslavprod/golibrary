//Определение метода для типа
package main

import "fmt"

type Mytype string
type Mytype2 string


func (m Mytype) sayHi() {
	fmt.Println("Hi")
}

func (m Mytype2) sayHi2() {
	fmt.Println("Hi from", m) //Параметр получателя m - value, another = Mytype
}

func main() {
	value := Mytype("a MyType value")
	value.sayHi()
	anotherValue := Mytype("another value")
	anotherValue.sayHi()

	value2 := Mytype2("a MyType value")
	value2.sayHi2()
	anotherValue2 := Mytype2("another value")
	anotherValue2.sayHi2()
}
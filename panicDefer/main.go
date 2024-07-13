package main

import "fmt"

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	} else {
		panic(p) //Непредвиденная паника
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
}
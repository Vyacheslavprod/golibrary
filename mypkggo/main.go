//Вызов методов интерфейса
package main

import (
	"fmt"
	"github.com/vyacheslavprod/golibrary/mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(2)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
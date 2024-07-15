package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func greeting(myChan chan string)  {
	myChan <- "hi"
}

func main() {
	go a()
	go b()

	myChan := make(chan string)
	go greeting(myChan)
	//fmt.Println(<-myChan)
	receivedValue := <- myChan
	fmt.Println(receivedValue)

	time.Sleep(2 * time.Second)
	fmt.Println("end main()")


}
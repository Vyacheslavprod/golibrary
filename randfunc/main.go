package main

import (
	"fmt"
	"math/rand"
)

//Если вызвать recover() через fmt - вернет значение паники
func calmDown() {
	fmt.Println(recover())
}

func awardPrize() {
	doorNumber := rand.Intn(3)
	fmt.Println(doorNumber)
	if doorNumber == 1 {
		fmt.Println("You win a cruise")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		defer calmDown()
		panic("invalid door number")
	}
}

func main() {
	awardPrize()
}
//pass_fail сообщает, сдал ли пользователь экзамен.

package main

import (
	"fmt"
	"log"
	"github.com/vyacheslavprod/golibrary/keyboard"
)



func main() {
	fmt.Print("Введите число: ") //Запросить у пользователя значение
	grade, err := keyboard.GetFloat()
	if err != nil{
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "Экзамен сдан"
	} else {
		status = "Экзамен не сдан"
	}
	fmt.Printf("Ваши баллы: %v - %s!\n", grade, status)
}



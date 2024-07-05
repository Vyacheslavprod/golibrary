package main

import (
	"errors"
	"fmt"
	"log"
)

//Расход краски
func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a heith of %0.2f is invalid", height)
	}
	
	area := width * height
	return area / 10.0, nil
}

func main() {
	err := errors.New("height can't be negative") // Создание ошибки
	
	fmt.Println(err.Error())
	fmt.Println(err) // функции fmt и log выведут ошибку без метода Error()

	password := 7475
	// форматирование числа или другого значения в выводе об ошибке
	err = fmt.Errorf("Пароль: %d - не верный!", password)
	fmt.Println(err)

	




	//Расход краски
	
	amount, err := paintNeeded(5.2, -3.0)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)

	amount1, err := paintNeeded(4.5, -3.2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount1)
	}

	amount2, err := paintNeeded(-3.4, -21.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount2)
}
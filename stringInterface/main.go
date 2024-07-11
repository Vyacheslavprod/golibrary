package main

import "fmt"

type CoffePot string

func (c CoffePot) String() string {
	return string(c) + " coffee pot"
}

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f ml", m)
}

func main() {
	coffePot := CoffePot("LuxBrew")
	fmt.Print(coffePot, "\n")
	fmt.Println(coffePot)
	fmt.Printf("%s\n", coffePot)

	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
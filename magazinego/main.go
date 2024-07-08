package main

import (
	"fmt"

	"github.com/vyacheslavprod/golibrary/magazine"
)

func main() {
	//Обращение к структуре Address. Вариант - 1
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber)
	//Обращение к структуре Address. Вариант - 2
	var subscriber1 magazine.Subscriber
	subscriber1.HomeAddress.PostalCode = "23445"
	fmt.Println("Postal code:", subscriber1.HomeAddress.PostalCode)

	employee := magazine.Employee{Name: "Joy Car"}
	employee.HomeAddress.State = "OR"
	employee.HomeAddress.City = "Portland"
	employee.HomeAddress.Street = "456 Elm St"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println(employee)

}
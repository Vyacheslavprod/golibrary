// average вычисляет среднее значение, считывая строки из файла и преобразуя их в float
package main

import (
	"fmt"
	"log"

	"github.com/vyacheslavprod/golibrary/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt") //Разбирает содержащиеся в нем числа и сохраняет массив
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum / sampleCount)
}
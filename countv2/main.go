//count v2 - подсчитывает количество вхождений каждой строки в файле
package main

import (
	"fmt"
	"log"

	"github.com/vyacheslavprod/golibrary/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes2.txt") //Получает слайс строк - Имен
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++ //В карте создается ключ с нулевым значением, который увеличивается
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
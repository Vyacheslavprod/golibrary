//count подсчитывает количество вхождений каждой строки в файле
package main

import (
	"fmt"
	"log"

	"github.com/vyacheslavprod/golibrary/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string //Переменная для хранения сегмента с именами кандидатов
	var counts []int //Сегмент с количеством вхождений каждого имени
	for _, line := range lines {  
		matched := false
		for i, name := range names { //Перебор всех значений из сегмента names
			if name == line { //Если эта строка совпадает с текущим именем...
				counts[i]++ //...увеличивает соответствующий счетчик
				matched = true //Устанавливаем признак обнаруженного совпадения
			}
		}
		if !matched { //Если совпадение не найдено...
			names = append(names, line) //...добавить его как новое имя в сегмент
			counts = append(counts, 1) //И добавить новый счетчик(текущая строка станет первым вхождением)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
//Пакет datafile предназначен для чтения данных из файлов.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloats читает значение float64 из каждой строки файла
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0 //Переменная для хранения индекса, по которому должно выполняться присваивание
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //Строка прочитанная из файла преобразуется в float
		if err != nil {
			return numbers, err
		}
		i++ //Переход к следующему индексу массива
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil { //Если при сканировании файла произошла ошибка, вернуть ее. строка 19
		return numbers, err
	}

	return numbers, nil
}
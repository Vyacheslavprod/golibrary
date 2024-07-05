//Пакет datafile предназначен для чтения данных из файлов.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloats читает значение float64 из каждой строки файла
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64) //Строка прочитанная из файла преобразуется в float
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number) //Переход к следующему индексу массива
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil { //Если при сканировании файла произошла ошибка, вернуть ее. строка 19
		return nil, scanner.Err()
	}

	return numbers, nil
}
// Пакет datafile предназначен для чтения данных из файлов.
package datafile

import (
	"bufio"
	"os"
)

//GetStrings читает строку из каждой строки данных файла.
func GetStrings(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return lines, nil
}


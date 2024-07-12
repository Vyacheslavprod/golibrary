package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Рекурсивная функция, которая получает путь для обработки
func scanDirectory(path string) error {
	fmt.Println(path) //Выводит текущий каталог
	files, err := os.ReadDir(path) //Получает сегмент с содержимым каталога
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filepath := filepath.Join(path, file.Name()) //Соединяет путь каталога и имя файла через символ /
		if file.IsDir() { //Если это каталог...
			scanDirectory(filepath) //..рекурсивно вызывает scanDirectory, но на этот раз с путем подкатолога
		} else {
			fmt.Println(filepath) //Если это обычный файл, просто вывести его имя с путем
		}
	}
	return nil
}

func main() {
	scanDirectory("my_directory") 
}
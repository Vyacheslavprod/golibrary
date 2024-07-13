package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

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
	defer reportPanic()
	scanDirectory("my_directory") 
}
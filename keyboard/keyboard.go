//Package keyboard считать строку, удаляет все символы-пропуски в начале и конце строки, преобразовать строку в float64
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloat получение текста с клавиатуры
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) //Создать средство чнения для получения текста с клавиатуры
	input, err := reader.ReadString('\n') //Возвращает текст, введеный до нажатия Enter
	if err != nil{
		return 0, err
	}

	input = strings.TrimSpace(input) //удаляем символ новой строки из ввода \n
	number, err := strconv.ParseFloat(input, 64) //преобразовать строку в float64
	if err != nil{
		return 0, err
	}
	return number, nil
}
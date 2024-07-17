package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//GuestBook - структура, используемая при отображении view.html
type GuestBook struct {
	SignatureCount int
	Signatures []string
}

//chek вызывает log.Fatal для любых ошибок, отличных от nil
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//viewHandler читает записи гостевой книги и выводит их вместе со счетчиком записей
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestBook := GuestBook{
		SignatureCount: len(signatures),
		Signatures: signatures,
	}
	err = html.Execute(writer, guestBook) //Данные структуры GuestBook вставляются в шаблон, а результат записывается в ResponseWriter
	check(err)
}

//newHandler отображает форму для ввода записи
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

//createHandler получает запрос POST с добавляемой записью и присоединяет ее к файлу signatures
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature") //Получает значение поля формы
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature) //Добавляет содержимое поля формы в файл
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	//Если возвращается ошибка, указывающая на то, что файл не существует...
	//...вернуть nil вместо сегмента строк
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines

}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
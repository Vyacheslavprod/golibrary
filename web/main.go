package main

import (
	"log"
	"net/http"
)

//Функция для генерирования ответа
//writer - значение для обновления ответа, который будет отправлен браузеру
//request - значение, представляющее запрос от браузера
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
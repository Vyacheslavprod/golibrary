package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Clients struct {
	Id int
	Document string
	Language string
	Status string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var database *sql.DB
 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("select * from productdb.Products")
	check(err)
	defer rows.Close()
	clients := []Clients{}

	for rows.Next(){
		p := Clients{}
		err := rows.Scan(&p.Id, &p.Document, &p.Language, &p.Status)
		if err != nil{
			fmt.Println(err)
			continue
		}
		clients = append(clients, p)
	}

	tmpl, _ := template.ParseFiles("myproject2/index.html")
	tmpl.Execute(w, clients)
}
 
func main() {
	db, err := sql.Open("mysql", "root:password@/productdb")
	check(err)
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server is listening...")
	err = http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
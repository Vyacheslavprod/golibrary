package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Clients struct {
	Id       int
	Document string
	Language string
	Status   string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from documentsdb.Documents")
	check(err)
	defer rows.Close()
	clients := []Clients{}

	for rows.Next() {
		p := Clients{}
		err := rows.Scan(&p.Id, &p.Document, &p.Language, &p.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		clients = append(clients, p)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, clients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EditPage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
 
    row := db.QueryRow("select * from documentsdb.Documents where id = ?", id)
    p := Clients{}
    err := row.Scan(&p.Id, &p.Document, &p.Language, &p.Status)
    if err != nil{
        log.Println(err)
        http.Error(w, http.StatusText(404), http.StatusNotFound)
    }else{
        tmpl, _ := template.ParseFiles("templates/edit.html")
        tmpl.Execute(w, p)
    }
}


func EditHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Println(err)
    }
    id := r.FormValue("id")
    document := r.FormValue("document")
    language := r.FormValue("language")
    status := r.FormValue("status")
 
    _, err = db.Exec("update documentsdb.Documents set document=?, language=?, status = ? where id = ?", 
	document, language, status, id)
 
    if err != nil {
        log.Println(err)
    }
    http.Redirect(w, r, "/", 301)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:IDIlfoVeTNGt9T8sifMJ0tGFYQyW+kC4n3ZwEaHwUvSR@/documentsdb")
	check(err)
	err = db.Ping()
	check(err)
	defer db.Close()

	router := mux.NewRouter()
    router.HandleFunc("/", IndexHandler)

    router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
    router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
     
    http.Handle("/",router)
	
	// open file
	fs := http.FileServer(http.Dir("document/"))
    http.Handle("/static}", http.StripPrefix("/static/", fs))

	fmt.Println("Server is listening...")
	err = http.ListenAndServe("localhost:8181", nil)
	check(err)
}

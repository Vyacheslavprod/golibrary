package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

type Part struct {
	Name string
	Count int
}

type Subscriber struct {
	Name string
	Rate float64
	Active bool
}

func main() {
	templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})

	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)

	subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)
}
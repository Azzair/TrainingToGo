package main

import (
	"log"
	"os"
	"text/template"
)

var (
	appTemplate *template.Template
)

func init() {
	appTemplate = template.Must(template.ParseGlob("template/*"))
}

func main() {
	err := appTemplate.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln("Opps!")
	}
}

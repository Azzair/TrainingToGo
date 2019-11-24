package main

import (
	"log"
	"os"
	"text/template"
)

var (
	appTemplate *template.Template
	funcTemp    = template.FuncMap{
		"sHello": sHello,
	}
)

func init() {
	appTemplate = template.New("")
	appTemplate = appTemplate.Funcs(funcTemp)
	var err error
	appTemplate, err = appTemplate.ParseFiles("template/hello.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := appTemplate.ExecuteTemplate(os.Stdout, "hello.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func sHello() string {
	return "Hello from func!"
}

package main

import (
	"log"
	"os"
	"text/template"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", "MEMEMEMEME")
	if err != nil {
		log.Fatalln(err)
	}
}

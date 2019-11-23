package main

import (
	"log"
	"os"
	"text/template"
)

var (
	appTemplates *template.Template
)

type dataProvider struct {
	Name string
	Text string
}

func init() {
	appTemplates = template.Must(template.ParseGlob("template/*"))
}

func main() {
	appData := dataProvider{
		"My Name",
		"Досить простий текст для зразка та перевірки роботоздатності",
	}
	err := appTemplates.ExecuteTemplate(os.Stdout, "index.html", appData)
	if err != nil {
		log.Fatalln("OOps")
	}
}

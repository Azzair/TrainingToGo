package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	tmp, err := template.ParseFiles("tmp.html")
	if err != nil {
		log.Fatalln("File not found")
	}
	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("ERROR to execute template")
	}

}

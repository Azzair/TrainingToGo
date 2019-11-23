package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tmp, err := template.ParseFiles("tmp.html")
	if err != nil {
		log.Fatalln("File not found")
	}
	fmt.Fprint("Hello")
	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("ERROR to execute template")
	}

}

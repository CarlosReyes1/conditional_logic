package main

import (
	"log"
	"os"
	"text/template"
)

type animal struct {
	Breed string
	Size  string
}

type doubleAni struct {
	animal
	KillOnCommand bool
}

func main() {
	a1 := doubleAni{
		animal: animal{
			Breed: "Pitbull",
			Size:  "medium",
		},
		KillOnCommand: false,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, a1)
	if err != nil {
		log.Fatalln(err)
	}
}

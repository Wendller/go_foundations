package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	courses := Courses{
		Course{Name: "GO", Duration: 60},
		Course{Name: "Node", Duration: 100},
		Course{Name: "FullCycle", Duration: 180},
	}

	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"html/template"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := tmp.Execute(w, courses)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}

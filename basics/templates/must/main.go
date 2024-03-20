package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Duration int
}

func main() {
	course := Course{Name: "GO", Duration: 40}
	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Duration: {{.Duration}}\n"))

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

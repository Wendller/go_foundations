package main

const expression = "Hello, World"

var boo bool

type ID int

var (
	b  bool    = true
	i  int     = 10
	s  string  = "wendler"
	f  float64 = 1.2
	id ID      = 30
)

func main() {
	shorthand := "Infered string"
	shorthand = "already declared"

	println(shorthand)
	println(id)
}

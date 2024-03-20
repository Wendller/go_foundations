package main

import "fmt"

type ID int

var (
	b  bool    = true
	i  int     = 10
	s  string  = "wendler"
	f  float64 = 1.2
	id ID      = 30
)

func main() {
	fmt.Printf("The f type is %T", f)
	fmt.Printf("The f value is %v", f)
}

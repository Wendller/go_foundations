package main

import "fmt"

type Address struct {
	street string
	number int
	state  string
}

type Person interface {
	deactivate()
}

type Client struct {
	name    string
	active  bool
	address Address
}

func (c Client) deactivate() {
	c.active = false
}

func deactivation(person Person) {
	person.deactivate()
}

func main() {
	wend := Client{name: "Wendler", active: true}
	wend.address.state = "PI"

	wend.deactivate()
	deactivation(wend)

	fmt.Println(wend.name)
}

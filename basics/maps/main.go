package main

import "fmt"

func main() {
	salaries := map[string]int{
		"wendler":  5800,
		"wendy":    7500,
		"wendler2": 9000,
	}

	fmt.Println(salaries["wendy"])
	delete(salaries, "wendler")

	hashMap := make(map[string]string)
	hashMap2 := map[string]string{}
	hashMap["node"] = "18.15.1"

	fmt.Println(hashMap["node"])
	fmt.Println(hashMap2["node"])

	for key, value := range salaries {
		fmt.Printf("The salary of %s is %d\n", key, value)
	}
}

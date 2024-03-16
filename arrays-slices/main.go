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
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 3
	myArray[2] = 5

	fmt.Println(len(myArray))

	for index, value := range myArray {
		fmt.Printf("The index is %d and value is %d\n", index, value)
	}

	// More ways to use for
	// 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 2 (as a while)
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 3
	// for {
	// }

	mySlice := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[:4]), cap(mySlice[:4]), mySlice[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])

	mySlice = append(mySlice, 12)
}

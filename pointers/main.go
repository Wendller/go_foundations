package main

func main() {
	// a := 10
	// println(a)
	// println(&a) // memory address for variable

	// var pointer *int = &a // set a variable to the address
	// *pointer = 20         // change the value direct from the address
	// println(a)

	// b := &a
	// println(b)
	// println(*b) // the value at that address in memory

	workWithCopy()
	workWithPointer()
}

func sum(a, b int) int {
	a = 50
	return a + b
}

func sum_with_pointer(a, b *int) int {
	*a = 50
	return *a + *b
}

func workWithCopy() {
	a := 10
	b := 20

	sum(a, b)
	println(a)
}

func workWithPointer() {
	a := 10
	b := 20

	sum_with_pointer(&a, &b)
	println(a)
}

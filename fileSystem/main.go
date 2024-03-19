package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Writing
	file, err := os.Create("doc.txt")
	if err != nil {
		panic(err)
	}

	size, err := file.WriteString("Long live to warriors\n")
	// size, err := file.Write([]byte("Long live to warriors"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("File have %d bytes\n", size)

	file.Close()

	// Reading
	read_file, err := os.ReadFile("doc.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(read_file))

	// Lazy reading
	big_file, err := os.Open("doc.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(big_file)
	buffer := make([]byte, 10)

	for {
		pack, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:pack]))
	}
}

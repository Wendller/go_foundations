package main

import (
	"errors"
	"fmt"
)

func main() {
	total, err := isOverLimit(1500, 2200)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(total)

	// fmt.Println(sumNums(1,4,6,7,9,15))
	// fmt.Println(sumNums(1,4,7))

	// CLOJURE
	FUNC_MODIFIER := 3
	total_clojured := func() int {
		return sumNums(1, 3, 5, 8, 18, 27, 83, 90) * FUNC_MODIFIER
	}()

	fmt.Println(total_clojured)
}

func isOverLimit(tax int, costs int) (int, error) {
	if total := tax + costs; total < 4000 {
		return total, nil
	}
	return tax + costs, errors.New("total costs is invalid")
}

func sumNums(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// CLOJURES - Funcoes anonimas que recebem outra funcao

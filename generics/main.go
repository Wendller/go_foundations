package main

type MapValue interface {
	int | float64
}

func sumMapValues[T MapValue](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func main() {
	mInt := map[string]int{"paul": 1000, "chani": 1500, "muabdib": 3000}
	mFloat := map[string]float64{"paul": 100.0, "chani": 150.20, "muabdib": 300.50}

	println(sumMapValues(mInt))
	println(sumMapValues(mFloat))
}

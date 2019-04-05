package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 15, 2, 8, 32, 14, 22, 1, 3, 6, 9, 0}

	sort.Ints(numbers)

	fmt.Printf("%+v", numbers)
}

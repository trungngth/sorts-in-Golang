package main

import (
	"fmt"

	"sort/sort"
)

func main() {
	input := []int64{1, 2, 3, 4, 23, 5, 456, 132, 1, -99, -2}
	fmt.Println(sort.BubbleSort(input))
}

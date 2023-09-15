package hw1

import (
	"fmt"
	"sort"
)

func compareSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3, 2, 1}
	slice3 := []int{1, 2, 3, 4, 6}

	fmt.Println(compareSlices(slice1, slice2))
	fmt.Println(compareSlices(slice1, slice3))
}

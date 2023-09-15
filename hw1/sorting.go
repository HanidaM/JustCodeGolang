package hw1

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	slice := IntSlice{5, 2, 8, 1, 9, 3}
	fmt.Println("Before sorting:", slice)

	sort.Sort(slice)
	fmt.Println("After sorting:", slice)
}

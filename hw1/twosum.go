package hw1

func twoSum(nums []int, target int) []int {
	cmp := make(map[int]int)

	for i, num := range nums {
		if complementIndex, ok := cmp[num]; ok {
			return []int{complementIndex, i}
		}

		cmp[target-num] = i
	}

	return []int{}
}

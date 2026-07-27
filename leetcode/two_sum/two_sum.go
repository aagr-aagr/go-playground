package twosum

func twoSum(nums []int, target int) []int {

	seenNumber := make(map[int]int)
	result := make([]int, 2, 2)

	for x, y := range nums {
		_, ok := seenNumber[target-y]
		if ok {
			result[1] = x
			result[0] = seenNumber[target-y]
		}
		seenNumber[y] = x
	}

	return result
}

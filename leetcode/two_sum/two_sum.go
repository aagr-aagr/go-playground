package twosum

func twoSum(nums []int, target int) []int {

	//use map[int]int instead:w http.ResponseWriter, r *http.Request
	result := make([]int, 2, 2)

	for x, y := range nums {
		target = target - y
		if target == 0 {
			result[1] = x
			break
		} else {
			result[0] = x
			continue
		}
	}
	return result
}

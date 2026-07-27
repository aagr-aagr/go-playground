package containsduplicate

func containsDuplicate(nums []int) bool {

	numsMap := make(map[int]bool)

	for _, y := range nums {
		_, ok := numsMap[y]

		if ok {
			return true
		} else {
			numsMap[y] = false
		}
	}

	return false

}

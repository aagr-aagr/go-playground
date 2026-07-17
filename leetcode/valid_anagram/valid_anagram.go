package validanagram

func validAnagram(s string, t string) bool {
	wordsCount := make(map[rune]int)

	for _, y := range s {
		wordsCount[y]++
	}

	for _, y := range t {
		wordsCount[y]--
	}

	for _, j := range wordsCount {
		if j != 0 {
			return false
		}
	}
	return true

}

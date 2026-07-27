package groupanagrams

import "slices"

func groupAnagrams(strs []string) [][]string {
	sortedStrings := make(map[string][]string)

	for _, y := range strs {
		byteString := []byte(y)
		slices.Sort(byteString)
		s := string(byteString)
		sortedStrings[s] = append(sortedStrings[s], y)

	}

	result := make([][]string, 0, len(sortedStrings))
	for _, v := range sortedStrings {
		result = append(result, v)
	}

	return result

}

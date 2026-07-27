package groupanagrams

import (
	"cmp"
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	tests := []struct {
		name   string
		strs   []string
		result [][]string
	}{
		{"groupped", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"empty", []string{""}, [][]string{{""}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			got = sortSlices(got)
			tt.result = sortSlices(tt.result)
			if len(got) != len(tt.result) {
				t.Errorf("The result: %v is not what's expected: %v", got, tt.result)
			} else {
				for i := range got {
					if !slices.Equal(got[i], tt.result[i]) {
						t.Errorf("The result: %v is not what's expected: %v", got, tt.result)
					}
				}
			}

		})

	}

}

func sortSlices(str [][]string) [][]string {
	for _, x := range str {
		slices.Sort(x)
	}
	slices.SortFunc(str, func(a, b []string) int {
		return cmp.Compare(a[0], b[0])
	})
	return str
}

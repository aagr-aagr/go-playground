package topkfrequentelements

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		count  int
		result []int
	}{
		{"test1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"test2", []int{3, 1, 3, 2, 2, 2, 3, 3, 4, 5, 1, 1, 1, 1, 3, 2}, 3, []int{1, 2, 3}},
		{"test3", []int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.input, tt.count)
			slices.Sort(got)
			slices.Sort(tt.result)
			if !slices.Equal(got, tt.result) {
				t.Errorf("Slices %v and %v -- not equal", got, tt.result)
			}

		})
	}

}

package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic_pair", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !slices.Equal(got, tt.want) {
				t.Errorf("The array %v is different from expected %v", got, tt.want)
			}
		})
	}
}

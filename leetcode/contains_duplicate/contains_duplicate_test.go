package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"contains", []int{1, 2, 3, 1}, true},
		{"no_dups", []int{1, 2, 3, 4}, false},
		{"dups_long", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("The result is different. Expected: %v, but got: %v", tt.want, got)
			}
		})
	}

}

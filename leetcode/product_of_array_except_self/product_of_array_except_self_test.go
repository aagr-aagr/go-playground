package productofarrayexceptself

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		result []int
	}{
		{"test1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"test1", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.input)
			want := tt.result
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("Slices don't match. Want: %v and got: %v", want, got)
			}

		})

	}

}

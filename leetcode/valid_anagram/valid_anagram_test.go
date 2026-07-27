package validanagram

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name   string
		str1   string
		str2   string
		result bool
	}{
		{"valid", "anagram", "nagaram", true},
		{"not_valid", "rat", "car", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnagram := validAnagram(tt.str1, tt.str2)
			if gotAnagram != tt.result {
				t.Errorf("The stings %v is different from expected %v", gotAnagram, tt.result)
			}

		})
	}

}

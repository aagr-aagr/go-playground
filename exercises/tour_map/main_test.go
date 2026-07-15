package main

import (
	"maps"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want map[string]int
	}{
		{"basic", "I love learning GO!", map[string]int{"I": 1, "love": 1, "learning": 1, "GO!": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordCount(tt.str)
			if !maps.Equal(got, tt.want) {
				t.Errorf("The array %v is different from expected %v", got, tt.want)
			}
		})
	}
}

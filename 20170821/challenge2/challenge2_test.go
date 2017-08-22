package main

import (
	"testing"
)

func TestPalindrones(t *testing.T) {
	testTable := map[string]bool{
		"ABBA": true,
		"X":    true,
		"소있고지게지고있소": true,
		"いかにもにがい":   false,
		"ANNA":      true,
		"RACECAR":   true,
		"DANIEL":    false,
	}
	for text, expected := range testTable {
		got := IsPalindrome(text)
		if got != expected {
			t.Errorf("IsPalindrome(%s) failed, expected %v, got %v", text, expected, got)
		}
	}
}

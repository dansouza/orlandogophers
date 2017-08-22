package main

import (
	"testing"
)

func TestAnagrams(t *testing.T) {
	dict := []string{"BANANA", "APPLE", "DICE", "ANNA MADRIGAL"}
	testTable := map[string][]string{
		"ICED":             []string{"DICE"},
		"ELPAP":            []string{"APPLE"},
		"A MAN AND A GIRL": []string{"ANNA MADRIGAL"},
	}
	for text, expected := range testTable {
		gotList := FindAnagram(text, dict)
		if len(gotList) != len(expected) {
			t.Errorf("FindAnagram(%s, %v) failed, expected %v, got %v", text, dict, expected, gotList)
		}
		for l := 0; l <= len(gotList)-1; l++ {
			if gotList[l] != expected[l] {
				t.Errorf("FindAnagram(%s, %v) failed, expected %v, got %v", text, dict, expected[l], gotList[l])
			}
		}
	}
}

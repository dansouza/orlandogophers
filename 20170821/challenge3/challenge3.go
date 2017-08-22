package main

import (
	"sort"
	"strings"

	"reflect"

	"github.com/cznic/sortutil"
)

// FindAnagram returns an slice of words in the dictionary `dict` that are anagrams of the string `s`
func FindAnagram(s string, dict []string) []string {
	var normalize = func(s string) []rune {
		output := sortutil.RuneSlice{}
		for _, r := range []rune(s) {
			if r != ' ' {
				tmp := []rune(strings.ToLower(string(r)))
				output = append(output, tmp[0])
			}
		}
		sort.Sort(output)
		return output
	}

	output := []string{}
	normalizedInput := normalize(s)

	for _, word := range dict {
		if reflect.DeepEqual(normalizedInput, normalize(word)) {
			output = append(output, word)
		}
	}

	return output
}

func main() {
	// FindAnagram(s string, dict []string)
}

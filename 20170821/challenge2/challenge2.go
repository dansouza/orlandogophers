package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

// IsPalindrome tells whether the string in `text` is a valid UTF-8 palindrome
func IsPalindrome(text string) bool {
	var i int
	runes := []rune(text)

	size := len(runes)
	log.Printf("IsPalindrome(%s), size=%d, runes=%s", text, size, spew.Sdump(runes))
	for i = 0; i <= size/2; i++ {
		if runes[i] != runes[size-1-i] {
			return false
		}
	}
	return true
}

func main() {

}

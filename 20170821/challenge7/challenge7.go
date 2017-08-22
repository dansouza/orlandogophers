// We need a program that picks good random PowerBall numbers.
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Powerball is played by picking 5 unique numbers between 1 and 69 inclusive,
// plus 1 “power play” number between 1 and 26 inclusive. PowerBall picks are
// usually displayed in ascending order plus the power play number.
// The function should return a slice of integers representing the 5 random
// numbers in ascending order plus the power play.
func Powerball() []int {
	var pickRandom = func(from int, to int) int {
		return from + rand.Intn(to-from)
	}

	var i int

	output := []int{0, 0, 0, 0, 0}
	for i = 0; i <= 4; i++ {
	PICKAGAIN:
		r := pickRandom(1, 69)
		for _, v := range output {
			if v == 0 {
				// reached the end of already-picked numbers, break out
				break
			}
			if v == r {
				// picked number already in output, retry
				goto PICKAGAIN
			}
		}
		output[i] = r
	}
	sort.Ints(output)
	// pick a 'power play'
	output = append(output, pickRandom(1, 26))
	return output
}

func main() {
	fmt.Printf("%s", spew.Sdump(Powerball()))
}

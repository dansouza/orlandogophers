package main

import (
	"testing"
)

func TestPowerball(t *testing.T) {
	t.Log("testing 1,000,000 powerball picks...")
	for i := 1; i <= 1000000; i++ {
		pick := Powerball()
		if len(pick) != 6 {
			t.Errorf("pick #%d failed because it doesn't have 6 elements: %v", i, pick)
			break
		}
		got := map[int]struct{}{}
		for idx := 0; idx <= 4; idx++ {
			val := pick[idx]
			if idx > 0 && val < pick[idx-1] {
				t.Errorf("pick #%d failed because it isn't sorted: %v", i, pick)
				break
			}
			if _, ok := got[val]; ok {
				t.Errorf("pick #%d failed because it repeats the number '%d': %v", i, val, pick)
				break
			}
			got[val] = struct{}{}
		}
	}
}

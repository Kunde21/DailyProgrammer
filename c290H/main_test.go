package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name              string
		dog, gopher       animal
		target, intercept XYpair
	}{
		{"Example:", animal{XYpair{10, 10}, 1}, animal{XYpair{1, 10}, 0.25}, XYpair{0, 0}, XYpair{1, 7}},
		{"Challenge", animal{XYpair{5, 3}, 1.2}, animal{XYpair{2, 8}, 0.5}, XYpair{10, 9}, XYpair{0, 0}},
		// TODO: Add test cases.
	}
	for i, v := range tests {
		result := intercept(v.dog, v.gopher, v.target)
		fmt.Println(i, result)
	}
}

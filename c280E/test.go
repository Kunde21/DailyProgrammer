package main

import (
	"fmt"
	"strings"
)

	type hands struct {
		left, right int
	}

	func calc(s string) (*hands, bool) {
		if strings.Contains(s[0:4], "10") || strings.Contains(s[6:], "01") {
			// left hand 						right hand
			return nil, false
		}
		return &hands{
			//  Thumb * 5 + fingers
			left:  strings.Count(s[:4], "1") + 5*strings.Count(s[4:5], "1"),
			right: strings.Count(s[6:], "1") + 5*strings.Count(s[5:6], "1"),
		}, true
	}

func main() {
	for _, v := range []string{"0111011100", "1010010000", "0011101110", "0000110000", "1111110001"} {
		hnds, success := calc(v)
		fmt.Println(v, hnds, success)
		if success {
			fmt.Println(hnds.left*10+hnds.right)
		}
	}
}

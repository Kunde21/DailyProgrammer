package main

import (
	"fmt"
	"math/rand"
)

var bag = []byte{'O', 'I', 'S', 'Z', 'L', 'J', 'T'}

func main() {
	pick := make([]bool, len(bag))
	otp := ""
	r := 0
	for i := 0; i < 50; i++ {
		r = rand.Intn(len(bag))
		if !(pick[r]) {
			otp += string(bag[r])
			pick[r] = true
		}
		if len(pick)%len(bag) == 0 {
			for i := 0; i < len(pick); i++ {
				pick[i] = false
			}
		}
	}
	fmt.Println(otp)
}

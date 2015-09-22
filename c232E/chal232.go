package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
)

func pTest() {
	f, err := ioutil.ReadFile("input2.txt")
	check(err)

	if len(f) < 2 {
		fmt.Println("Not a palindrome.")
		return
	}

	rx, err := regexp.Compile("[^a-z]")
	check(err)

	f = rx.ReplaceAll(bytes.ToLower(f), nil)
	for i, j := 0, len(f)-1; i <= j; i, j = i+1, j-1 {
		if f[i] != f[j] {
			fmt.Println("Not a palindrome.")
			return
		}
	}
	fmt.Println("Palindrome.")
}

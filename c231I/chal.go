package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	d, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	cards := strings.Split(strings.TrimSpace(string(dat)), "\n")
	compl := make(chan struct{})

	for i := 0; i < len(cards)-3; i++ {
		go sets(cards[i:], compl)
	}
	for i := 0; i < len(cards)-3; i++ {
		<-compl
	}
}

func sets(cards []string, compl chan struct{}) {
	for i, c2 := range cards[1:] {
		for _, c3 := range cards[i+2:] {
			if pairs(cards[0], c2, c3) {
				fmt.Println(cards[0], c2, c3)
			}
		}
	}
	compl <- struct{}{}
}

func pairs(s1, s2, s3 string) bool {
	n := 0
	for i := 0; i < len(s1); i++ {
		n = 0
		if s1[i] == s2[i] {
			n++
		}
		if s1[i] == s3[i] {
			n++
		}
		if s2[i] == s3[i] {
			n++
		}
		if n == 1 {
			return false
		}
	}
	return true
}

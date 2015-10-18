package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRead(T *testing.T) {
	b := readInput("input1.txt")
	for i, v := range b {
		fmt.Println(i, ":", string(v), "#")
	}
}

func TestLife1(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	w := readInput("input3.txt")
	rc := make(chan []byte)

	for i := range w {
		go lifeCur(w, i, rc)
	}
	tmp := make([][]byte, len(w))
	for _ = range w {
		t := <-rc
		tmp[t[0]] = t[1:]
	}
	w = tmp

	for _, v := range w {
		fmt.Println(string(v))
	}
}

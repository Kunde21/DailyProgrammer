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

/*
func TestExt(T *testing.T) {
	b := extIfNeed([][]byte{[]byte("AAA"), []byte("AAA"), []byte("AAA")})
	for i, v := range b {
		fmt.Println(i, ":", string(v), "#")
	}
	b = extIfNeed([][]byte{[]byte("A A"), []byte("AAA"), []byte("A A")})
	for i, v := range b {
		fmt.Println(i, ":", string(v), "#")
	}
	b = extIfNeed([][]byte{[]byte(" AA"), []byte("AAA"), []byte("AA ")})
	for i, v := range b {
		fmt.Println(i, ":", string(v), "#")
	}
	b = extIfNeed([][]byte{[]byte("AAA"), []byte(" A "), []byte("AAA")})
	for i, v := range b {
		fmt.Println(i, ":", string(v), "#")
	}
}
//*/

func TestLife1(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	w := readInput("input1.txt")
	giveMeLife(&w, 0, 0)
	for _, v := range w {
		fmt.Println(string(v))
	}

}

func TestLife2(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	w := readInput("input2.txt")
	giveMeLife(&w, 0, 0)
	for _, v := range w {
		fmt.Println(string(v))
	}
}

func TestLife3(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	w := readInput("input3.txt")
	giveMeLife(&w, 0, 0)
	for _, v := range w {
		fmt.Println(string(v))
	}
}
func TestLife4(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	w := readInput("c233I.go")
	giveMeLife(&w, 0, 0)
	for _, v := range w {
		fmt.Println(string(v))
	}
}

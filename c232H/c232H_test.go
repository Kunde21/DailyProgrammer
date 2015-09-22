package main

import (
	"fmt"
	"os"
	"testing"
)

var state [][]int

func init() {

	state = append(state, []int{2, 1})
	state = append(state, []int{2, 1})
}

func TestAdd(T *testing.T) {
	fmt.Println(shortestSplitLine(state, 3, []int{0, len(state) - 1, 0, len(state[0]) - 1, 0}))
}

func TestParse(T *testing.T) {
	fmt.Println(strToInt("4 5 10"))
}

func TestRead(T *testing.T) {
	n, r := readIn("input.txt")
	fmt.Println("N: ", n)
	for _, row := range r {
		for _, cell := range row {
			fmt.Printf("%v ", cell)
		}
		fmt.Println()
	}
}

func TestReadSplit(T *testing.T) {
	n, r := readIn("input.txt")
	res := shortestSplitLine(r, n+1, []int{0, len(r) - 1, 0, len(r) - 1, 0})
	prettyPrint(r, res)
}

func BenchmarkFull(B *testing.B) {
	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}
	B.ReportAllocs()
	for i := 0; i < B.N; i++ {
		n, r := readIn("input.txt")
		res := shortestSplitLine(r, n+1, []int{0, len(r) - 1, 0, len(r) - 1, 0})
		prettyPrint(r, res)
	}
	os.Stdout = tmp
}

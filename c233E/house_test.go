package main

import (
	"fmt"
	"testing"
)

var tmpl [][]byte = [][]byte{[]byte("    * "), []byte(" **** "), []byte(" **** ")}

func TestRead(T *testing.T) {

	b, ln := readFile("input.txt")
	buildFloors(b, ln)
}

func TestFloors(T *testing.T) {
	buildFloors(tmpl, len(tmpl[0]))
}

func TestWalls(T *testing.T) {
	w := buildWalls(buildFloors(tmpl, len(tmpl[0])))
	fmt.Println("Walls")
	for i, v := range w {
		fmt.Println(i, ":", string(v), "#")
	}
	fmt.Println(len(w[0]))
}

func TestRoof(T *testing.T) {
	w := buildWalls(buildFloors(tmpl, len(tmpl[0])))
	w = buildRoof(w)
	for _, v := range w {
		fmt.Println(string(v))
	}
	fmt.Println(len(w[0]))
}

func TestInput1(T *testing.T) {
	w, ln := readFile("input.txt")
	w = addWindows(buildRoof(buildWalls(buildFloors(w, ln))))
	for _, v := range w {
		fmt.Println(string(v))
	}
}

func TestInput2(T *testing.T) {
	w, ln := readFile("input2.txt")
	w = addWindows(buildRoof(buildWalls(buildFloors(w, ln))))
	for _, v := range w {
		fmt.Println(string(v))
	}
}

func TestInputc1(T *testing.T) {
	w, ln := readFile("inputc1.txt")
	w = addWindows(buildRoof(buildWalls(buildFloors(w, ln))))
	for _, v := range w {
		fmt.Println(string(v))
	}
}

func TestInputc2(T *testing.T) {
	w, ln := readFile("inputc2.txt")
	w = addWindows(buildRoof(buildWalls(buildFloors(w, ln))))
	for _, v := range w {
		fmt.Println(string(v))
	}
}

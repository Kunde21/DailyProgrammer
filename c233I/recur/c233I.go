package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	w := readInput("input2.txt")
	giveMeLife(&w, 0, 0)
	for _, v := range w {
		fmt.Println(string(v))
	}
}

func readInput(fn string) (in [][]byte) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)
	in = make([][]byte, 0)
	ln := 0
	for i := 0; scan.Scan(); i++ {
		in = append(in, bytes.Replace(scan.Bytes(), []byte{'\t'}, []byte("        "), -1))
		if len(in[i]) > ln {
			ln = len(in[i])
		}
	}
	for i := 0; i < len(in); i++ {
		tmp := make([]byte, 0, ln)
		tmp = append(tmp, in[i]...)
		in[i] = append(tmp, bytes.Repeat([]byte(" "), ln-len(in[i]))...)
	}
	return

}

func giveMeLife(b *[][]byte, x, y int) {
	if y >= len(*b) {
		return
	}
	if x >= len((*b)[0]) {
		giveMeLife(b, 0, y+1)
		return
	}

	var res byte = 0
	tmp := make([]byte, 0, 9)
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if j >= 0 && i >= 0 && j < len(*b) && i < len((*b)[j]) && (*b)[j][i] != ' ' && (i != x || j != y) {
				res, tmp = res+1, append(tmp, (*b)[j][i])
			}
		}
	}

	switch {
	case (*b)[y][x] != ' ' && res < 2:
		res = byte(' ')
	case (*b)[y][x] != ' ' && res > 3:
		res = byte(' ')
	case (*b)[y][x] == ' ' && res == 3:
		res = tmp[rand.Intn(len(tmp))]
	case (*b)[y][x] != ' ' && res < 4:
		res = (*b)[y][x]
	default:
		res = ' '
	}

	giveMeLife(b, x+1, y)
	(*b)[y][x] = res
}

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

	rc := make(chan []byte)

	for i := range w {
		go lifeCur(w, i, rc)
	}
	tmp := make([][]byte, len(w))
	for _ = range w {
		t := <-rc
		fmt.Println(t[0], string(t[1:]))
		tmp[t[0]] = t[1:]
	}
	w = tmp

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

func lifeCur(b [][]byte, y int, retCh chan []byte) {
	if y >= len(b) {
		retCh <- nil
	}

	ret := make([]byte, len(b[0])+1)
	for x := 0; x < len(b[0]); x++ {
		tmp := make([]byte, 0, 9)
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if j >= 0 && i >= 0 && j < len(b) && i < len(b[j]) && b[j][i] != ' ' && (i != x || j != y) {
					ret[x+1], tmp = ret[x+1]+1, append(tmp, b[j][i])
				}
			}
		}
		switch {
		case b[y][x] != ' ' && ret[x+1] < 2:
			ret[x+1] = byte(' ')
		case b[y][x] != ' ' && ret[x+1] > 3:
			ret[x+1] = byte(' ')
		case b[y][x] == ' ' && ret[x+1] == 3:
			ret[x+1] = tmp[rand.Intn(len(tmp))]
		case b[y][x] != ' ' && ret[x+1] < 4:
			ret[x+1] = b[y][x]
		default:
			ret[x+1] = ' '
		}
	}
	ret[0] = byte(y)
	retCh <- ret
	return
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	w, ln := readFile("inputc2.txt")
	w = addWindows(buildRoof(buildWalls(buildFloors(w, ln))))
	for _, v := range w {
		fmt.Println(string(v))
	}
}

func readFile(fn string) (in [][]byte, lmx int) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Scan()

	ln, err := strconv.Atoi(strings.TrimSpace(s.Text()))
	if err != nil {
		log.Fatal(err)
	}

	in = make([][]byte, ln)

	for i := 0; i < ln; i++ {
		s.Scan()
		in[i] = append(s.Bytes(), byte(' '))
		in[i] = append([]byte{' '}, in[i]...)
		if l := len(in[i]); l > lmx {
			lmx = l
		}
	}

	return
}

func buildFloors(b [][]byte, ln int) (ret [][]byte) {

	ret = make([][]byte, len(b)*2+1)

	for i := 0; i < len(b); i++ {
		v := make([]byte, ln, ln)
		copy(v, b[i])
		v = bytes.Replace(v, []byte{0}, []byte{' '}, -1)
		v = bytes.Replace(v, []byte{' '}, bytes.Repeat([]byte{' '}, 4), -1)
		v = bytes.Replace(v, []byte{'*'}, []byte("+---"), -1)
		v = bytes.Replace(v, []byte("- "), []byte("-+"), -1)
		ret[i*2] = v
		ret[i*2+1] = bytes.Repeat([]byte{' '}, ln*4)
	}

	ret[len(ret)-1] = make([]byte, len(ret[0]))
	copy(ret[len(ret)-1], ret[len(ret)-3])

	for i := len(ret) - 5; i >= 0; i -= 2 {
		for j, v := range ret[i] {
			if v == '-' && ret[i+2][j] == '-' {
				ret[i+2][j] = ' '
			}
		}
		ret[i+2] = bytes.Replace(ret[i+2], []byte(" + "), []byte("   "), -1)
		ret[i+2] = bytes.Replace(ret[i+2], []byte("-+-"), []byte("---"), -1)
	}
	ret[len(ret)-1] = bytes.Replace(ret[len(ret)-1], []byte("-+-"), []byte("---"), -1)
	ret[0] = bytes.Replace(ret[0], []byte("-+-"), []byte("---"), -1)

	return ret

}

func buildWalls(b [][]byte) [][]byte {

	for i := 0; i < len(b)-2; i += 2 {
		for j, v := range b[i] {
			if v != '+' {
				continue
			}
			for k := i + 1; k < len(b)-1; k++ {
				if b[k][j] == '+' || i > 0 && b[i-1][j] == '|' {
					break
				}
				b[k][j] = '|'
			}
		}
	}
	//Remove parse padding
	for i := 0; i < len(b); i++ {
		b[i] = b[i][4 : len(b[i])-3]
	}
	return b
}

func buildRoof(b [][]byte) [][]byte {

	//Height adjust
	add := 0
	for i := 0; i < len(b)-2; i += 2 {
		for k := 0; bytes.IndexRune(b[i][k:], '+') >= 0; k += 1 {
			k += bytes.IndexRune(b[i][k:], '+') + 1
			tmp := bytes.IndexRune(b[i][k+1:], '+')
			if tmp/2-i+1 > add {
				add = tmp/2 - i + 1
			}
			k += tmp + 1
		}
	}

	pad := make([][]byte, add)
	for i := 0; i < add; i++ {
		pad[i] = bytes.Repeat([]byte{' '}, len(b[0]))
	}
	b = append(pad, b...)

	for i := len(b) - 2; i > 0; i-- {
		for j, v := range b[i] {
			switch {
			case v == '+' && j == 0:
				b[i-1][j+1] = '/'
			case v == '+' && j == len(b[i]):
				b[i-1][j-1] = '\\'
			case v == '+' && b[i][j-1] == '-':
				b[i-1][j-1] = '\\'
			case v == '+' && b[i][j+1] == '-':
				b[i-1][j+1] = '/'
			case v == '/' && b[i][j+2] == '\\':
				b[i-1][j+1] = 'A'
			case v == '/' && b[i][j+2] != '\\':
				b[i-1][j+1] = '/'
			case v == '\\' && b[i][j-2] != '/':
				b[i-1][j-1] = '\\'
			}
		}
	}
	return b
}

func addWindows(b [][]byte) [][]byte {
	rand.Seed(time.Now().UnixNano())
	door := rand.Intn(len(b[0]) - 4)

	b[len(b)-2][door+1] = '|'
	b[len(b)-2][door+3] = '|'

	var w bool
	for i := len(b) - 2; i >= 1; i -= 2 {
		w = false
		for j := 0; j < len(b[i])-1; j++ {
			if b[i][j] != '|' && w && j%4 != 0 && j%2 == 0 && rand.Intn(2) == 0 {
				b[i][j] = 'o'
			}
			if b[i][j] == '|' {
				w = !w
			}
		}
	}
	return b
}

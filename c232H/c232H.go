package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, r := readIn("input.txt")
	res := shortestSplitLine(r, n+1, []int{0, len(r) - 1, 0, len(r) - 1, 0})
	prettyPrint(r, res)
}

func shortestSplitLine(state [][]int, N int, area []int) [][]int {
	sum, h, v := blockSums(state, area)

	if N == 1 {
		area[4] = sum
		return [][]int{area}
	}

	a, b := N/2, (N+1)/2
	ratio := float64(a) / float64(a+b)

	var (
		tsum, div1, div2 int
		horiz            bool
		minratio         float64
	)

	for n, i := range v {
		tsum += i
		if tmp := math.Abs(float64(tsum)/float64(sum) - ratio); n == 0 || tmp < minratio {
			div1, div2, minratio = n+1, 0, tmp
		}
		if tmp := math.Abs(1 - float64(tsum)/float64(sum) - ratio); tmp < minratio {
			div2, div1, minratio = n+1, 0, tmp
		}
	}
	tsum = 0
	for n, i := range h {
		tsum += i
		if tmp := math.Abs(float64(tsum)/float64(sum) - ratio); tmp < minratio {
			div1, div2, minratio, horiz = n+1, 0, tmp, true
		}
		if tmp := math.Abs(1 - float64(tsum)/float64(sum) - ratio); tmp < minratio {
			div2, div1, minratio, horiz = n+1, 0, tmp, true
		}
	}

	retchan := make(chan [][]int)
	switch {
	case horiz && div1 > 0:
		go func(ch chan [][]int) {
			retchan <- shortestSplitLine(state, a, []int{area[0], area[0] + div1 - 1, area[2], area[3], 0})
		}(retchan)
		return append(<-retchan, shortestSplitLine(state, b, []int{area[0] + div1, area[1], area[2], area[3], 0})...)
	case horiz && div2 > 0:
		go func(ch chan [][]int) {
			retchan <- shortestSplitLine(state, b, []int{area[0], area[0] + div2 - 1, area[2], area[3], 0})
		}(retchan)
		return append(<-retchan, shortestSplitLine(state, a, []int{area[0] + div2, area[1], area[2], area[3], 0})...)
	case div1 > 0:
		go func(ch chan [][]int) {
			retchan <- shortestSplitLine(state, a, []int{area[0], area[1], area[2], area[2] + div1 - 1, 0})
		}(retchan)
		return append(<-retchan, shortestSplitLine(state, b, []int{area[0], area[1], area[2] + div1, area[3], 0})...)
	case div2 > 0:
		go func(ch chan [][]int) {
			retchan <- shortestSplitLine(state, b, []int{area[0], area[1], area[2], area[2] + div2 - 1, 0})
		}(retchan)
		return append(<-retchan, shortestSplitLine(state, a, []int{area[0], area[1], area[2] + div2, area[3], 0})...)
	default:
		return nil
	}
}

func blockSums(state [][]int, area []int) (sum int, h, v []int) {
	h, v = make([]int, area[1]-area[0]+1), make([]int, area[3]-area[2]+1)
	sum = 0
	for i := area[0]; i <= area[1]; i++ {
		for j := area[2]; j <= area[3]; j++ {
			h[i-area[0]] = h[i-area[0]] + state[i][j]
			v[j-area[2]] = v[j-area[2]] + state[i][j]
		}
		sum += h[i-area[0]]
	}
	return
}

func readIn(fn string) (n int, st [][]int) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	scan.Scan() //First Line (N, row, col)
	info := strToInt(scan.Text())
	if len(info) != 3 {
		//Fail fast if poorly formatted
		panic("Info line missing in file")
	}
	n = info[0]
	st = make([][]int, 0, info[1])

	for i := 0; i < info[1] && scan.Scan(); i++ {
		t := strToInt(scan.Text())
		if len(t) != info[2] {
			panic("Improper format")
		}
		st = append(st, t)
	}

	return
}

func strToInt(inp string) (ret []int) {
	s := strings.SplitN(strings.TrimSpace(inp), " ", -1)
	ret = make([]int, len(s))
	var t int64
	for i, v := range s {
		t, _ = strconv.ParseInt(v, 10, 0)
		ret[i] = int(t)
	}
	return
}

func prettyPrint(state [][]int, res [][]int) {
	for i, v := range res {
		fmt.Printf("Section: %v   (%v,%v)\t(%v,%v)    \tPopulation: %v\n", i, v[0], v[2], v[1], v[3], v[4])
	}

	fmt.Printf("\n%s\n", strings.Repeat("-", len(state[0])*4+1))

	for i, v := range state {
		ln, br := splits(res, i, len(state[0]))
		for j, c := range v {
			if ln[j] {
				fmt.Printf("| %v ", c)
			} else {
				fmt.Printf("  %v ", c)
			}
		}
		fmt.Printf("|\n%s|\n", br)

	}
}

func splits(res [][]int, row, len int) (ln []bool, s string) {
	ln = make([]bool, len)
	br := make([]byte, len*4)
	for _, r := range res {
		if row <= r[1] && row >= r[0] {
			ln[r[2]] = true
			br[r[2]*4] = byte('|')
			if row == r[1] {
				br[r[2]*4] = byte('*')
				br = bytes.Replace(br, []byte{'*'}, bytes.Repeat([]byte{'-'}, 4*(r[3]-r[2]+1)), 1)
			}
		}
	}
	br = bytes.Replace(br, []byte{0}, []byte{' '}, -1)

	s = string(br[:len*4])

	return
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type loc struct {
	pts  []byte
	x, y float64
}

func main() {
	f, err := os.Open("bonusGen.txt")
	check(err)
	l := fillArr(bufio.NewScanner(f))
	f.Close()
	fmt.Println(len(l))
	i, j, d := findNearest(l)
	fmt.Println(string(l[i].pts), string(l[j].pts), d)
}

func runfile(fn string) {
	f, err := os.Open(fn)
	check(err)
	l := fillArr(bufio.NewScanner(f))
	f.Close()
	fmt.Println(len(l))
	i, j, d := findNearest(l)
	fmt.Println(string(l[i].pts), string(l[j].pts), d)
}

func findNearest(la []loc) (a, b int, d float64) {
	if len(la) < 2 {
		return
	}
	a, b, d = 0, 1, la[0].dist(la[1])
	for i, l := range la {
		for j := i + 1; j < len(la) && la[j].x-l.x < d; j++ {
			if td := l.dist(la[j]); td < d {
				a, b, d = i, j, td
			}
		}
	}
	return
}

func findNearestCt(la []loc) (a, b, ct int, d float64) {
	if len(la) < 2 {
		return
	}
	a, b, d = 0, 1, la[0].dist(la[1])
	for i, l := range la {
		for j := i + 1; j < len(la) && math.Abs(l.x-la[j].x) < d; j++ {
			ct++
			if td := l.dist(la[j]); td < d {
				a, b, d = i, j, td
			}
		}
	}
	return
}

func (l1 loc) dist(l2 loc) float64 {
	return math.Sqrt((l1.x-l2.x)*(l1.x-l2.x) + (l1.y-l2.y)*(l1.y-l2.y))
}

func fillArr(s *bufio.Scanner) (l []loc) {
	var err error
	for s.Scan() {
		if bytes.Count(s.Bytes(), []byte{','}) != 1 {
			if len(l) == 0 {
				ln, err := strconv.ParseInt(string(bytes.TrimSpace(s.Bytes())), 10, 0)
				check(err)
				l = make([]loc, 0, ln+10)
			}
			continue
		}

		t1 := make([]byte, len(s.Bytes()))
		copy(t1, s.Bytes())
		tmploc := loc{t1, 0, 0}
		tmp := bytes.SplitN(bytes.Trim(tmploc.pts, "() "), []byte{','}, 3)
		tmploc.x, err = strconv.ParseFloat(string(bytes.TrimSpace(tmp[0])), 64)
		check(err)
		tmploc.y, err = strconv.ParseFloat(string(bytes.TrimSpace(tmp[1])), 64)
		check(err)
		l = append(l, tmploc)
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	sort.Sort(locA(l))
	return
}

func check(err error) {
	if err != nil {
		log.Fatal("Error", err)
	}
}

type locA []loc

func (lc locA) Len() int {
	return len(lc)
}

func (lc locA) Less(i, j int) bool {
	switch {
	case lc[i].x < lc[j].x:
		fallthrough
	case lc[i].x == lc[j].x && lc[i].y < lc[j].y:
		return true
	default:
		return false
	}
}

func (l locA) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

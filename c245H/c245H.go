package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	root      *Node
	allRngs   []*Range
	unkC      uint32
	treeDepth int = 16
	sNum      int = 25
)

func main() {
}

func loadSrc(f io.Reader) {
	s := bufio.NewScanner(f)
	root, allRngs, unkC = NewNode(0, 0, 1<<32-1), make([]*Range, 0, 1000), 0
	for s.Scan() {
		tmp := bytes.SplitN(bytes.TrimSpace(s.Bytes()), []byte(" "), 3)
		min, max := convertIP(tmp[0]), convertIP(tmp[1])
		root.load(&Range{min, max, max - min, 0, string(tmp[2])})
	}
}

func loadStats(f io.Reader) {
	s := bufio.NewScanner(f)
	wg := new(sync.WaitGroup)
	tmp := make([][]byte, sNum)
	unkC = 0
	srch := func(b [][]byte) {
		for _, v := range b {
			t := root.search(convertIP(v))
			if t == nil {
				atomic.AddUint32(&unkC, 1)
			} else {
				atomic.AddUint32(&t.count, 1)
			}
		}
		wg.Done()
	}

	i := 0
	for s.Scan() {
		tmp[i] = make([]byte, len(s.Bytes()))
		copy(tmp[i], s.Bytes())
		i++
		if i == sNum {
			wg.Add(1)
			go srch(tmp)
			i, tmp = 0, make([][]byte, sNum)
		}
	}
	if i > 0 {
		wg.Add(1)
		go srch(tmp[:i])
	}
	wg.Wait()
}

//*****************Interval Tree structures***************
type Range struct {
	min, max, rng, count uint32
	name                 string
}

func (r Range) String() string {
	return fmt.Sprint(numToIp(r.min), " - ", numToIp(r.max), " ", r.rng, " ", r.name)
}

type Node struct {
	mid  uint32
	chld []*Node
	rngs []*Range
}

func NewNode(depth int, min, max uint32) (n *Node) {
	mid := uint32((uint64(max) + uint64(min)) / 2)
	if depth < treeDepth {
		return &Node{mid,
			[]*Node{
				NewNode(depth+1, min, mid),
				NewNode(depth+1, mid, max),
			}, make([]*Range, 0, 200)}
	}
	return &Node{mid, []*Node{}, make([]*Range, 0, 200)}
}

func (n *Node) load(r *Range) {
	if len(n.chld) == 0 {
		n.rngs = append(n.rngs, r)
		allRngs = append(allRngs, r)
		return
	}

	switch {
	case r.max < n.mid:
		n.chld[0].load(r)
	case r.min > n.mid:
		n.chld[1].load(r)
	default:
		n.rngs = append(n.rngs, r)
		allRngs = append(allRngs, r)
	}
}

func (n *Node) sort() {
	sort.Sort(RngA(n.rngs))
	if len(n.chld) != 0 {
		n.chld[0].sort()
		n.chld[1].sort()
	}
}

func (n *Node) search(num uint32) (r *Range) {
	switch {
	case len(n.chld) == 0:
		for _, v := range n.rngs {
			if v.min <= num && num <= v.max {
				return v
			}
		}
		return nil
	case num < n.mid:
		r = n.chld[0].search(num)
	default:
		r = n.chld[1].search(num)
	}
	for _, v := range n.rngs {
		switch {
		case r != nil && v.rng > r.rng:
			return r
		case v.min <= num && num <= v.max:
			return v
		}
	}
	return
}

//***************Conversion Helpers*******************************
func convertIP(ip []byte) (val uint32) {
	qtrs := bytes.Split(bytes.TrimSpace(ip), []byte("."))
	if len(qtrs) != 4 {
		return 0
	}
	for i, v := range qtrs {
		t, _ := strconv.Atoi(string(v))
		val |= uint32(t << uint8((3-i)*8))
	}
	return val
}

func numToIp(inp uint32) (s string) {
	for i := 0; i < 4; i++ {
		s += strconv.Itoa(int(inp >> (uint8((3 - i) * 8)) & 255))
		if i < 3 {
			s += "."
		}
	}
	return s
}

//************************Result Output***************************
func output() {
	sort.Sort(RngB(allRngs))
	for _, v := range allRngs {
		if v.count > 0 {
			fmt.Println(v.count, v.name)
		}
	}
	if unkC > 0 {
		fmt.Println(unkC, "<unknown>")
	}
}

//************************Sort helpers****************************
type RngA []*Range //Range width sort

func (r RngA) Len() int {
	return len(r)
}

func (r RngA) Less(i, j int) bool {
	if r[i].rng < r[j].rng {
		return true
	}
	return false
}

func (r RngA) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type RngB []*Range //Resulting count sort(Print order)

func (r RngB) Len() int {
	return len(r)
}

func (r RngB) Less(i, j int) bool {
	if r[i].count > r[j].count {
		return true
	}
	return false
}

func (r RngB) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

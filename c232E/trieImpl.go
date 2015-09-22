package main

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
	"time"
	"unsafe"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	nf, nr := newNode(), newNode()
	load(nf, nr)

	var buf []byte
	ch := make(chan struct{})
	ct := make(chan struct{}, 50)
	snd := func(c chan struct{}, buf []byte, i int64, ct chan struct{}) {
		palinTrie(nf.ptrs[i], nr.ptrs[i], 1, 0, 0, buf, nf, nr, ct)
		c <- struct{}{}
	}

	gor := 0
	for i := nf.minIdx; i <= nf.maxIdx; i++ {
		buf = make([]byte, 50)
		buf[0] = byte(i + 'a')
		go snd(ch, buf, i, ct)
		gor++
	}

	cnt := 0
	for gor > 0 {
		select {
		case <-ct:
			cnt++
		case <-ch:
			gor--
		}
	}
	for {
		select {
		case <-ct:
			cnt++
		default:
			fmt.Println("Palindromes", cnt)
			close(ct)
			close(ch)
			return
		}
	}
}

type Node struct {
	word           int64
	minIdx, maxIdx int64
	ptrs           []*Node
}

func newNode() (n *Node) {
	n = new(Node)
	n.ptrs = make([]*Node, 26)
	n.minIdx, n.maxIdx = 25, 0
	return
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type wrk struct {
	tmp    unsafe.Pointer
	mn, mx int64
	i      *int64
	tmpN   Node
}

func (node *Node) addChar(c *byte, w *wrk) *Node {
	*w.i = int64(*c - 'a')
	if *w.i < 0 || *w.i > 25 {
		return node
	}
	if w.tmp = (atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&(node.ptrs[*w.i]))))); w.tmp == nil {
		atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer((&node.ptrs[*w.i]))), w.tmp, unsafe.Pointer(newNode()))
		w.mn, w.mx = atomic.LoadInt64(&(node.minIdx)), atomic.LoadInt64(&(node.maxIdx))
		for {
			switch {
			case w.mn > *w.i:
				if !atomic.CompareAndSwapInt64(&(node.minIdx), w.mn, *w.i) {
					w.mn = atomic.LoadInt64(&(node.minIdx))
				} else {
					w.mn = *w.i
				}
			case w.mx < *w.i:
				if !atomic.CompareAndSwapInt64(&(node.maxIdx), w.mx, *w.i) {
					w.mx = atomic.LoadInt64(&(node.maxIdx))
				} else {
					w.mx = *w.i
				}
			default:
				return node.ptrs[*w.i]
			}
		}
	}
	return node.ptrs[*w.i]
}

type inp struct {
	wd int64
	rn []byte
}

func loadWord(chi, cho chan inp, tr, rtr *Node) {

	var nf, nr *Node
	var ln, j int
	i := new(wrk)
	i.i = new(int64)
	for c := range chi {
		nf, nr = tr, rtr
		ln = len(c.rn) - 1
		for j = 0; j < ln; j++ {
			nf = nf.addChar(&c.rn[j], i)
			nr = nr.addChar(&c.rn[ln-j-1], i)
		}
		nf.word, nr.word = c.wd, c.wd

		select {
		case cho <- c:
		default:
		}
	}
}

func load(tr, rtr *Node) {
	f, err := os.Open("enable1.txt")
	defer f.Close()
	check(err)

	var w []byte
	var i, wd int64
	var c inp

	r := bufio.NewReader(f)

	chi, cho := make(chan inp, 30), make(chan inp, 30)

	for i = 0; i < 28; i++ {
		go loadWord(chi, cho, tr, rtr)
	}

	for w, err = r.ReadBytes('\n'); err == nil; w, err = r.ReadBytes('\n') {
		wd++
		select {
		case c = <-cho:
			c.wd, c.rn = wd, w
		default:
			c = inp{wd, w}
		}

		chi <- c
	}

	for {
		select {
		case <-cho:
		case <-time.After(10 * time.Millisecond):
			close(cho)
			close(chi)
			return
		}
	}
}

func palinTrie(tr, rtr *Node, dpth, lastW, lastWr int64, buf []byte, t, r *Node, num chan struct{}) {

	var nx, rnx *Node
Loop:
	for i := tr.minIdx; i <= tr.maxIdx; i++ {
		nx, rnx = tr.ptrs[i], rtr.ptrs[i]

		if nx == nil || rnx == nil {
			continue Loop
		}

		buf[dpth] = byte(i + 'a')

		if nx.word != 0 && lastW == 0 {
			buf[dpth+1] = ' '
			palinTrie(t, rnx, dpth+2, nx.word, lastWr, buf, t, r, num)
		}
		if rnx.word != 0 && lastWr == 0 {
			palinTrie(nx, r, dpth+1, lastW, rnx.word, buf, t, r, num)
		}
		if nx.word != 0 && rnx.word != 0 {
			switch {
			case lastW == rnx.word && nx.word == lastWr && nx.word != lastW:
				num <- struct{}{}
			case lastW == 0 && lastWr == 0 && nx.word != lastWr:
				buf[dpth+1] = ' '
				palinTrie(t, r, dpth+2, nx.word, rnx.word, buf, t, r, num)
			}
		}
		palinTrie(nx, rnx, dpth+1, lastW, lastWr, buf, t, r, num)
	}
}

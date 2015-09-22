package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var nfg, nrg *Node

func init() {
	nfg, nrg = newNode(), newNode()
	load(nfg, nrg)
}

func TestNewNode(t *testing.T) {
	_ = newNode()
}

func TestPTest(t *testing.T) {
	pTest()
}

func BenchmarkPTest(b *testing.B) {

	b.ReportAllocs()
	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}

	for i := 0; i < b.N; i++ {
		pTest()
	}
	os.Stdout = tmp
}

func TestAddChar(t *testing.T) {
	n := newNode()
	i := new(wrk)
	i.i = new(int64)
	a := byte('a')
	r := n.addChar(&a, i)
	if n.ptrs[0] != r {
		log.Fatal("Incorrect Placement")
	}
}

func TestLoad(t *testing.T) {
	nf, nr := newNode(), newNode()
	load(nf, nr)
}

func BenchmarkLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		nf, nr := newNode(), newNode()
		load(nf, nr)
	}
}

func TestPalinTrie(t *testing.T) {
	var buf []byte

	ch := make(chan struct{})
	ct := make(chan struct{}, 50)

	snd := func(c chan struct{}, buf []byte, i int64, ct chan struct{}) {
		palinTrie(nfg.ptrs[i], nrg.ptrs[i], 1, 0, 0, buf, nfg, nrg, ct)
		c <- struct{}{}
	}

	gor := 0
	for i := nfg.minIdx; i <= nfg.maxIdx; i++ {
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
		case <-time.After(5 * time.Millisecond):
			fmt.Println("Palindromes", cnt)
			close(ct)
			close(ch)
			return
		}
	}
}

func BenchmarkPalinTrie(b *testing.B) {

	b.ReportAllocs()
	b.SetParallelism(1)

	var buf []byte

	snd := func(c chan struct{}, buf []byte, i int64, ct chan struct{}) {
		palinTrie(nfg.ptrs[i], nrg.ptrs[i], 1, 0, 0, buf, nfg, nrg, ct)
		c <- struct{}{}
	}

	var ch, ct chan struct{}
	gor, cnt := 0, 0
	for bi := 0; bi < b.N; bi++ {
		ch = make(chan struct{})
		ct = make(chan struct{}, 50)

		for i := nfg.minIdx; i <= nfg.maxIdx; i++ {
			buf = make([]byte, 100)
			buf[0] = byte(i + 'a')
			go snd(ch, buf, i, ct)
			gor++
		}

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
				if cnt != 6501 {
					log.Fatal("Palindromes", cnt)
				}
				close(ct)
				close(ch)
				return
			}
		}
	}
}

func TestFull(t *testing.T) {
	nfg, nrg := newNode(), newNode()
	load(nfg, nrg)

	var buf []byte

	ch := make(chan struct{})
	ct := make(chan struct{}, 50)

	snd := func(c chan struct{}, buf []byte, i int64, ct chan struct{}) {
		palinTrie(nfg.ptrs[i], nrg.ptrs[i], 1, 0, 0, buf, nfg, nrg, ct)
		c <- struct{}{}
	}

	gor := 0
	for i := nfg.minIdx; i <= nfg.maxIdx; i++ {
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
			if cnt != 6501 {
				log.Fatal("Palindromes", cnt)
			}
			close(ct)
			close(ch)
			return
		}
	}
}

func benchmarkFull(b *testing.B) {
	for nm := 0; nm < b.N; nm++ {
		nfg, nrg := newNode(), newNode()
		load(nfg, nrg)

		var buf []byte

		ch := make(chan struct{})
		ct := make(chan struct{}, 50)

		snd := func(c chan struct{}, buf []byte, i int64, ct chan struct{}) {
			palinTrie(nfg.ptrs[i], nrg.ptrs[i], 1, 0, 0, buf, nfg, nrg, ct)
			c <- struct{}{}
		}

		gor := 0
		for i := nfg.minIdx; i <= nfg.maxIdx; i++ {
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
				if cnt != 6501 {
					log.Fatal("Palindromes", cnt)
				}
				close(ct)
				close(ch)
				return
			}
		}

	}
}

func BenchmarkFullAlloc1(b *testing.B) {
	b.SetParallelism(1)
	b.ReportAllocs()
	benchmarkFull(b)
}

func BenchmarkFullAlloc2(b *testing.B) {
	b.SetParallelism(2)
	b.ReportAllocs()
	benchmarkFull(b)
}

func BenchmarkFullAlloc3(b *testing.B) {
	b.SetParallelism(16)
	b.ReportAllocs()
	benchmarkFull(b)
}

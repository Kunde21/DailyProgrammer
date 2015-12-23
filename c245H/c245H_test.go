package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSrch(t *testing.T) {
	f, err := os.Open("tstIPs.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	loadSrc(f)
	f, err = os.Open("tstLk.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	loadStats(f)
	output()
}

func TestConv(t *testing.T) {
	tst := []struct {
		ip  []byte
		err bool
		num uint32
	}{
		{[]byte("127.0.0.1"), false, 2130706433},
		{[]byte("1.0.0.0"), false, 1 << 24},
		{[]byte("0.1.0.0"), false, 1 << 16},
		{[]byte("0.0.1.0"), false, 1 << 8},
		{[]byte("0.0.0.1"), false, 1},
	}

	for _, v := range tst {
		if tm := convertIP(v.ip); (tm == 0) != v.err || tm != v.num {
			t.Log("IP:", string(v.ip), "Num:", tm, v.num, "Err:", v.err)
			t.Fail()
		}
	}
}

func TestRead(t *testing.T) {
	f, err := os.Open("ips500.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	loadSrc(f)

}

func TestCalc(t *testing.T) {
	f, err := os.Open("tstIPs.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	loadSrc(f)
	f, err = os.Open("tstLk.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	loadStats(f)
	tm := time.Now()
	output()
	fmt.Println(time.Since(tm))
}

func BenchmarkRead(b *testing.B) {
	f, err := os.Open("ips1mil.txt")
	if err != nil {
		b.Log(err)
		b.FailNow()
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		loadSrc(f)
		f.Seek(0, 0)
	}
}

func Benchmark300k10k(b *testing.B) {
	f, err := os.Open("ips300k.txt")
	if err != nil {
		b.Log(err)
		b.FailNow()
	}
	loadSrc(f)
	b.StartTimer()
	fs, err := os.Open("query10k.txt")
	if err != nil {
		b.Log(err)
		b.FailNow()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loadStats(fs)
		b.StopTimer()

		for _, v := range allRngs {
			v.count = 0
		}
		unkC = 0
		fs.Seek(0, 0)
		b.StartTimer()
	}
}

func Benchmark1mil10k(b *testing.B) {
	f, err := os.Open("ips1mil.txt")
	if err != nil {
		b.Log(err)
		b.FailNow()
	}
	loadSrc(f)
	b.StartTimer()
	fs, err := os.Open("query10k.txt")
	if err != nil {
		b.Log(err)
		b.FailNow()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loadStats(fs)
		b.StopTimer()

		for _, v := range allRngs {
			v.count = 0
		}
		unkC = 0
		fs.Seek(0, 0)
		b.StartTimer()
	}
}

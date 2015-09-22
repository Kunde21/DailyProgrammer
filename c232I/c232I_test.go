package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"testing"
)

func TestFillArr(T *testing.T) {
	f, err := os.Open("inp.txt")
	if err != nil {
		log.Fatal("Open File error", err)
	}
	defer f.Close()
	l := fillArr(bufio.NewScanner(f))
	fmt.Println(len(l))
}

func TestFindNearestTest(T *testing.T) {
	f, err := os.Open("bonusGen.txt")
	if err != nil {
		log.Fatal("Open File error", err)
	}
	defer f.Close()
	l := fillArr(bufio.NewScanner(f))
	i, j, ct, d := findNearestCt(l)
	fmt.Println(i, j)
	fmt.Println("Point A: ", string(l[i].pts))
	fmt.Println("Point B: ", string(l[j].pts))
	fmt.Println("Distance: ", d)
	fmt.Println("Comparisons: ", ct)

	i, j, d = findNearest(l)
	fmt.Println(i, j)
	fmt.Println("Point A: ", string(l[i].pts))
	fmt.Println("Point B: ", string(l[j].pts))
	fmt.Println("Distance: ", d)
}

func TestFindNearestChal(T *testing.T) {
	f, err := os.Open("inp2.txt")
	if err != nil {
		log.Fatal("Open File error", err)
	}
	defer f.Close()
	l := fillArr(bufio.NewScanner(f))
	i, j, ct, d := findNearestCt(l)
	fmt.Println(i, j)
	fmt.Println("Point A: ", string(l[i].pts))
	fmt.Println("Point B: ", string(l[j].pts))
	fmt.Println("Distance: ", d)
	fmt.Println("Comparisons: ", ct)

	i, j, d = findNearest(l)
	fmt.Println(i, j)
	fmt.Println("Point A: ", string(l[i].pts))
	fmt.Println("Point B: ", string(l[j].pts))
	fmt.Println("Distance: ", d)
}

func TestFindNearestTenK(T *testing.T) {
	f, err := os.Open("bonus.txt")
	if err != nil {
		log.Fatal("Open File error", err)
	}
	defer f.Close()
	l := fillArr(bufio.NewScanner(f))
	fmt.Println(len(l))
	i, j, ct, d := findNearestCt(l)
	fmt.Println(i, j)
	fmt.Println("Point A: ", string(l[i].pts))
	fmt.Println("Point B: ", string(l[j].pts))
	fmt.Println("Distance: ", d)
	fmt.Println("Comparisons: ", ct)

	i, j, d = findNearest(l)
	fmt.Println(i, j)
	fmt.Println("Point A: ", string(l[i].pts))
	fmt.Println("Point B: ", string(l[j].pts))
	fmt.Println("Distance: ", d)
}

func TestFullChal(T *testing.T) {
	main()
}

/*
func TestFix(T *testing.T) {
	r, err := os.Open("bonus2.txt")
	check(err)
	w, err := os.OpenFile("bonus.txt", os.O_WRONLY, 0666)
	check(err)
	s := bufio.NewScanner(r)
	for s.Scan() {
		if bytes.Count(s.Bytes(), []byte{' '}) != 1 {
			fmt.Fprintf(w, "%s\n", s.Text())
			continue
		}
		fmt.Fprintf(w, "(%s)\n", strings.Replace(s.Text(), " ", ", ", 1))
	}
}
//*

func TestGenerate(T *testing.T) {
	w, err := os.OpenFile("bonusGen.txt", (os.O_CREATE | os.O_WRONLY), 0666)
	check(err)
	fmt.Fprintf(w, "%v\n", 10000000)
	for i := 0; i < 10000000; i++ {
		fmt.Fprintf(w, "(%v, %v)\n", 10*rand.Float64(), 10*rand.Float64())
	}
}

//*/

func BenchmarkFullChal5k(b *testing.B) {

	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main()
	}
	os.Stdout = tmp
}

func BenchmarkFullChal10K(b *testing.B) {

	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		runfile("bonus.txt")
	}
	os.Stdout = tmp
}

func BenchmarkFullChal1M(b *testing.B) {

	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		runfile("bonusGen.txt")
	}
	os.Stdout = tmp
}

/*
var lst []loc

func init() {
	f, err := os.Open("bonusGen.txt")
	if err != nil {
		log.Fatal("Open File error", err)
	}
	lst = fillArr(bufio.NewScanner(f))
	sort.Sort(locA(lst))
	f.Close()
}

func BenchmarkSort1M(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, _ = findNearest(lst)
	}
} //*/

func BenchmarkLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f, err := os.Open("bonusGen.txt")
		if err != nil {
			log.Fatal("Open File error", err)
		}
		_ = fillArr(bufio.NewScanner(f))
		f.Close()
	}
}

func BenchmarkLoadSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f, err := os.Open("bonusGen.txt")
		if err != nil {
			log.Fatal("Open File error", err)
		}
		l := fillArr(bufio.NewScanner(f))
		sort.Sort(locA(l))
		f.Close()
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestThrees(t *testing.T) {
	threes(3)
	fmt.Println()
	threes(4)
	fmt.Println()
	threes(5)
}

func TestThreesRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		threes(rand.Intn(3000))
		fmt.Println()
	}
}

func BenchmarkChallenge(b *testing.B) {
	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		threes(31337357)
	}
	os.Stdout = tmp
}

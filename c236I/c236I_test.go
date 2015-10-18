package main

import (
	"os"
	"testing"
)

func Testmain(T *testing.T) {
	main()
}

func BenchmarkMain(B *testing.B) {
	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}

	for i := 0; i < B.N; i++ {
		main()
	}

	os.Stdout = tmp
}

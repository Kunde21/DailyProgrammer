package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	ch := make(chan int64)
	go fibGen(ch)

	f, err := os.Open("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var inp, tmp, mod int64
	fib := make([]int64, 1, 25)
	fib[0] = <-ch
	fmt.Fscanln(f, &inp)
	for inp >= 0 {
		minMult := inp
		for i := 0; fib[i] <= inp; i++ {
			if fib[i] != 0 {
				tmp, mod = inp/fib[i], inp%fib[i]
				if mod == 0 && tmp < minMult {
					minMult = tmp
				}
			}
			if i == len(fib)-1 {
				fib = append(fib, <-ch)
			}
		}
		if minMult == 0 {
			minMult = 1
		}
		fmt.Println("\nFor num: ", inp, " i =", minMult)
		fmt.Print("Series: ")
		for i := 0; i < len(fib) && minMult*fib[i] <= inp; i++ {
			fmt.Print(minMult*fib[i], " ")
		}
		fmt.Println()
		fmt.Fscanln(f, &inp)
	}
}

func fibGen(ch chan<- int64) {
	a, b := int64(0), int64(1)

	for {
		ch <- a
		a += b
		a, b = b, a
	}
}

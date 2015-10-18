package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	ch := make(chan *big.Int)
	go fibGen(ch)

	f, err := os.Open("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var inp, tmp, mod big.Int
	fib := make([]big.Int, 1, 50)
	fib[0] = *(<-ch)
	fmt.Fscanln(f, &inp)
	for inp.Sign() >= 0 {
		minMult := big.NewInt(0).Set(&inp)
		for i := 0; fib[i].Cmp(&inp) <= 0; i++ {
			if fib[i].Sign() != 0 {
				tmp.QuoRem(&inp, &fib[i], &mod)
				if mod.Sign() == 0 && tmp.Cmp(minMult) == -1 {
					minMult.Set(&tmp)
				}
			}
			if i == len(fib)-1 {
				fib = append(fib, *(<-ch))
			}
		}
		if minMult.Sign() == 0 {
			minMult.SetInt64(1)
		}
		fmt.Println("\nFor num: ", &inp, " i =", minMult)
		fmt.Print("Series: ")
		var prnt big.Int
		for i := 0; i < len(fib) && prnt.Mul(minMult, &fib[i]).Cmp(&inp) <= 0; i++ {
			fmt.Print(&prnt, " ")
		}
		fmt.Println()
		fmt.Fscanln(f, &inp)
	}
}

func fibGen(ch chan<- *big.Int) {
	a, b := big.NewInt(0), big.NewInt(1)

	for {
		ch <- big.NewInt(0).Set(a)
		a.Add(a, b)
		a, b = b, a
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	var i int
	flag.IntVar(&i, "i", -1, "Starting value greater than 1.")
	flag.Parse()

	if i == -1 {
		fmt.Print("Enter start value: ")
		_, err := fmt.Scanf("%d", &i)

		if err != nil {
			log.Fatal(err)
		}
	}

	threes(i)

}

func threes(in int) {
	for in > 1 {
		switch r := in % 3; r {
		case 0:
			fmt.Println(in, 0)
			in /= 3
		case 1:
			fmt.Println(in, -1)
			in = (in - 1) / 3
		case 2:
			fmt.Println(in, 1)
			in = (in + 1) / 3
		}
	}
	fmt.Println(in)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
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
	threes(uint64(i), 0, nil)
}

type node struct {
	parent *node
	val    int
}

var tree *node
var wg *sync.WaitGroup = new(sync.WaitGroup)

func threes(in uint64, sum int, root *node) {
	if root == nil {
		root = new(node)
		root.parent = nil
		tree = root
	}

	if in == 1 && sum == 0 {
		s, i := "1", 1
		for r := root; r != tree; r = r.parent {
			s = fmt.Sprintf("%d %d | %s", (i-1)*3+(3-r.val), r.val, s)
			i = (i-1)*3 + (3 - r.val)
		}
		fmt.Printf("Done\n%s\n", s)
		return
	}
	if in <= 1 {
		return
	}

	switch in % 3 {
	case 0:
		child := node{root, 0}
		threes(in/3, sum, &child)
	case 1:
		wg.Add(1)
		go func(in uint64, sum int, root *node) {
			child := node{root, -1}
			threes((in-1)/3, sum-1, &child)
			wg.Done()
		}(in, sum, root)
		child := node{root, +2}
		threes((in+2)/3, sum+2, &child)
	case 2:
		wg.Add(1)
		go func(in uint64, sum int, root *node) {
			child := node{root, 1}
			threes((in+1)/3, sum+1, &child)
			wg.Done()
		}(in, sum, root)
		child := node{root, -2}
		threes((in-2)/3, sum-2, &child)
	}
}

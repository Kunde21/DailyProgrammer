package main

import "bytes"

func extIfNeed(b [][]byte) [][]byte {

	ext := func(b [][]byte) [][]byte {
		for i, v := range b {
			tmp := make([]byte, 0, len(v)+2)
			b[i] = append(append(append(tmp, byte(' ')), v...), byte(' '))
		}
		return b
	}
	//Test need to extend Horizontally
	ln := len(b[0])
	for i := 1; i < len(b)-1; i++ {
		if b[i][0] != ' ' && b[i-1][0] != ' ' && b[i+1][0] != ' ' {
			b, ln = ext(b), ln+2
			break
		}
		if b[i][ln-1] != ' ' && b[i-1][ln-1] != ' ' && b[i+1][ln-1] != ' ' {
			b, ln = ext(b), ln+2
			break
		}
	}

	//Test need to extend Vertically
	for i := 1; i < len(b[0])-1; i++ {
		if b[0][i] != ' ' && b[0][i-1] != ' ' && b[0][i+1] != ' ' {
			tmp := make([][]byte, 0, len(b)+1)
			tmp = append(tmp, bytes.Repeat([]byte{' '}, ln))
			b = append(tmp, b...)
		}
		if b[len(b)-1][i] != ' ' && b[len(b)-1][i-1] != ' ' && b[len(b)-1][i+1] != ' ' {
			b = append(b, bytes.Repeat([]byte{' '}, ln))
		}
	}

	return b
}

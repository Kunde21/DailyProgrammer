package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in io.ReadCloser
	var out io.Writer
	fname := "file.txt"
	fmt.Println(byte('3'))

	if len(os.Args) > 1 {
		fin, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalln(os.Args[1], err)
		}
		in = fin
		fname = os.Args[1]
	} else {
		in = os.Stdin
	}
	if len(os.Args) > 2 {
		fout, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
		if err != nil {
			log.Fatalln("Out:", os.Args[2], err)
		}
		out = fout
		defer fout.Close()
	} else {
		out = bytes.NewBuffer(make([]byte, 0, 60))
		defer fmt.Println(out)
	}
	err := encode(fname, in, out)
	if err != nil {
		log.Println(err)
	}
	in.Close()
}

func encode(fname string, input io.Reader, output io.Writer) error {
	w := bufio.NewWriter(output)
	buf := make([]byte, 45)
	out := make([]byte, 61)
	n := 0
	var err error

	w.WriteString("begin 644 " + fname + "\n")
	for n, err = input.Read(buf); err == nil; n, err = input.Read(buf) {
		out[0] = byte(n) + 32
		if n%3 != 0 {
			buf = append(buf[:n], make([]byte, 3-n%3)...)
			n += 3 - n%3
		}
		var inIdx, outIdx int
		for inIdx, outIdx = 0, 1; inIdx < n; inIdx, outIdx = inIdx+3, outIdx+4 {
			out[outIdx] = (buf[inIdx] >> 2) + 32
			out[outIdx+1] = ((buf[inIdx]&3)<<4 | (buf[inIdx+1] >> 4)) + 32
			out[outIdx+2] = ((buf[inIdx+1]&0xF)<<2 | (buf[inIdx+2] >> 6)) + 32
			out[outIdx+3] = (buf[inIdx+2] & 0x3F) + 32
		}

		w.Write(out[:n/3*4+1])
		w.WriteRune('\n')
	}
	if err != io.EOF {
		return err
	}
	w.Write([]byte("`\nend\n"))
	if err = w.Flush(); err != nil {
		return err
	}
	return nil
}

func makeOutputFile(input io.Reader) (*os.File, error) {
	scan := bufio.NewScanner(input)
	if !scan.Scan() {
		return nil, io.ErrUnexpectedEOF
	}

	fields := strings.Split(scan.Text(), " ")
	if len(fields) != 3 || fields[0] != "begin" {
		return nil, fmt.Errorf("Format Error")
	}
	perm, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}
	return os.OpenFile(fields[2], os.O_CREATE|os.O_WRONLY, os.FileMode(perm))
}

func decode(input io.Reader, output io.Writer) error {
	scan := bufio.NewScanner(input)
	if !scan.Scan() { // Throw out first line
		return io.ErrUnexpectedEOF
	}
	w := bufio.NewWriter(output)
	buf := make([]byte, 45)

	for n := 0; scan.Scan(); {
		in := scan.Bytes()
		if in[0] == '`' {
			break
		}
		n = int(in[0] - 32)
		var inIdx, outIdx int
		for inIdx, outIdx = 1, 0; outIdx < n; inIdx, outIdx = inIdx+4, outIdx+3 {
			buf[outIdx] = (in[inIdx]-32)<<2 | (in[inIdx+1]-32)>>4
			buf[outIdx+1] = ((in[inIdx+1]-32)<<4 | (in[inIdx+2]-32)>>2)
			buf[outIdx+2] = ((in[inIdx+2]-32)<<6 | (in[inIdx+3] - 32))
		}
		w.Write(buf[:n])
	}
	if scan.Err() != nil {
		return scan.Err()
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

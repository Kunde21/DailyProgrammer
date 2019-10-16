package main

import (
	"bytes"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	buf := []byte("I feel very strongly about you doing duty. Would you give me a little more documentation about your reading in French? I am glad you are happy — but I never believe much in happiness. I never believe in misery either. Those are things you see on the stage or the screen or the printed pages, they never really happen to you in life.")
	target := bytes.NewBuffer(make([]byte, 0, 25))
	final := bytes.NewBuffer(make([]byte, 0, 25))

	encode("file.txt", bytes.NewReader(buf), target)
	decode(target, final)
	if res := bytes.Compare(buf, final.Bytes()); res != 0 {
		t.Error(res)
		t.Log(len(buf), buf)
		t.Log(final.Len(), buf)
	}
}

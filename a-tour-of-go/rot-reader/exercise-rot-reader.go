package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	asciiLower := byte('a')
	asciiUpper := byte('A')

	bufLen, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}

	for pos := 0; pos < bufLen; pos++ {
		if b[pos] >= asciiLower && b[pos] <= asciiLower+26 {
			b[pos] = asciiLower + ((b[pos] + 13 - asciiLower) % 26)
		} else if b[pos] >= asciiUpper && b[pos] <= asciiUpper+26 {
			b[pos] = asciiUpper + ((b[pos] + 13 - asciiUpper) % 26)
		}
	}
	return bufLen, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

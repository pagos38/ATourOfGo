package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	for i, v := range b {
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			b[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			b[i] -= 13
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrqgur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func rot13(b byte) byte {
	var beg byte
	if b >= 'A' && b <= 'Z' {
		beg = 'A'
	} else if b >= 'a' && b <= 'z' {
		beg = 'a'
	} else {
		return b
	}
	return (((b - beg) + 13) % 26) + beg
}

func _rot13(b byte) byte {
	var result byte
	switch {
	case b >= 'A' && b <= 'Z':
		result = (((b - 'A') + 13) % 26) + 'A'
	case b >= 'a' && b <= 'z':
		result = (((b - 'a') + 13) % 26) + 'a'
	default:
		result = b
	}
	return result
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	size, err := r.r.Read(b)
	for i, v := range b {
		b[i] = _rot13(v)
	}
	return size, err
}

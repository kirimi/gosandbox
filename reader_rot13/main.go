package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err == io.EOF {
		return n, err
	}

	for i, c := range b {
		decoded := decode(c)
		b[i] = decoded
	}

	return n, err
}

func decode(c byte) byte {
	const origin = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const encoded = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

	i := strings.Index(encoded, string(c))
	if i == -1 {
		return c
	}

	return origin[i]
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

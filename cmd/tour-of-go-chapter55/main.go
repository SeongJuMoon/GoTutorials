package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func encryptRot13(sb byte) byte {
	s := rune(sb)

	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		sb += 13
	}

	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		sb -= 13
	}

	return sb
}

func (rot13 rot13Reader) Read(data []byte) (len int, err error) {
	len, err = rot13.r.Read(data)
	for i := 0; i < len; i++ {
		data[i] = encryptRot13(data[i])
	}

	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr")

	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}

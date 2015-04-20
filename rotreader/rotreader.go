package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotr *rot13Reader) Read(buf []byte) (int, error) {

	tmp := make([]byte, len(buf))

	n, err := rotr.r.Read(tmp)

	//s := string(tmp[:n])
	s := []rune(string(tmp))

	var outs []rune

	for _, c := range s {
		switch diff := 'n' - 'a'; {

		case (c >= 'a' && c <= 'm') || (c >= 'A' && c <= 'M'):
			outs = append(outs, c+diff)
		case (c > 'm' && c <= 'z') || (c > 'M' && c <= 'Z'):
			outs = append(outs, c-diff)
		default:
			outs = append(outs, c)

		}
	}

	copy(buf, []byte(string(outs)))

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

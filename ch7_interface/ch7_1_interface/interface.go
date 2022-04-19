package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}
	if scanner.Err() != nil {
		panic(scanner.Err)
	}

	return int(*c), nil

}

// Ex 7.2
type CountWriter struct {
	wr   io.Writer
	nptr *int64
}

func (w CountWriter) Write(p []byte) (int, error) {
	var err error
	var n int
	n, err = w.wr.Write(p)
	*w.nptr = int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var p *int64
	nw := CountWriter{wr: w, nptr: p}
	return nw, p
}

// Ex 7.4
// type HtmlDecoder struct {
// 	rd io.Reader
// }

// func (r HtmlDecoder) Read(d []byte) (int, error) {
// 	n, err := r.rd.Read(d)
// 	if err != nil {
// 		return -1, err
// 	}
// 	return n, err
// }

func NewHtmlDecoder(s string) io.Reader {
	br := bufio.NewReader(strings.NewReader(s))
	return br
}

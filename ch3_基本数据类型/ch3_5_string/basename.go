package main

import (
	"fmt"
	"os"
	"strings"
)

// basename 3.5.4
func main() {
	var s = ""
	if len(os.Args) == 2 {
		s = os.Args[1]
	} else {
		s = "hello/world/go.go"
	}
	fmt.Println(basename(s))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s); i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}

func basename2(s string) string {
	pos := strings.LastIndex(s, "/")
	s = s[pos+1:]
	if pos = strings.Index(s, "."); pos >= 0 {
		s = s[:pos]
	}
	return s
}

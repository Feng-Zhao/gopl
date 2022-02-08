package main

import (
	"fmt"
	s "strings"
)

// 给方法起别名
var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	// 最后一个参数为替换次数，当小于 0 时，全部替换
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foooooo", "o", "0", 3))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

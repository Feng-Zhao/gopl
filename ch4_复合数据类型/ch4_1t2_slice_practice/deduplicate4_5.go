package main

import "fmt"

func main() {
	s := []string{"123", "123", "456", "456"}
	fmt.Println(s)
	s = deduplicate(s)
	fmt.Println(s)
}
func deduplicate(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	count := 0
	str := ss[0]
	for i := 1; i < len(ss)-count; {
		if str == ss[i] {
			temp := ss[i]
			copy(ss[i:], ss[i+1:])
			ss[len(ss)-1] = temp
			count++
		} else {
			str = ss[i]
			i++
		}
	}
	return ss[:len(ss)-count]
}

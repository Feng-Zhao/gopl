package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
// 这里每一个 uint64 有 64 bit,即,每一个 unit 以位图的形式,记录 64 个元素是否存在于集合中
// 每次更新集合时,先以 x/64 作为索引,找到 x 对应的在 []中的位置,
// 再以 x%64 计算 x 在 unit 中所代表的 bit 点
// 将该点置为1(使用`|`运算),表示 x 加入集合
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// 每次扩容 64 bit, 直到集合的的范围能够包含 x
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
// 作并集
func (s *IntSet) Unionwith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
// 这个 String 方法是为了友好的显示而做,
// fmt 会默认调用 String() 方法,因此可以实现 String()接口来做自定义 struct 的格式化输出
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

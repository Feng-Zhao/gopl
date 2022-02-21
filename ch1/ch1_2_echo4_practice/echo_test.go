package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkCh1_3_a(b *testing.B) {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	str := ""
	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			str += v
		}
		str = ""
	}
}

func BenchmarkCh1_3_b(b *testing.B) {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	str := ""
	for i := 0; i < b.N; i++ {
		str = strings.Join(slice, ",")
	}
	fmt.Println(str)
}

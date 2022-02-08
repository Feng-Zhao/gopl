package main

import (
	"fmt"
	"testing"
)

// 通常 test代码应该和源码在同一个文件夹下
// 例如 intutils.go, 测试代码一般命名为 intutils_test.go.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 测试代码的 func 名字以 Test 开头，go test 工具会自动识别这类 func
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* 报告错误，然后继续执行
		// t.Fatal* 报告错误，然后直接停止测试
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	// 将测试用例和 预期结果使用数组存储
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// 对于每次测试 使用 t.Run() 执行子测试
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// go test -v 运行测试

// benchmark 方法以 Benchmark 开头
func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

// go test -bench=. 运行 benchmark

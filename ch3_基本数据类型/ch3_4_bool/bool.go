package ch3_4_bool

// bool 类型无法隐式转换为 0/1
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
func itob(i int) bool {
	return i != 0
}

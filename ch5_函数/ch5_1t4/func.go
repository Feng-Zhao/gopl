package functions

type Receiver struct {
	this string
}

// 函数声明
// func 关键字 [(接收实体,即,这个函数属于那种 class)] funcName (参数列表) (返回值列表)
func (r Receiver) funcName(argA string, argB string, args ...string) (resultA string, resultB string) {
	return "", ""
}

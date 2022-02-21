package ch2_5_type

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	// 两种 type 并不相同,并不能直接赋值
	// 需要强制转换
	return Fahrenheit(c*9/5 + 32)
}
func FtoC(f Fahrenheit) Celsius {
	// 两种 type 并不相同,并不能直接赋值
	// 需要强制转换
	return Celsius((f - 32) * 5 / 9)
}

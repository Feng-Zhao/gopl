# ch3_基本数据类型

## 整型 int
- 整型大体上可以分为有符号,无符号两类,有符号整型使用补码表示.无符号整型使用原码表示
- go 中的类型无论底层是否相同,几乎都需要显示的类型转换.

### 二元操作符有优先级
```ini
* / % << >> & &^
+ - | ^
== != < <= > >=
&&
||
```
二元运算要求两个操作数的类型相同.否则编译报错

- 浮点型向整型转换时,会损失精度,并且转换时是趋向 0 取整,并截去小数部分

## 文字符号 rune
使用 '' 括起来表示 rune
使用 %c %q 打印

```go
package main

import "fmt"

func main() {
	// 使用 '' 括起来表示 rune
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	// 使用 %c %q 打印
	fmt.Printf("%d %[1]c %[1]q\n",ascii)
	fmt.Printf("%d %[1]c %[1]q\n",unicode)
	fmt.Printf("%d %[1]q\n",newline)
}
```

## 整型 float
- 十进制下 float32 约为 6 位数, float64 约为 15 位数
- 浮点值打印时的谓词 %g(自动规划精度) %f(无指数形式) %e(有指数形式)

## 复数
```go
package main
import "fmt"
// complex 两种类型 complex64/complex128
// 内部是使用 float32/float64 实现
var x complex128 = complex(1, 2) // 复数 1+2i
//var y complex64 = complex(3, 4)
var y complex128 = complex(3, 4)

func main() {
	fmt.Println(x * y)
	fmt.Println(real(x * y)) // 实部
	fmt.Println(imag(x * y)) // 虚部
}
```

## 布尔型 bool
- 只有两个值 true/false
- 不能隐式转换为 int 的 1/0
- && || 运算符有短路

## 字符串
- 字符串不可变
- 使用 `...` 表示字面量字符串,不理会转义字符\

## Unicode & UTF-8
- unicode 每个字符 32 位,即,4 字节
- utf-8对 Unicode 做了简化,每个字符 1-4 字节不等,对 acsii 码完全相同
- go 以 utf-8 编码


utf-8 编码

| 编码                                  | 长度                 |
|-------------------------------------|--------------------|
| 0xxxxxxx                            | 0-127 完全兼容 ASCII 码 |
| 110xxxxx 10xxxxxx                   | 两字节编码              |
| 1110xxxx 10xxxxxx 10xxxxxx          | 三字节编码              |
| 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx | 四字节编码              |

## 常量
- 使用 const 声明
- 常量生成器 iota 从 0 开始,每次调用加 1,
- 常量可以没有类型,称为无类型常量,这种常量的精度更高,可以看做是至少 256 位精度
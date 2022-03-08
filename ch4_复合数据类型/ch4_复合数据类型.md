# 复合数据类型

## 数组 slice
- `var a [3]int` 数组的长度是数组类型的一部分
- **函数传参的时候,参数会复制一个副本,副本供函数内部使用**,所以把大数组通过参数传递会极大降低程序效率,**并且函数内部改变数组,并不会影响外部原来的数组**
- 数组可以用 == 做比较,只有当数组长度,类型,和数组中的每个元素都一样时,两个数组才相等

## slice
- `var a []int` slice 的长度可变,可以理解为数组的可变长度视图
- slice中包含底层数组的指针,因此如果使用 slice 作为参数,那个函数对 slice 的修改会反映在外部 slice 中
- slice 不能用 == 比较,若想比较 slice 则需要自己实现对应的函数
- slice 的零值为 nil
- 检查 slice 是否为空时,应使用 len(slice) == 0
- slice 使用 append(slice,element) 向自己添加元素,并且会自动扩容(每次扩容 2 倍,每次扩容会调用 copy 函数),
由于这个过程有可能会创建新的底层数组,go 不保证 slice 的底层数组不变,所以需要将 append() 函数返回的 slice 重新赋值给 slice 变量
即,s = append(s,element)

## map
- map 的键必须是可以通过 == 比较的数据类型
- map 中无对应值的键,会返回零值/空值,并且 ok 变量会返回 false
- map中的元素并不是变量,所以无法通过&获取元素地址(注:其中一个原因为 map 可能会在resize 过程中移动元素,并不能保证元素的地址不变)
- map 的零值是 nil,向零值 map 中设置元素会引起 panic
- map 不可通过 == 比较, 若想比较 map 则需要自己实现对应的函数
- range map 遍历的顺序是随机的

## 结构体 struct 
- 与 C 的结构体类似,是组织/组合数据的一种方式
- 结构体中成员变量的顺序是结构体定义的一部分,更改成员变量顺序会得到不同的结构体
- 结构体中不能包含自己,但是可以包含自己类型的指针.
- 结构体作为参数时,是复制一份然后值传递,所以从效率考虑,或者需要函数内部对结构体的更改能反映到外部变量上时,需要传递结构体指针
- 如果结构体内的所有成员变量都是可以比较的,则结构体可以比较
- 结构体可以通过\`匿名成员\`的方式(即,没有变量名,只声明一个结构体类型或者结构体类型的指针)呈现组合的效果.实际上这类成员变量并不是"匿名"的,
他们的名字就是类型的名字,不过在使用`.`访问时,可以省略.也因此结构体中不能有两个想同类型的匿名成员,这会引起成员变量的隐式变量名冲突
- 结构体的组合不仅会获得子类的变量,还会获得子类的方法

## Json
- 使用json.Marshal() json.Unmarshal() 来对 go结构体和 json形式的 []byte 进行转换
- 另一种方式是使用 json.Encoder() json.Decoder() 进行流式编解码
- 编解码只针对于可导出的字段
- 可通过成员标签 ``json:"name"`` 来指定序列化时对应字段的名字

## 文本模板
```go
package main

import (
	"fmt"
	"os"
	"text/template"
)
// 定义模板字符串
const temp = `{{.Name}} is a template, the value of a is {{.A}}, the value of b is {{.B}}
a+b is {{Add .A .B}}`

type Data struct {
	Name string
	A    int
	B    int
}

func main() {
	data := Data{Name: "data", A: 1, B: 2}
	// 创建模板
	result, err := template.New("report").Funcs(template.FuncMap{"Add": Add}).Parse(temp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 渲染模板
	result.Execute(os.Stdout, data)
}

func Add(a, b int) int {
	return a + b
}
```
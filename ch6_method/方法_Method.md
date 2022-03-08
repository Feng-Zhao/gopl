# 方法 Method

method 是一种函数(func),不过 method 有其接受者类型(receiver),将func 和具体的 type 绑定,从而实现面向对象的效果.
method 只能通过具体的 receiver 实例来调用

## 方法声明

```go
package functions

type Receiver struct {
 this string
}

// 函数声明
// func 关键字 funcName (参数列表,可以有变长参数) (返回值列表)
func (r Receiver) funcName(argA string, argB string, args ...string) (resultA string, resultB string) {
 return "", ""
}
```

- 相同命名空间下,在此章节即为 同一个 struct 下,方法和字段是同属同一命名空间的,所以同一 struct 中,方法和字段不能重名

## 指针作为 receiver

- 当我们需要在方法内部对 receiver 进行更改时,或者 receiver 很大,想要避免调用过程中的复制时,可以使用指针作为 receiver
- 习惯上,如果对象的一个方法使用指针作为接受者,那么这个对象的所有方法都应该使用指针作为接收者
- 另外,不允许本身就是指针的类型被用作为 receiver (会提示编译错误)
- nil 可以作为 receiver,这种情况一般要在注释中标注清楚,nil 作为 receiver 时的行为

## 对象的组合

- go 中对象的组合通过在结构体中内嵌完成,主体可以使用被嵌入的结构体的所有方法

## 方法变量

包含具体接收者的方法可以整体赋值给一个变量,这个方法隐含的把接收者实例绑定到方法变量上

## 方法表达式

可以不指定接收者实例,而把 类型.方法 赋值给变量.这个变量相当于一个函数,其第一个形参为接收者类型

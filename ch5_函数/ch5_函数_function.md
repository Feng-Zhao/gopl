# 函数

[toc]

## 函数声明

```go
package functions

// 函数声明
// func 关键字 funcName (参数列表,可以有变长参数) (返回值列表)
func funcName(argA string, argB string, args ...string) (resultA string, resultB string) {
 return "", ""
}
```

- 函数传递时都是值传递,如果需要函数对参数的修改影响调用者,则需传递指针类型的参数(指针,slice,map,chan,func)
- 只有声明没有函数体的函数,是使用了 go 之外的语言实现的.

## 函数处理错误的几种形式

- 直接向上返回:

  ```go
  if err != nil {
    return err
  }  
  ```

- fmt.Errorf() 包装错误信息

  ```go
  if err != nil {
    return fmt.Errorf("err msg:%v",err)
  }
  ```

- 报告错误并退出程序 fmt.Fprintf(os.Stderr,"err msg: %v",err) + os.Exit(code)

  ```go
  if err != nil {
    fmt.Fprintf(os.Stderr,"err msg: %v",err)
    os.Exit(code)
    // 等价于
    log.Fatalf(os.Stderr,"err msg: %v",err) 
  }
  ```

- 记录到 log 或者 os.Stderr
- 直接忽略

### 特殊错误处理, 这种 err 是 error 类型,但是并不表示程序出错,而是表示程序遇到某种特殊状态

  典型的场景就是 IO 读取文件时,读到文件最后会有 io.EOF 错误,这种错误需要特殊判断,然后挑出读取文件的循环

  ```go
    if err != nil{
        if err == io.EOF{ // 文件读取结束
            break
        } else {
            return err // 遇到其他错误
        }  
    }
  ```

## 函数变量

- 函数是 first-class value, 它可以赋值给变量,可以作为参数传递,也可以作为函数的结果被返回
- 函数变量的零值为 nil, 调用空函数会导致 panic

  `panic: runtime error: invalid memory address or nil pointer dereference`
- 函数变量可以和 nil 比较,但是函数变量不可以相互比较,所以也不可以作为 map 的键
- 使用函数变量可以将逻辑和实际行为分离,非常方便的实现策略模式,模板方法等,有助于更加灵活的系统设计

## 匿名函数

- 作用域相关: 命名函数只能在包级别中声明,而匿名函数可以使用函数字面量来定义,没有作用域级别限制,并且匿名函数可以获取到整个词法环境,即,能够使用外层作用域中的变量
- 闭包相关: 匿名函数作为函数返回值,相当于返回一个闭包,这个闭包维护自己的变量池,即,函数变量可以拥有状态
- 在使用匿名函数时需要注意不要把迭代变量,即,从 for 语句中初始化的变量赋值给函数(不光是匿名函数),因为这些迭代变量是被 for 循环的每次迭代内部共享的,
是一个地址(指针),而不是一个值.尤其在 defer 和 go 语句中,不恰当的使用迭代变量会引起错误.正确使用引用迭代变量的方法是在 for 循环内部声明一个内部变量,
并用迭代变量初始化其值,这样这个变量的会在每次迭代中都被声明和初始化,保持当次迭代时迭代变量的值.

## 延迟函数 defer

- 函数调用前使用 defer 修饰,则该函数会在 return 结束后,函数结束前被调用,简化了复杂逻辑路径下的后处理操作,通常用来关闭系统资源
- 同一函数下 defer 函数的调用顺序是后进先出的,即,栈式
- defer + 匿名函数 可以更改 return 的返回值

## 宕机 panic

- 当发生 panic 时,defer 语句执行,函数退出,打印栈调用信息,异常退出函数
- 宕机只应该用于处理"不可能到达的逻辑",可以预见的错误状态应该用 error 来标识

## 宕机恢复 recover

- recover 在 defer 函数中,可以在宕机时反馈宕机信息,做处理后恢复系统然后让系统继续运行
- 未发生 panic 的情况下调用 recover() 则没有任何效果,只返回 nil

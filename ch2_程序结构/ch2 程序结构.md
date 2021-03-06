# ch2 程序结构

## 变量及其声明

- 包级别的变量在 main 函数之前初始化, 局部变量在函数执行期间初始化
- 短变量声明的方式只能在 func 内部使用,包级别的变量都需要使用 var 关键字声明
- 变量逃逸(生命周期相关):当局部变量在代码块结束后仍可访问(一般就是将此变量的地址付给了一个外部指针),则称这个变量从函数/代码块中逃逸.
此时编译器在分配这个变量时,会将变量分配在堆中.
而未逃逸的变量可以分配在栈上,随着函数/代码块的退栈操作而自动被销毁,从而达到更高的性能和效率

## 指针

- 指针是变量的地址,go中指针不能参与运算
- 取地址符 &
- 指针符 *
- 指针的零值为 nil

## new

- new(Type) 返回的是指针
- 比起正常声明和取地址符,只是不需要给变量一个名字,相当于返回一个 type 类型的匿名变量的指针

## 赋值

- go 支持多重赋值
- go 支持空标识符 _ 来接收多个返回值中不需要的部分

## 类型声明

- 使用 type 定义的新类型,即使底层类型相同,两种类型也是不同的类型,即,两者无法相互赋值

## 包

- go 使用包来组织代码,同属于同一个包(package)下的文件可以看做是同一个文件,同一文件夹下的文件只允许使用有且仅有一个包名
- 包中的变量/函数是否可导出取决于变量/函数的首字母是否为大写

## 作用域

- 变量的作用于取决于其被声明的语法块
- 作用域不等价于变量的生命周期,作用域表示可见性,是一个编译时属性.生命周期是变量在程序执行期间可否被引用的起止时间,是一个运行时属性

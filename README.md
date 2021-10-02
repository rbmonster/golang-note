# golang-note



声明：
- 变量 var
- 常量 const
- 类型 type
- 函数 func


## 程序结构

### 变量声明
**变量通用形式**\
`var name type = expression`\
类型type 及 表达式expression 部分可以省略一个

```go
var a = ""
fmt.Println(a)

var b string
fmt.Println(b)

var i,j, k int
fmt.Println( i,j,k)

var c, f, s = true, 2.3, "four"
fmt.Println(c,f,s)
```

**短变量声明**：作为变量声明的一个可选形式用来声明和初始化局部变量\
`name := expression` 

```go

i := 100
j, k := 10, 20

// 交换j与k值
j, k = k, j

// 短变量声明复用，err变量实现复用，声明中必须要有一个变量不重复
resp, err := http.Get(url)
resp1, err := http.Get(url)
```

**指针**：变量存储值的地方，指针的值为一个变量的地址。

```go
x := 1
p := &x
fmt.Println(x, *p, p)  // 1 1 0xc00001e070
*p = 3
fmt.Println(x, *p, p) //  3 3 0xc00001e070
```

**new函数创建变量**: 初始化为T类型的零值，并返回其地址

```go
p := new(int)
q := new(int)
fmt.Println(p == q)  // false
fmt.Println(*p == *q) // true
```

**多重赋值**

```go
v, ok = m[key]

_, err = io.Copy(dst, src) // _丢弃返回值
```

### 变量的生命周期
- 包级别变量的生命周期：整个程序的执行时间
- 局部变量：动态的生命周期，变量一直生存到它变得不可访问，这时它占用的空间被回收
> 垃圾回收思路：每一个包级别的变量，以及每一个当前执行函数的局部变量，可以作为追溯该变量的路径源头，通过指针和其他方式的引用都无法找到变量，那么说明变量不可访问，因此可以进行回收。

```go
// g函数返回，变量y回收
func g() {
	y := new(int)
	*y = 1
}


// f函数执行完成还可以用global变量访问，x从f函数逃逸
var global *int
func f() {
	var x int
	x = 1
	global = &x
}
```

### 类型声明
type声明定义一个新的明明类型，它和某个已有类型使用相同的底层类型

```go

package tempconv

import "fmt"

 // 类型声明
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
```

### 包和文件

#### 包文件导入

实体的第一个字母大小写决定其可见性是否跨包

```go
package basic

func Fib(n int) int {
	var x,y = 0, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
```

```go
package main

import (
	"fmt"
	"golang-note/src/basic"
)

func main() {
	var res = basic.Fib(100)
	fmt.Println(res)
}

```

#### 包初始化
1. 包级别变量: 声明顺序初始化
2. 包存在多个.go文件，编译器排序吼，根据文件顺序初始化
3. 执行`func init() { /* ...*/}`函数

```go
package main

import (
	"fmt"
	"golang-note/src/basic"
)

func init() {
	fmt.Println("initing")
}

func main() {
	var res = basic.Fib(10)
	fmt.Println(res)
}

```

### 作用域
声明的作用域：指用到变量时所声明名字的源代码段

函数中，词法块嵌套，一个局部变量声明可能覆盖另一个。
```go
package main

import "fmt"

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x :=x[i]
		if x != '!' {
			x := x+ 'A'- 'a'
			fmt.Println(x)
		}
	}
	fmt.Println(x)
}
/**
output =>
72
69
76
76
79
 */
```


**隐式词法块的作用域**
else中仍可以使用if中的声明
```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if x := f(); x == 0 {
		fmt.Println("x match zero")
	} else if y := f(); y == x {
		fmt.Println("range num tow match")
	} else {
		fmt.Println("nomatch num:", x, y)
	}
}

func f() int  {
	return rand.Intn(10)
}


/**
output:
nomatch num: 1 7
*/
```

## 基本类型
- 整型：int8, int16, int32, int64
- 浮点数：float32, float64
- 复数：complex64, complex128
- 布尔值：true, false
- 字符串

> 取模的余数正负号总与被除数一致 `-5%3=-2  | -5%-3=-2 | 5%-3=2`\
> 布尔值判断短路行为：运算符左边的判断已经能确定结果，则右边不计算\
> 布尔值无法转换0，1

### 字符串
![image](https://github.com/rbmonster/file-storage/blob/main/golang-note/basic/stringconstruct.png)

```go
var s = "hello,world!"

s[0] = 'L' // 编译错误，字符串内部不可赋值
```
字符串本身所包含的字节序列永不可变。

**字符串字面量**
```
\a 警告
\b 退格符
\f 换页符
\n 换行符
\r 回车符
\t 制表符
\v 垂直制表符
\' 单引号
\" 双引号
\\ 反斜杠
```

工具包：
- strings: 搜索、替换、比较、修整、切分与连接字符串
- bytes: 操作字节slice。`bytes.Buffer` 频繁操作字符使用
- strconv：转换布尔值、整数、浮点数等转换相关
- unicode：判别字符符号特性的函数，如IsDigit、IsLetter、IsUpper和IsLower
- FormatInt 和formatUint ：进制位格式化


### 常量
const常量：保证在编译阶段就计算出表达式的值，并不需要等到运行时。本质上都属于基本类型如布尔型、字符串或数字
```go
const (
a = 1
b = 2
)

//同时声明，其他项目省略则复用前面的表达式
const (
	a = 1
	b        // value 1
	c = 2
	d       // value 2
)
```

**常量生成器 iota**: 从0开始取值，逐项加1
```go
package main

import (
	"fmt"
)

type Weekend int
const (
	Sunday Weekend = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Monday)
}
```

## 复合数据类型

### 数组
数组声明：
```go

var q [3]int = [3]int{1, 2, 3}
//var q = [3]int{1, 2, 3} 
for i,v := range q {
    fmt.Printf("%d %d\n", i, v)
}
p := [...]int{1,2,3}
for _, v := range p {
    fmt.Printf("%d \n", v)
}

fmt.Println(q == p)  // 数组支持直接 == 判断equal

// 数组的调用函数参数传递，必须指定数组的大小
func zero(ptr [3]int) {
    for i := range ptr {
        ptr[i] = 0
        fmt.Printf("%d %d \n", i, ptr[i])
    }
}
```

### 值传递与引用传递
Go把数组和其他类型都看成值传递。在函数调用时候传入参数都会创建一个副本，然后赋值给对应的函数变量，所以传递大的数组会变得很低效
> 对于数组类型的函数参数，可以显示地传递一个数组指针，防止数组参数传递的低效复制

```go
func main() {
	zero(&p)
	fmt.Println(p)
}

// 数组的个数必须匹配
func zero(ptr *[3]int) {
	for i := range ptr {
		ptr[i] = 0
		fmt.Printf("%d %d \n", i, ptr[i])
	}
}
```

### slice
slice 标识一个拥有相同类型元素的可变长度的序列。
> ⚠️ slice和数组的区别在于：
> 1. 数组是需要指定个数的，而切片则不需要。
> 2. slice 无法直接使用==比较，而数组可以直接使用

**slice创建**
1. slice 操作符 `s[i:j]` 可以用于创建一个新的slice
2. make 函数可以创建一个无名数组并返回它的一个slice。表达式`make([]T, len, cap)`

```go
s := []int{1,2,3,4,5}   // 切面的创建 未指定大小
s2 := s[:3]             // 使用操作符创建
var s3 = make([]int, 4, 10)  // 使用make函数创建 
array1 := [...]int{1,2,3,4,5} // 数组的创建，省略了大小
array2 := [5]int{1,2,3,4,5} // 数组的创建，指定了大小
```


slice有三个属性：指针、长度和容量
- 指针：数组第一个元素
- 长度：slice的长度
- 容量：从slice的起始元素到底层数组的最后一个元素间的个数


```go
package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5}
	s2 := s[:3]
	fmt.Println(s2)
	extendLen := s2[:4]     // 在s2 slice的范围内扩展了slice到4位，最终比s2长
	fmt.Println(extendLen)
	extendCap := s2[:9]     // 超过capacity容量，运行报错
	fmt.Println(extendCap)
}
// output
/**
[1 2 3]
[1 2 3 4]
panic: runtime error: slice bounds out of range [:9] with capacity 5

goroutine 1 [running]:
main.main()
        /Users/sanwuhong/private/golang-note/src/test/test.go:22 +0xf6

*/
```

**地址传递**：slice包含了指向数组元素的指针，所以将一个slice传递给函数，可以在函数内部修改底层数组的元素
```go
package main

import "fmt"
func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

```


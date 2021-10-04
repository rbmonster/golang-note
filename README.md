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

```
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

```

i := 100
j, k := 10, 20

// 交换j与k值
j, k = k, j

// 短变量声明复用，err变量实现复用，声明中必须要有一个变量不重复
resp, err := http.Get(url)
resp1, err := http.Get(url)
```

**指针**：变量存储值的地方，指针的值为一个变量的地址。

```
x := 1
p := &x
fmt.Println(x, *p, p)  // 1 1 0xc00001e070
*p = 3
fmt.Println(x, *p, p) //  3 3 0xc00001e070
```

**new函数创建变量**: 初始化为T类型的零值，并返回其地址

```
p := new(int)
q := new(int)
fmt.Println(p == q)  // false
fmt.Println(*p == *q) // true
```

**多重赋值**

```
v, ok = m[key]

_, err = io.Copy(dst, src) // _丢弃返回值
```

### 变量的生命周期
- 包级别变量的生命周期：整个程序的执行时间
- 局部变量：动态的生命周期，变量一直生存到它变得不可访问，这时它占用的空间被回收
> 垃圾回收思路：每一个包级别的变量，以及每一个当前执行函数的局部变量，可以作为追溯该变量的路径源头，通过指针和其他方式的引用都无法找到变量，那么说明变量不可访问，因此可以进行回收。

```
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

```
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
```
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
```

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

```
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

```
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

**append函数**：用来将元素追加到slice后面
```
s := []int{0, 1, 2, 3, 4, 5}
var s2 = append(s, 1,2,3)   // [0 1 2 3 4 5 1 2 3]
var s3 = append(s, s...)    // [0 1 2 3 4 5 0 1 2 3 4 5]
```



### slice 模拟栈操作
```go
package main

import "fmt"

func main() {
	var stack = []int {}
	var nums =  [6]int {1,2,3,4,5,6}
	for _,v := range nums{
		stack = append(stack, v)	// 栈push操作
	}
	fmt.Println(stack)
	for len(stack) != 0 {
		var top = stack[len(stack)-1]	// 栈顶元素
		fmt.Println(top)    
		stack = stack[:len(stack)-1]  // 栈pop操作
	}
}
```

### map
**map创建**：
- 内置map函数创建map
- map字面量创建
```
arg := make(map[string]int)
arg["alice"] = 32
fmt.Println(arg)

args1 := map[string]int {
    "alice" : 32 ,
}
fmt.Println(args1)

// 创建map<key, map<key, bool>> 
var graph = make(map[string]map[string]bool)
```


**map操作**
```

delete(args1, "alice")	 	// 删除map元素

args1["bob"] = 12		// map 对于空值，直接返回默认值0
fmt.Println(args1)

// 判断元素是否存在的两种方式
if age, ok := args1["bob1"]; !ok { fmt.Println("123", ok, age) }

age, ok := args1["bob"]
if !ok {
    fmt.Println("error for not bob age,use default", age)
}
```


### map用图的连通性测试
```go
package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	addEdge("a", "d")
	addEdge("d", "a")
	fmt.Println(hasEdge("a", "b"))
}
```

### 结构体
**结构体**：将零个或多个任意类型变量组合的聚合数据类型
> 结构体的成员变量名称是首字母答谢，那么表示该变量可以导出


#### 结构体初始化
支持两种初始化，一种按顺序，一种指定变量名.
```
type Point struct {
	X, Y int
}

func main() {
	var p1 = Point{1,2}
	p2 := Point{X:3, Y:4}  
	p3 := Point{Y:4}    // {0, 4}
	fmt.Println(p1, p2, p3)
}
```

**嵌套结构体的两种初始化方式**
```go
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Redius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	// 第一种创建方式
	w := Wheel{Circle{Point{X: 8, Y: 9}, 10}, 12}
	// 第二种创建方式
	w2 := Wheel{
		Circle: Circle{
			Center: Point{ X: 1, Y:2},
			Redius: 10,
		},
		Spokes: 12,
	}
	fmt.Printf("%#v\n", w2)
	// 嵌套结构体的访问
	fmt.Println(w.Circle.Center.X)
}

```

#### 结构体应用树的排序
结构体的应用：
```go
package treesort

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func TreeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
```

### JSON 
Go的数据结构转换为JSON称为marshal，Marshal生成一个字节slice。
> 只有可导出的成员可以转换为JSON字段
> MarshalIndent将输出格式化的结果 ''
```
func main() {
	w := Wheel{
		Circle: Circle{
			Center: Point{ X: 1, Y:2},
			Redius: 10,
		},
		Spokes: 12,
	}
	data, err := json.Marshal(w)
	if err!=nil {
		log.Fatalf("json marshaling fail: %s", err)
	}
	fmt.Printf("%s\n", data)

	dataFormat, err := json.MarshalIndent(w, "","    ")
	if err!=nil {
		log.Fatalf("json marshaling fail: %s", err)
	}
	fmt.Printf("%s\n", dataFormat)
}
// output
{"Circle":{"Center":{"X":1,"Y":2},"Redius":10},"Spokes":12}
{
    "Circle": {
        "Center": {
            "X": 1,
            "Y": 2
        },
        "Redius": 10
    },
    "Spokes": 12
}
```

unmarshal为将JSON字符串解码为Go数据结构的操作
> unmarshal阶段JSON字段关联到Go结构成员的名称是忽略大小写的，结构体的成员必须首字母大写保证可访问进而转换成功
```

func main() {
	jsonStr := "{\"Circle\":{\"Center\":{\"X\":1,\"Y\":2},\"Redius\":10},\"Spokes\":12}"
	var w2 Wheel
	// 参数传递指针
	if err := json.Unmarshal([]byte(jsonStr), &w2) ;err != nil{
		log.Fatalf("unmarshal error:%s",err)
	}
	fmt.Println(w2)
}
```


### 其他
`json.Decoder` 流式解码器，用来依次从字节流里面解码出多个JSON实体

`import "html/template"`: 将字符映射成既定的html模版 
`import "text/template"`: 将字符映射成既定的text模版


## 函数
函数的形式：
```
func name(parameter-list) (result-list) {
    body
}

// 指定每个参数
func add(x int, y int) int {
	return x+y
}

// 省略型参声明，指定了返回值的参数
// 一个函数如果有命名的返回值，可以省略return语句的操作数
func sub(x, y int) (z int) {
	z = x-y
	return
}

func first(x int, _ int) int {
	return x
}
```

**多返回值**
```
func HourMinSec(t time.Time) (hour, minute, second int) {
	return t.Hour(), t.Minute(), t.Second()
}
```

### 错误
Go语言通过使用普通的值而非异常来报告错误，Go语言使用通常的控制流机制应对错误，这种方式在错误处理逻辑方面要求更加小心谨慎。


`fmt.Errorf`使用`fmt.Sprintf`格式化一条错误信息并且返回一个新的错误值
```
err1 := fmt.Errorf("http get url %s error:%v","www.baidu.com", err)
```

因为错误消息会频繁的拼接起来，字符串首字母不应该大写而且应该避免换行，这样可以有利于使用grep工具进行查询

log.Fatalf，默认将时间和日期作为前缀添加到错误消息前。支持自定义设置命令名称，通过设置log的属性
```
log.Fatalf("http get fail:%v", err)
// output
// 2021/10/03 19:52:53 http get fail:Get "http//:www.baidu.com": unsupported protocol scheme ""
```

**文件结束标识的特殊判断**：io包保证任何由文件结束引起的读取错误均抛出`io.EOF`，因此对于文件读取的时候该错误需要特殊判断

### 函数变量
函数变量使得函数不仅将数据进行参数化，还将函数行为当作参数进行传递。

```go
package main

import (
	"fmt"
)

func main() {
	x := []int { 1,2,3,4,5,6}
	function(x, func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	})
	fmt.Println(x)
}

func function(x []int, reverse func(s [] int)) {
	fmt.Println(x)
	reverse(x)
}
```

**匿名函数**：func关键字后面没有函数的名称，这种方式定义的函数可以获取到整个词法环境
> 函数变量类似于使用闭包方式实现的变量
```
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
```

闭包易忽略的问题：⚠️捕获迭代变量
```
var rmdirs []func() 
for _,dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)       // 错误，dir在循环中赋值一直改变
    })
}
for _,rmdir := range rmdirs {
    rmdir()
}
```

**变长函数**：使用省略号代表参数，正常用于格式化中
```
func sum(val ...int) int {
	total := 0
	for v := range val {
		total +=v
	}
	return total
}
```

### 延迟函数
延迟函数defer就是一个普通函数活着方法调用。无论是正常调用执行return或函数执行完毕，还是不正常的情况如宕机，**实际的调用推迟到⚠️包含defer语句函数执行完毕**之后才执行。

```go
package main

import (
	"log"
	"time"
)

func bigSLowOperation() {
	// 延迟调用trace函数，等待该方法结束之后，再调用trace返回的函数参数执行，实现记录方法执行时间的效果
	defer trace("bigSLowOperation") ()
	time.Sleep(3*time.Second)
}

// defer调用的时候，前两行语句已经执行
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start))}
}
```

互斥锁与延迟函数的应用
```
var mu sync.Mutex
var m = make(map[string] int)

func lookup(key string) int {
    mu.Lock()
    // do some logic handle common var m
    defer mu.Unlock()
    return m[key]
}

// 延迟函数执行
func double(x int) (result int) {
    defer func() { fmt.Printf("double(%d) = %d\n", x, result)} ()
    return x + x
}
```

defer语句经常适用于成对的操作，比如打开和关闭，连接和断开，加锁和解锁\
不适合延迟函数的场景： 如本地文件关闭，因为文件系统的修改，经常在关闭文件的时候才做文件检查。


## 方法
Go语言中的面向对象编程：对象就是一个简单的值活着变量，并且拥有其方法，而**方法是某种特定类型的函数**。面向对象编程就是使用方法来描述每个数据结构的属性和操作。

类型拥有的所有方法名必须都是唯一的，但不同的类型可以使用相同的方法名。无java中重载的特性。

方法声明：
```
// 相比函数在方法名之前增加了参数，表示该方法的接受者
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

func main() {
    p := Point{1,2}
    q := Point{3,5}
    // 方法的调用
    fmt.Println(q.Distance(p))
}
```

`q.Distance`称为选择子，编译器会通过方法名与接收者的类型决定调用哪一个函数

### 指针方法
指针方法调用，可以直接使用类型的方法调用，编译器会进行变量的&p的隐式转换
> `nil`也可以作为一个方法的接收者，比如map和slice类型中的应用。
```
func main() {
    p := Point{1,2}
    p.ScaleBy(12)
    (&p).ScaleBy(12)
    fmt.Println(p)
}

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

// 该方法支持nil作为方法接收者
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}
```

### 结构体内嵌类型的方法调用
嵌套的结构体类型组成复杂的类型，**内嵌的方法都被纳入到上层类型**中。调用的时候编译器会进行方法包装再调用
> 编译器处理方法调用顺序：直接声明方法 -> 内嵌字段方法 -> 内嵌字段的内嵌字段方法...   遇到同名方法直接报错

```
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Circle struct {
	Point
	Redius int
}

func main() {
	c1 := Circle{Point{1,2}, 3}
	c2 := Circle{Point{3,4}, 1}
	// 编译器会使用包装方法调用，Point的Distance方法
	fmt.Println(c1.Distance(c2.Point))
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}
```
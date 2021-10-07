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

### import导包
`import _ "xxxx""`**空导入**：导入的包在文件中无引用，近视利用其副作用，对包级别变量执行初始化表达式的求值，并执行init函数

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
    
    // 指定方法参数
	dis := p.Distance
	res := dis(q.Point)
	fmt.Println(res)
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

### 封装
Go语言只有一种方式控制命名的可见行：定义的时候首字母大写的标识符是可以从包中导出的，而首字母没有大写不导出。

封装三个优点
1. 防止使用者肆意修改对象内变量。
2. 隐藏实现细节防止使用方依赖的属性发生改变。
3. 使用者不能直接修改，因此不需要更多的语句来检查变量的值。


## 接口

### 接口定义
Go语言的接口独特之处在于接口是隐式实现的。对于一个具体的类型，无须声明它实现了哪些接口，只要提供接口所必须的方法即可。

接口是一种抽象类型，一个具体类型只要实现了接口的所有方法，就该接口参数就可以指向该类型。这种把一种类型替换，为满足同一接口的另一种类型的特性称为可取代性。
> 一个接口类型定义了一套方法，如果一个具体类型要实现该接口，那么必须实现接口类型定义中的所有方法。
> 空接口类型对实现类型没有任何要求，所以可以把任何之都赋给空接口类型

```
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = time.Second   // error 实现中缺少write方法


var any interface{}
any = true
any = 123
any = map[string]int{"123": 3}
```

接口类型的值包含两部分：一个**具体类型**和**该类型的一个值**，分别称为**动态类型**和**动态值**

![image](https://github.com/rbmonster/file-storage/blob/main/golang-note/basic/interfacevalue.png)

图中分别为声明一个接口及为接口赋值
```
var w io.Writer
w = os.Stdout
```

**接口值**：接口可以直接用 `==` 比较，如果两个接口值都是nil或者二者的动态值相等，那么两个接口值相等

[comment]: <> (> 若一个变量仅进行参数声明，未进行赋值，当调用`w != nil` 的防御性检查将失效)

**类型断言**：检查作为操作数的动态类型是否满足指定的断言类型，类似于x(T)
> 如果断言类型是个具体类型，那么类型断言会检查x的动态类型是否就是T\
> 如果断言类型是接口类型，那么类型断言检查x的动态类型是否满足T\
> 断言中T是否使用`*` 在于支持的类型实现的方法，是否为指针方法，当然编译器会进行优化
```
var w io.Writer
w = os.Stdout
f := w.(*os.File)
c := w.(*bytes.Buffer) // 执行报错，持有的接口为 *os.file

c, ok := w.(*bytes.Buffer)  // 通过ok返回值确定类型是否正确，直接避免报错

if !ok {
    fmt.Println(c)
}
```


类型分支：
1. 子类型多态：如http.Handler与sort.Interface 等接口，各种方法突出了满足这个接口的具体类型，但隐藏了各个具体类型的布局和各自特有的功能。
2. 特设多态(可联合识别):充分利用接口值能容纳各种具体类型的能力，把接口作为这些类型的联合来使用。


### http.Handler接口
一个httpserver 服务需要实现handler接口
```
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

database类型实现了handler接口的方法

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }


func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

```


net/http包请求**多工转发器serverMux**，简化URL与处理程序之间的关联
```
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// ⚠️http.HandlerFunc 为一个方法类型， HandlerFunc定义了类型的ServeHTTP方法，因此实现该类型即可实现handler方法
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
```

### Go接口编程的建议
思想转变：强类型语言中，经常需要先定义接口，再实现具体方法，而Go中，先定义接口的行为是不必要的抽象，应该充分利用Go的隐式实现的特性，在有两个或者多个具体类型需要按统一的方式处理时才需要接口。
> 因为接口仅在有两个或者多个类型满足的情况下，所以接口就必然会抽象掉具体的实现细节。这样的设计的结果会出现更简单和更少方法的接口。


## goroutine 和通道
Go 有两种并发风格，一种是共享内存多线程的传统模型，一种是goroutine和通道(channel)，他们支持通信顺序进程(Communication Sequential Process, CSP)，CSP是一种并发的模式，在不同的执行体(goroutine)之间传递值。

### goroutine
每一个并发执行的活动称为goroutine。语法上，一个go语句在普通的函数或者方法调用前加上`go`关键字前缀
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine 调用
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// 实现类似于pending 的等待
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
```

### 通道
通道是goroutine之间的连接，可以让一个goroutine发送特定值到另一个goroutine的通讯机制
> 当通道复制或者作为参数传递到一个函数时，复制的是引用，这样调用者和被调用者都引用同一份数据结构

```
// 创建一个通道传递 int值
ch := make(chan int)
var x := 12
// 通道发送
ch <- x

// 通道接收
y = <- ch 

// close函数关闭通道
close(ch)


func main() {
    // 创建一个通道传递 int值
	ch := make(chan int)
	go f1(ch)
    // 通道接收
	y := <- ch
	fmt.Println(y)
}

func f1(ch chan int) {
	var x = 12
    // 通道发送
	ch <- x
}
```

单向通道类型：即仅仅支持导出发送或接收操作的通道
```
// 只能发送的通道，允许发送不允许接收
chan <- int 

// 只允许接收的通道，不允许发送
<- chan int

func f1(out chan<- int){
}

func f2(in <-chan int){
}
```
#### 缓存通道
**无缓冲通道**上的发送及接收操作都是阻塞的
1. 无缓冲通道的发送会阻塞，直到另一个goroutine在执行对应通道的接收操作才执行完成。
2. 如果接收操作先执行，接收方goroutine将阻塞，直到另一个goroutine在同一通道上发送值。
> 使用无缓冲通道进行的通信将导致发送和接收goroutine同步化
```
ch := make(chan int) // 无缓冲通道
ch := make(chan int, 3) // 容量为3的缓冲通道
```

**缓冲通道**：有一个元素队列，队列的最大长度在创建的时候通过make的容量来设置，如果通道无goroutine接收，超过容量时将阻塞。
> ⚠️不能将缓冲通道当成队列使用，若无元素进行接收，发送者有被永久阻塞的风险
```
ch := make(chan string, 2)
ch <- "A"
ch <- "B"
// 超过容量阻塞
ch <- "C"

//获取通道的容量
cap(ch)

// 获取通道里面的元素数量
len(ch)s
```

无缓冲的通道提供强同步保障，发送和接收都是阻塞的。而对于缓冲通道，发送和接收操作是解耦的
> goroutine 泄漏：指的是一个无缓冲通道，发送响应没有 goroutine进行接收

#### 管道

通道可以用来连接goroutine，`chan`中一个输入一个输出，概念上叫做管道。

```
var ch1 = make(chan int)
// 关闭管道
close(ch1)

// 接收操作的变种，可以检测是否接收成功。true接收成功，false表示当前的接收操作在一个关闭的并且读完的通道上
x, ok <- ch1
```

管道通信关闭，无缓冲通道的阻塞传输，最终保障程序正常传输完成。
```go
package main

import "fmt"

//!+
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

```


### select多路复用
select 关键字一直等待，直到某个通道通知有消息可以执行接收或发送。如果多个情况同时满足，select随机选择一个通道，这样保证每个通道都有同样的机会被选中
>select 想保证非阻塞的通信，正常需要指定一个default方法，用于select轮训通道时的空转
```
select {
  case <- ch1:
    // ...
  case x := <- ch2
    // ...
  case ch3 <- y
    // ...
  default: 
    
// 缓冲通道为1，偶数发送，奇数接收。如果使用的是一个缓冲通道，会导致程序报错deadlock，因为发送跟接收均无其他goroutine交互，程序认为导致死锁
func selectSend() {
    ch := make(chan int, 1)
    for i := 0; i < 10; i++ {
        select {
        case x := <-ch:
            fmt.Println(x)
        case ch <- i:
        }
    }
}
```

### 取消
需要程序自己控制，比如在goroutine中，每次轮训是否取消的状态。在主函数中把需要创建goroutine的资源消耗掉，保证不再创建新的资源


### 应用
**计数信号量**：可以用来控制并发的量
```
var token = make(chan struct{}, 20)
func f() {
    tokens <- struct{}
    
    <- tokens
}
```

一种main函数等待所有goroutine处理完逻辑再结束的控制
```
var finish = 0
var done = make(chan string)

// goroutine ++finish; done<- "deliver message"

for ; finsh>0;finish-- {
    <- done
}
```

`fatal error: all goroutines are asleep - deadlock!`
出现该错误，说明无缓冲队列存在无接收阻塞或者发送阻塞，而当前所有的goroutine 都已经结束。

`tick := time.Tick(1 * time.Second)` 可以用于定期的发送事件

## 使用共享变量实现并发

并发概念：如果无法确定一个goroutine的事件x与另一个goroutine的事件y的先后顺序， 那么这两个事件就是并发的。

Go并发核心：**不要通过共享内存来通信，应该通过通信来共享内存**。使用通道请求来代理一个受限变量的所有访问的goroutine称为该变量的监控goroutine(monitor goroutine)


竞态：多个goroutine按某些交错顺序执行时程序无法给出正确的结果。

数据静态：两个goroutine并发读写同一个变量并且至少其中一个是写入时。\
避免数据竟态：
1. 方法不要修改变量。
2. 避免从多个goroutine访问同一个变量。
3. 允许多个goroutine访问同一变量，但是同一事件只有一个goroutine可以访问。(互斥机制)

**串行受限**：借助通道把共享变量的地址从上一步到下一步，从而在流水线上的多个goroutine之间共享该变量。而该共享变量就受限于流水线的第一步，这种受限就是串行受限


### sync

#### sync.Mutex

`sync.Mutex`:互斥锁，互斥量保护共享变量，**互斥量是不可重入**的。
```
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
// unlock
mu.Unlock()
```

#### sync.RWMutex
`sync.RWMutex`：读写互斥锁，适用于获取读锁且锁竞争激烈比较有优势，因为RWMutex需要更复杂的内部设置工作，所以在竞争不激烈的情况反而比普通的互斥锁慢。

```
var mw sync.RWMutex

mw.RLock() // 读锁
mw.RUnlock()
mw.Lock() // 写锁
mw.Unlock() 
```

#### sync.Once

`sync.Once`:针对一次性初始化问题，Once中包含一个boolean的变量和一个互斥量，boolean变量标志记录是否初始化完成。

```
var once sync.Once

var f = func() {
    fmt.Println("init")
}
// 首次调用布尔变量为false，调用完成后boolean为true
once.Do(f)
```



### 其他

**竞争检测器**：Go命令行中内置的竞争检测功能`-race` 可以帮忙调试检测程序中的竞争情况

## goroutine与线程

goroutine的栈空间是不固定的，它可以按需增大和缩小，栈空间的限制可以达到1GB，比线程典型的固定大小栈高几个数量级。
> 正常的操作系统(OS)线程则是使用固定大小的栈空间，可能造成浪费，而对于DFS的场景又可能导致太小。因此可拓展的栈空间会更灵活。


### goroutine调度
操作系统线程(OS)由内核进行调度，线程切换需要一个完整的上下文切换，切换的操作需要耗费事件
> 上下文切换：保存一个线程的状态到内存，再恢复另一个线程的状态，最后更新调度器的数据结构。涉及内存的局限性、内存的访问数量、访问内存所需的CPU周期数量的增加

Go运行时包含一个自己的调度器，使用`m:n`调度技术(可以复用/调度m个goroutine到n个OS线程)。当一个goroutine调用time.Sleep、被通道阻塞、互斥量操作时，调度器会设置goroutine为休眠模式，并运行其他goroutine直到被唤醒，该**操作不涉及内核环境的切换**。

### `GOMAXPROCS`
`GOMAXPROCS`参数用来确定需要多少个OS线程来同时执行Go代码。默认值是机器上的CPU数量。
> 阻塞在IO、其他系统调用中、调用非Go语言写的函数的goroutine需要一个独立的OS线程，这个线程不再该参数内。

```
for {
    go fmt.Print(0)
    fmt.Print(1)
}
// output～
// 111111100011111100000011...

// GOMAXPROS=1 go run test.go
// 指定一个核心进行运行，每次最多只能一个goroutine运行，因此无法保证0101的输出
```

### goroutine没有标识(局部存储空间)
goroutine无类似于线程的局部存储，如Java中的ThreadLocal，设计决定的。

1. 避免"超距作用"，即函数的行为不仅取决于它的参数，还取决于它的线程标识存储空间，导致函数或方法的行为难以理解。'】
2. 鼓励更简单的编程风格，影响一个函数行为的仅由其参数决定。



## go工具包
- `go env`: 输出与工具链相关生效的环境变量
- `go get`: 包的下载，`-u`可以保证包更新到最新版本
- `go build`: 构建所有需要的包以及包中所需要的依赖，丢弃最终可执行外的代码
- `go install`: 与build类似，但是会保存所有的编译代码
- `go doc`: 工具输出命令行上所指定的内容声明和整个代码注释，`godoc` 可以用来启动一个文档服务器
- `go list`: 列出可用包的信息。可以使用`...`通配符，`-json`可以以json模式输出


## 测试
`go test`工具是Go语言包的测试驱动程序。
- `-v`：可以输出每个包中测试用力的名称和执行时间。
- `-run`：可以使用正则表达式指定运行的测试用力。
- `-coverprofile`：可以显示运行时的覆盖率
> 在一个包目录中，以`_test.go`结尾的文件不是`go build`命令编译的目标，而是`go test`编译的目标

demo
```
package main

import "testing"

func TestWheel_Write(t *testing.T) {

	w2 := Wheel{
		Circle: Circle{
			Point: Point{ X: 1, Y:2},
			Redius: 10,
		},
		Spokes: 12,
		owner: "12123",
	}

	_, err := w2.Write([]byte("asdfa"))
	if err!= nil {
		t.Errorf("error")
	}
}
```

**白盒测试**：测试中遇到与外部交互的如发送邮件等方法，可以在test中将方法替换成假方法，完成测试后再将方法替换回去，保证测试的进行又不与外部交互。


### Benchmark基准测试

Benchmark开头的测试是用来测试某些操作的性能，是在一定的工作负载下检测程序性能的方法。
- `-bench`：指定要运行的基准测试。`.`表示默认所有的基准测试
- `-benchmen`：在报告中包含了内存分配统计数据

```
func BenchmarkWheel_Write(b *testing.B) {
	w2 := Wheel{
		Circle: Circle{
			Point: Point{ X: 1, Y:2},
			Redius: 10,
		},
		Spokes: 12,
		owner: "12123",
	}
	// b.N 表示调用N次
	for i := 0; i < b.N; i++ {
		_, err := w2.Write([]byte("asdfa"))
		if err!= nil {
			b.Errorf("error")
		}
	}
}

// output ~
sanwuhong@sanwudeMacBook-Pro test % go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: golang-note/src/test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkWheel_Write-8         1000000000                0.3061 ns/op          0 B/op          0 allocs/op
PASS
ok      golang-note/src/test   0.603s
```


基准测试对检测具体操作的性能很有用，特别是在优化程序性的过程中，可以从基准测试的结果中，从内存及程序运行角度对程序进行优化。

进一步的优化程序可以进行程序的**性能剖析**
- cpu性能剖析`go test -cpuprofile=cpu.out`
- 堆性能剖析`go test -memprofile=block.out`
- 阻塞性能剖析`go test -blockprofile=block.out`
> 以上生成的报告文件需要pprof工具进行分析

### Example函数
Example的函数作用：
1. 作为文档的使用示例
2. 是通过`go test`可运行的样本
3. 提供手动实验的代码

## 反射
- `reflect.Type`：反射获取具体的值类型
- `reflect.Value`：反射获取具体的值
- `reflect.Value.Set`：反射设置具体的值


反射Type获取Kind方法区分不同类型：基础类型及以下聚合及引用类型
`reflect.Struct`
`reflect.Ptr`
`reflect.Inteface`
`reflect.Map`
`reflect.Slice`
`reflect.Array`

## unsafe包
- `unsafe.SizeOf`:获取参数在内存中的占用字节长度
- `unsafe.Alignof`:获取参数类型所要求的对其方式
- `unsafe.Pointer`:特殊类型的指针，可以存储任何变量地址
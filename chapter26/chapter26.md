# chapter 26

## 01. 指针基础
### 正经版：

&emsp;**指针是什么**：指针是一个变量，存储另一个变量的**内存地址**。通过指针可以间接访问或修改目标变量的值。
```go
var x int = 10
var p *int = &x // p指向x的地址
fmt.Println(*p)  // 输出：10（解引用获取x的值）
```

&emsp;**内存地址**：每个变量在内存中都有一个唯一地址，可通过 **`&`** 获取。
```go
a := 42
fmt.Println(&a) // 输出a的地址（如0xc0000160a0）
```

&emsp;**注意**：不能对字面量（如 **`&"hello"`**、**`&3`**）使用 **`&`**。

### 大白话版：

&emsp;**指针是什么？**

&emsp;&emsp;**打个比方**：指针就像你家地址的小纸条。纸条本身不是你家，但有了地址就能找到你家。

&emsp;&emsp;代码示例：
```go
x := 10         // 你家的房子（变量）
p := &x         // 写有你家的地址的小纸条（指针）
fmt.Println(*p) // 用地址找到你家，输出：10
```

&emsp;&emsp;**指针有什么用？**

&emsp;&emsp;&emsp;**省内存**：传地址比传整个房子（大对象）更高效。

&emsp;&emsp;&emsp;**改原值**：通过地址直接改房子里的东西（修改变量值）。
```go
*p = 20       // 通过地址直接把你家的电视换成新的
fmt.Println(x) // 输出：20
```

## 02. 解引用操作符 `*`
### 正经版：

&emsp;通过 **`*`** 可以获取指针指向的值，与 **`&`** 互为反向操作。
```go
p := &x
*p = 20       // 修改x的值为20
fmt.Println(x) // 输出：20
```

## 03. 指针类型
### 正经版：

&emsp;**声明指针类型**：**`*int`** 表示指向 **`int`** 的指针。
```go
var p *int     // 声明一个int指针
var s *string  // 声明一个string指针
```

&emsp;**类型限制**：指针只能指向声明类型的变量。
```go
var x int = 10
var p *int = &x // 合法
var s *string = &x // 编译错误（类型不匹配）
```

### 大白话版：

&emsp;**指针的类型**

&emsp;&emsp;**只能指向对应类型**：地址纸条上必须写清楚是“谁家”（比如 **`*int`** 只能指向 **`int`** 类型）。
```go
var s string = "hello"
var p *int = &s // 报错！纸条写的是“int家地址”，但s是“string家”
```

## 04. 指针的赋值与相等性
### 正经版：

&emsp;**赋值**：指针赋值是复制地址值，而非数据。
```go
p1 := &x
p2 := p1      // p2和p1指向同一地址
*p2 = 30      // 修改后x变为30
```

&emsp;**相等性**：两个指针指向同一地址时相等。
```go
p1 := &x
p2 := &x
fmt.Println(p1 == p2) // 输出：true（指向同一变量）
```

## 05. 指针在函数中的应用
### 正经版：

&emsp;**传值 vs 传指针**：Go函数默认传值（副本），通过指针可修改原变量。
```go
func modifyValue(x int) {
    x = 100 // 修改副本，不影响原值
}

func modifyPointer(x *int) {
    *x = 100 // 通过指针修改原值
}

func main() {
    a := 10
    modifyValue(a)
    fmt.Println(a)  // 输出：10
    modifyPointer(&a)
    fmt.Println(a)  // 输出：100
}
```

### 大白话版：

&emsp;**函数传指针 vs 传值**

&emsp;&emsp;**传值**：给你一个房子的复印件，你随便涂改不影响原房。
```go
func 改复印件(a int) {
    a = 100 // 改的是复印件
}
```

&emsp;&emsp;**传指针**：给你原房的地址，你可以直接进去装修。
```go
func 改原房(a *int) {
    *a = 100 // 通过地址直接改原房
}
```

## 06. 结构体与指针
### 正经版：

&emsp;**自动解引用**：访问结构体指针字段时，无需显式解引用。
```go
type Person struct {
    Name string
}

p := &Person{Name: "Alice"}
fmt.Println(p.Name) // 等价于 (*p).Name
```

### 大白话版：

&emsp;**结构体的指针**

&emsp;&emsp;**自动找门牌号**：拿地址纸条就能直接看房，不用敲门（自动解引用）。
```go
type 房子 struct {
    面积 int
}
p := &房子{面积: 100}
fmt.Println(p.面积) // 直接看房，输出：100（不用写成(*p).面积）
```

## 07. 数组与指针
### 正经版：

&emsp; **指向数组的指针**：可以通过指针操作数组，但Go中数组和指针是独立类型。
```go
arr := [3]int{1, 2, 3}
p := &arr
fmt.Println((*p)[0]) // 输出：1
fmt.Println(p[0])    // 自动解引用，输出：1
```

## 08. 切片与指针
### 正经版：

&emsp;**切片底层结构**：切片包含指向底层数组的指针、长度和容量。
```go
s := []int{1, 2, 3}
modifySlice(s) // 函数内部可修改底层数组
```

&emsp;**修改切片本身**：需传递切片指针。
```go
func resize(s *[]int) {
    *s = append(*s, 4) // 修改切片长度
}
```

### 大白话版：

&emsp;**数组和切片**

&emsp;&emsp;**数组指针**：指向整个小区（数组），但Go不让你随便溜达（禁止指针运算）。
```go
arr := [3]int{1, 2, 3}
p := &arr
fmt.Println(p[0]) // 输出：1（自动找第一户）
```

&emsp;&emsp;**切片**：自带小区地图（底层数组指针），传切片就能改小区里的房子。
```go
func 装修小区(s []int) {
    s[0] = 100 // 直接改底层数组
}
```

## 09. 方法接收者的一致性
### 正经版：

&emsp;**指针接收者**：若部分方法使用指针接收者，所有方法都应统一。
```go
type Counter struct {
    count int
}

func (c *Counter) Increment() { // 指针接收者
    c.count++
}

func (c *Counter) Get() int { // 也使用指针接收者
    return c.count
}
```

### 大白话版：

&emsp;**方法接收者的坑**

&emsp;&emsp;**统一用指针**：如果有的方法用指针，有的用值，会精分。
```go
type 电饭煲 struct{}

func (d *电饭煲) 煮饭() {} // 用指针
func (d 电饭煲) 保温() {}  // 用值——会报警告！
```

&emsp;&emsp;**结论**：要么全用指针，要么全用值。

## 10. 接口实现
### 正经版：

&emsp;**指针与非指针类型**：若值类型实现接口，指针类型自动实现。
```go
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() { // 值接收者
    fmt.Println("Woof!")
}

func main() {
    var s Speaker = &Dog{} // 合法
    s.Speak()
}
```

### 大白话版：

&emsp;**接口的玄学**

&emsp;&emsp;**值实现接口，指针也行**：如果你家狗会叫（值方法），它的照片（指针）也能叫。
```go
type 狗 struct{}
func (d 狗) 叫() { fmt.Println("汪！") }

var 宠物 会叫的 = &狗{} // 合法
宠物.叫()            // 输出：汪！
```

&emsp;&emsp;**反过来不行**：如果方法是指针接收者，值不能当接口用。

### 大白话版：

&emsp;**指针核心作用**：省内存、改原值。

&emsp;**避坑指南**：

&emsp;&emsp;**1. `&`** 不能对字面量用（比如 **`&3`** 是错的）。

&emsp;&emsp;**2.** 指针类型必须匹配（ **`*int`** 只能指向 **`int`** ）。

&emsp;&emsp;**3.** 方法接收者要统一风格。

&emsp;&emsp;**4.** 切片传值就能改底层数据，数组不行。

&emsp;**终极口诀：指针是地址，用星号取值；传参想改原，记得带地址！**
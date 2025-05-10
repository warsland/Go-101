# chapter 23

## 01. 结构体与组合
### 正经版：

&emsp;Go语言通过**组合（Composition）**而非继承来实现代码复用。组合的本质是将多个结构体组合在一起，形成更复杂的结构。

&emsp;例如：
```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person  // 嵌入Person结构体
    Salary float64
}
```

&emsp;&emsp;**`Employee`** 嵌入了 **`Person`** ，因此可以直接访问 

&emsp;&emsp;**`Person`** 的字段（如 **`Name`**、**`Age`**）和方法。

组合后的结构更清晰，相关字段被分组到不同的结构体中。

### 大白话版：

&emsp; **组合是什么？**

&emsp;&emsp;想象你搭积木：

&emsp;&emsp;&emsp;先搭一个「人」的积木块，有名字和年龄。

&emsp;&emsp;&emsp;再搭一个「员工」的积木块，把「人」的积木块直接塞进去，再加一个工资属性。

&emsp;&emsp;**代码示例**：
```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person  // 直接塞进去（嵌入）
    Salary float64
}
```

&emsp;&emsp;**用法**：
```go
emp := Employee{Person: Person{Name: "小明", Age: 25}, Salary: 8000}
fmt.Println(emp.Name) // 直接访问人的名字（不用写成 emp.Person.Name）
```

&emsp;&emsp;**核心**：组合就是“塞进去就能直接用”，不用继承那么麻烦。

## 02. 方法转发（Method Forwarding）
### 正经版：

&emsp;嵌入的结构体的方法会被**自动提升**到外层结构体中，可以直接调用。

&emsp;例如：
```go
func (p Person) SayHello() {
    fmt.Println("Hello, I'm", p.Name)
}

func main() {
    emp := Employee{Person: Person{Name: "Alice"}, Salary: 5000}
    emp.SayHello() // 直接调用嵌入结构体的方法
}
```

&emsp;**`Employee`** 实例可以直接调用 **`Person`** 的 **`SayHello`** 方法，无需通过 **`emp.Person.SayHello()`**。

### 大白话版：

&emsp; **方法转发：直接“继承”能力**

&emsp;&emsp;如果「人」会打招呼，那「员工」自动也会！
```go
// 给「人」加一个方法
func (p Person) SayHi() {
    fmt.Println("你好，我是", p.Name)
}

// 员工直接继承了这个方法
emp := Employee{Person: Person{Name: "小明"}}
emp.SayHi() // 输出：你好，我是小明
```

&emsp;&emsp; **重点**：嵌入后，员工自动获得人的所有技能（方法），不用重新写。

## 03. 嵌入的字段名规则
### 正经版：

&emsp;嵌入时不需要指定字段名，默认使用类型名作为字段名。

&emsp;可以直接访问嵌入结构体的字段：
```go
emp.Name = "Bob" // 等价于 emp.Person.Name
```

## 04. 嵌入任意类型
### 正经版：

&emsp;Go允许嵌入**任何类型**（包括基本类型、接口、自定义类型等）：
```go
type Counter int

func (c *Counter) Increment() {
    *c++
}

type App struct {
    Counter // 嵌入Counter类型
}
```

&emsp;&emsp;**`App`** 实例可以直接调用 **`Counter`** 的 **`Increment`** 方法。

### 大白话版：

&emsp; **能嵌入任何东西？**

&emsp;&emsp;是的！比如嵌入一个“计数器”：
```go
type Counter int

func (c *Counter) Add() {
    *c++
}

type App struct {
    Counter // 嵌入计数器
}

func main() {
    app := App{}
    app.Add() // 直接调用计数器的Add方法
    fmt.Println(app.Counter) // 输出：1
}
```

&emsp;&emsp;**用途**：比如给App加个统计功能，不用重新写计数逻辑。

## 05. 方法冲突与解决
### 正经版：

&emsp;如果多个嵌入类型有**同名方法**，编译器会报歧义错误。例如：
```go
type A struct{}

func (A) Foo() { fmt.Println("A") }

type B struct{}

func (B) Foo() { fmt.Println("B") }

type C struct {
    A
    B
}

func main() {
    c := C{}
    c.Foo() // 编译报错：ambiguous selector c.Foo
}
```

&emsp;解决方法：

&emsp;&emsp;**1.显式指定调用的方法**：
```go
c.A.Foo() // 调用A的Foo方法
c.B.Foo() // 调用B的Foo方法
```

&emsp;&emsp;**2.在外部结构体中重新实现同名方法**：
```go
func (c C) Foo() {
    c.A.Foo() // 手动转发到A的Foo方法
}
```

### 大白话版：

&emsp; **嵌入的坑：名字冲突**

&emsp;&emsp;假设你有两个老师傅：

&emsp;&emsp;&emsp;· 张师傅会修电脑

&emsp;&emsp;&emsp;· 李师傅也会修电脑

&emsp;&emsp;你把他俩都塞进一个「万能师傅」里：
```go
type 张师傅 struct{}
func (张师傅) 修() { fmt.Println("修电脑！") }

type 李师傅 struct{}
func (李师傅) 修() { fmt.Println("也修电脑！") }

type 万能师傅 struct {
    张师傅
    李师傅
}

func main() {
    老王 := 万能师傅{}
    老王.修() // 报错！Go懵了：到底调用谁的修？
}
```

&emsp;&emsp;解决方法：

&emsp;&emsp;**1.指名道姓**：
```go
老王.张师傅.修() // 明确用张师傅的修
老王.李师傅.修() // 明确用李师傅的修
```

&emsp;&emsp;**2.自己重新定义**：
```go
func (w 万能师傅) 修() {
    w.张师傅.修() // 手动选一个
}
```

## 06. 优先级规则
### 正经版：

&emsp;**外层结构体的方法优先级高于嵌入类型**。若外层结构体重写了同名方法，调用时会优先使用外层的方法。

## 总结

&emsp;**组合**：通过嵌入结构体实现代码复用，避免继承的复杂性。

&emsp;**方法自动转发**：嵌入类型的方法可直接调用。

&emsp;**冲突处理**：同名方法需显式指定或重写。

&emsp;**灵活性**：可嵌入任意类型，但需注意实际用途。


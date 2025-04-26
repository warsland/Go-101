# chapter 04

## 01. 作用域基础
### 正经版：

&emsp;**定义**：变量从声明时进入作用域，离开作用域后无法访问。

&emsp;**规则**：Go使用大括号 {} 定义块级作用域（如函数、循环、条件语句）。

```go
func main() {
    x := 10 // 作用域开始
    if true {
        y := 20 // y仅在if块内有效
        fmt.Println(x, y) // 合法
    }
    // fmt.Println(y) // 错误：y已脱离作用域
}
```

### 大白话版：
&emsp;**作用域是什么？**

&emsp;&emsp;简单说：变量就像你家养的小狗，它只能在你的院子里（大括号 {} 里）活动。出了院子，别人就看不见它了。

&emsp;&emsp;**例子**：

```go
func main() {
    dog := "旺财"  // 小狗在院子里（main函数里）
    if true {
        cat := "咪咪" // 小猫在屋里（if块里）
        fmt.Println(dog, cat) // 合法：院子里能看见狗和猫
    }
    // fmt.Println(cat) // 报错！猫在屋里，出了屋就找不到了
}
```

## 02. 简短声明（`:=`）
### 正经版：

&emsp;特点：替代 **`var`**，只能在函数内部使用（如 **`for`**、**`if`**、**`switch`**）。

&emsp;作用域：声明变量的作用域到当前语句块结束。

```go
if x := getValue(); x > 5 { // x在整个if-else块有效
    fmt.Println(x)
} else {
    fmt.Println("x is", x)
}
// fmt.Println(x) // 错误：x已脱离作用域
```
### 大白话版
&emsp;**简短声明（`:=`）**

&emsp;&emsp;干嘛用的：快速给变量起名字，但只能在“小任务”里用（比如 **`if`**、**`for`**、**`switch`**）。

```go
if score := 90; score > 60 { 
    fmt.Println("及格了！分数是", score) // 这里能用score
}
// fmt.Println(score) // 报错！score只在判断及格的时候有效
```

&emsp;&emsp;**注意**：全局变量（包作用域）不能用这个写法，必须老老实实用 **`var`**。

## 03. 包作用域的限制
### 正经版：

&emsp;**包级变量**：在函数外声明的变量（全局变量）必须使用 **`var`**。

```go
var globalVar int = 42 // 正确
// packageVar := 100   // 错误：包作用域不能用 :=
```

## 04. `for` 循环的作用域
### 正经版：

&emsp;**传统循环**：初始化语句中的变量作用域覆盖整个循环。
```go
for i := 0; i < 3; i++ { // i的作用域是整个循环
    fmt.Println(i)
}
// fmt.Println(i) // 错误：i已脱离作用域
```

### 大白话版：
&emsp;**`for`循环的作用域**

&emsp;&emsp;**循环里的变量**：像“临时工”，只在打工期间存在。

```go
for i := 0; i < 3; i++ { 
    fmt.Println(i) // 临时工i在打工（循环）
}
// fmt.Println(i) // 报错！临时工下班了，找不到人了
```

## 05. 作用域大小与代码可读性
### 正经版：

&emsp;**平衡原则**：

&emsp;&emsp;**过小作用域**：减少命名冲突，但可能导致重复代码。
```go
// 重复声明（合法但可能冗余）
for i := 0; i < 3; i++ { /* ... */ }
for i := 0; i < 5; i++ { /* ... */ }
```

&emsp;&emsp;**过大的作用域**：可能导致变量污染，需重构。
```go
func main() {
    var result int // 作用域过大
    if condition {
        result = calculate()
    }
    // ...其他使用result的代码
}
```
### 大白话版：
&emsp;**作用域大小怎么选？**

&emsp;&emsp;**太小的问题**：比如每次循环都重新定义一个变量，虽然安全但可能啰嗦。
```go
for i := 0; i < 3; i++ { ... } // 第一个i
for i := 0; i < 5; i++ { ... } // 第二个i，和第一个没关系
```

&emsp;&emsp;**太大的问题**：比如在函数开头定义变量，后面全用它，可能让代码难懂。
```go
func main() {
    var result int // 这个result能活到整个函数结束
    if 条件 {
        result = 100
    }
    // ...后面可能还有100行代码用result，容易看晕
}
```

## 06. 重构建议
### 正经版：

&emsp;**何时重构**：当重复声明导致代码冗余时，适当提升作用域。
```go
func process() {
    data := loadData() // 提升到函数作用域
    if len(data) > 0 {
        // 使用data
    }
    // 其他代码继续使用data
}
```
### 大白话版：
&emsp;**什么时候要改代码？**

&emsp;&emsp;**重复劳动时**：比如多次写同样的变量声明，可以试试把变量提到更大的作用域。
```go
func 处理数据() {
    data := 获取数据() // 提到函数开头，后面都能用
    if 数据有效 {
        // 用data
    }
    // 其他地方继续用data
}
```

## 常见误区：
&emsp;1.在包作用域使用 `:=`（需改用 `var`）。
&emsp;2.认为 `for` 循环中的变量在循环外仍有效（实际作用域仅限于循环）。

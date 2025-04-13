# chapter 03

## 01. 布尔类型
### 正经版：

&emsp;Go中的布尔类型只有两个值：`true` 和 `false`。
```go
var isTrue bool = true
var isFalse bool = false
```
### 大白话版：
&emsp;**布尔类型**：真和假的开关

&emsp;&emsp;Go语言里的布尔值就像灯泡开关，只有两种状态：

&emsp;&emsp;&emsp;· **`true`**：开灯（真）
&emsp;&emsp;&emsp;· **`false`**：关灯（假）

&emsp;&emsp;比如：`今天下雨吗？` 回答只能是 **`true`**（下了）或 **`false`**（没下）。

## 02. strings.Contains 函数
### 正经版：

&emsp;用于检查字符串是否包含子串，返回布尔值。
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    command := "go outside and play"
    fmt.Println(strings.Contains(command, "outside")) // 输出：true
}
```
### 大白话版：
&emsp;**查单词神器：strings.Contains**

&emsp;&emsp;想在一句话里找有没有某个词？比如在“去外面玩”里找“外面”这个词：
```go
strings.Contains("去外面玩", "外面") // 结果：true（有）
```
&emsp;&emsp;就像老师检查你作文里有没有写“努力”这个词一样。

## 03. 比较运算符
### 正经版：

&emsp;`==`（相等）、`!=`（不等）、`<`、`>`、`<=`、`>=`

&emsp;**Go不允许直接比较不同类型**：
```go
// 错误示例：字符串与数字比较
if "5" == 5 { // 编译错误：类型不匹配
}

// 正确做法（需类型转换）：
num := 5
str := "5"
if strconv.Itoa(num) == str { // 需要导入 "strconv"
    fmt.Println("Equal")
}
```
### 大白话版：
&emsp;**比较运算符** ：数学课代表

&emsp;&emsp;·**`==`**：等于（比如 `5 == 5` → `true`）
&emsp;&emsp;·**`!=`**：不等于（比如 `5 != 3` → `true`）
&emsp;&emsp;·**`>`**、**`<`**、**`>=`**、**`<=`**：和数学里一样（比如 `10 > 5` → `true`）

&emsp;&emsp;**`重点`**：Go语言很严格，不同类型不能直接比！
&emsp;&emsp;比如数字 **`5`** 和字符串 **`"5"`** 不能直接比，得先转成同类型（就像不能直接拿苹果和橘子比谁多）。

## 04. 条件语句 `if`/`else if`/`else`
### 正经版：

```go
x := 10
if x > 10 {
    fmt.Println("x大于10")
} else if x == 10 {
    fmt.Println("x等于10") // 输出：x等于10
} else {
    fmt.Println("x小于10")
}
```
### 大白话版：
&emsp;**条件判断**：if...else就像人生选择

&emsp;&emsp;如果...否则如果...否则...：
```go
钱包余额 := 100
if 钱包余额 > 200 {
    fmt.Println("吃大餐！")
} else if 钱包余额 > 50 {
    fmt.Println("吃沙县！")  // 输出这个
} else {
    fmt.Println("泡面...")
}
```

&emsp;&emsp;就像你妈说：“考90分以上去游乐场，80分以上买玩具，否则在家写作业。”

## 05. 赋值符 `=` vs 相等符 `==`
### 正经版：

&emsp;在条件中使用 `=` 会报错：
```go
if x = 5 { // 错误：非布尔类型（需改为 x == 5）
}
```
### 大白话版：
&emsp;**别手滑**：**`=`** 和 **`==`** 的区别

&emsp;&emsp;·**`=`** 是“赋值”（把右边的值塞给左边），比如 `x = 5`（把5放进x里）。
&emsp;&emsp;·**`==`** 是“判断相等”，比如 `if x == 5`（x是不是等于5？）

&emsp;&emsp;**手滑后果**：
```go
if x = 5 { ... }  // 报错！因为这是赋值，不是判断
```

## 06. 逻辑运算符
### 正经版：

&emsp;**`||`**（或）：任意一个为真则结果为真。

&emsp;**`&&`**（与）：两个都为真才为真。

&emsp;**`!`**（非）：取反。
```go
a, b := true, false
fmt.Println(a || b)  // true
fmt.Println(a && b)  // false
fmt.Println(!a)      // false
```

### 大白话版：
&emsp;**逻辑运算符**：或者、并且、反过来

&emsp;&emsp;·**`||`**（或者）：满足一个就行
&emsp;&emsp;比如：`下雨 || 刮风` → `带伞`（下雨或刮风都带伞）

&emsp;&emsp;·**`&&`**（并且）：两个都要满足
&emsp;&emsp;比如：`有钱 && 有闲` → `去旅游`（钱和时间缺一不可）

&emsp;&emsp;·**`!`**（取反）：
&emsp;&emsp;`!晴天` → `不是晴天`（可能是阴天或下雨）

## 07. switch 语句
### 正经版：

&emsp;基本用法：
```go
day := "Mon"
switch day {
case "Mon":
    fmt.Println("周一")
case "Tue":
    fmt.Println("周二")
default:
    fmt.Println("其他")
}
```

&emsp;条件分支（类似 `if-else`）：
```go
x := 15
switch {
case x > 10:
    fmt.Println("x大于10")
case x < 5:
    fmt.Println("x小于5")
}
```
&emsp;**`fallthrough`** 关键字：强制执行下一个分支（无论条件是否满足）。
```go
num := 1
switch num {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2") // 输出：1 2
}
```

### 大白话版：
&emsp;**switch语句**：选择题神器

&emsp;&emsp;**基础用法**：根据选项选结果
```go
switch 星期几 {
case "周一":
    fmt.Println("不想上班...")
case "周五":
    fmt.Println("快乐摸鱼！")
default:
    fmt.Println("普通一天")
}
```

&emsp;&emsp;**高级用法**：直接写条件（类似if-else）
```go
switch {
case 分数 >= 90:
    fmt.Println("优秀！")
case 分数 >= 60:
    fmt.Println("及格！")
}
```
&emsp;&emsp;**`fallthrough`**：强行执行下一题
&emsp;&emsp;比如答对第一题后，不管下一题对不对，都继续执行：
```go
switch 1 {
case 1:
    fmt.Println("答对了！")
    fallthrough  // 强行执行下一个case
case 2:
    fmt.Println("这题其实没答对，但也被执行了")
}
```


## 08. for 循环
### 正经版：

&emsp;传统形式：
```go
for i := 0; i < 3; i++ {
    fmt.Println(i) // 输出：0 1 2
}
```

&emsp;类似 `while` 的循环：
```go
j := 0
for j < 3 {
    fmt.Println(j) // 输出：0 1 2
    j++
}
```

&emsp;无限循环：
```go
for {
    fmt.Println("无限循环")
    break // 立即退出
}
```

### 大白话版：
&emsp;**for循环**：重复干同一件事

&emsp;&emsp;**正常循环**：数数到3
```go
for i := 1; i <= 3; i++ {
    fmt.Println(i)  // 输出1,2,3
}
```

&emsp;&emsp;**当...的时候循环**：类似while
```go
j := 0
for j < 3 {
    fmt.Println(j)  // 输出0,1,2
    j++
}
```
**无限循环**：停不下来，除非用 **`break`**
```go
for {
    fmt.Println("一直打印...")
    break  // 立刻跳出循环（只打印一次）
}
```

## 注意事项

&emsp;1.**比较不同类型**：Go严格要求类型匹配，需显式转换。
&emsp;2.**`strings.Contains`** 区分大小写：是的，**`"Go"`** 和 **`"go"`**不同。
&emsp;3.**`fallthrough`** 使用场景：通常用于多个分支共享代码的情况（但需谨慎使用）。

# chapter 10

## 01. 变量类型决定操作
### 正经版：

&emsp;Go是**静态类型语言**，变量类型一旦确定就不能改变。比如：
```go
var a int = 10
var b string = "20"
// 直接 a + b 会报错，因为类型不匹配
```

&emsp;如果要对不同类型操作，必须**显式转换**：
```go
sum := a + int(b) // 错误！因为字符串不能直接转数字
// 正确做法需要先通过 strconv 包转换字符串为数字
```

### 大白话版：

&emsp; **变量类型就像「工具用途」**

&emsp;&emsp;**错误例子**：你有一个数字（比如10）和一个文字（比如“20”），直接相加会报错，就像用菜刀拧螺丝——工具不对！

&emsp;&emsp;**正确操作**：必须先把文字“20”转成数字20，再相加。但Go里不能直接转，需要调用工具包（比如 **`strconv.Atoi("20")`**）。

## 02.  类型转换的作用
### 正经版：

&emsp;基本类型转换：比如 **`float64(10)`** 将整数10转为浮点数。

&emsp;字符类型转换：

&emsp;&emsp;**· `rune`**（Unicode字符）转字符串：**`string(rune('A'))`** 得到 **`"A"`**。

&emsp;&emsp;**· `byte`**（ASCII字符）转字符串：**`string(byte(65))`** 也得到 **`"A"`**。

### 大白话版：

&emsp; **类型转换就是「换马甲」**

&emsp;&emsp;**例子1**：字符'A'（ASCII码65）转成字符串 → **`string(65)`** 会变成"A"，而不是"65"！

&emsp;&emsp;**例子2**：想把数字65变成字符串"65"？必须用工具：**`strconv.Itoa(65)`**。

## 03. 数字转字符串的陷阱
### 正经版：

&emsp;直接转换数字到字符串会按ASCII码处理：
```go
num := 65
s1 := string(num) // 结果是 "A"（ASCII 65对应字符A）
```

&emsp;如果想将数字65转成字符串"65"，需要用 **`strconv.Itoa`**：
```go
s2 := strconv.Itoa(65) // 结果是 "65"
```

## 04. math包的常量
### 正经版：

&emsp;**`math.MaxInt16`** 和 **`math.MinInt16`** 是无类型常量，用于检查转换是否越界：
```go
value := 32768
if value > math.MaxInt16 {
    fmt.Println("超过int16范围！")
}
// 如果将32768转为int16会溢出，结果不可预测
```

### 大白话版：

&emsp; **数字范围检查像「尺子刻度」**

&emsp;&emsp;**`math.MaxInt16 = 32767`**（最大刻度）

&emsp;&emsp;**用法**：如果你想存一个数到int16类型里，先看看这个数有没有超过32767，超过就存不下（会出错或乱码）。

## 05. 变量的静态类型
### 正经版：

&emsp;变量声明后类型不可变：
```go
var x int = 10
x = "hello" // 报错！不能从int改为string
```

### 大白话版：

&emsp; **变量类型一旦定死，不能改！**

&emsp;&emsp;**例子**：你养了一只狗，名字叫 **`x`**，类型是狗。

&emsp;&emsp;**错误操作**：突然想让 **`x`** 变成猫 → **`x = "猫"`**（报错！Go不允许）。

## 06. 打印布尔值
### 正经版：

&emsp;使用 **`fmt.Print`** 系列函数时，布尔值会转为文本：
```go
fmt.Print(true)  // 输出 "true"
fmt.Printf("%v", false) // 输出 "false"
```

### 大白话版：

&emsp; **打印布尔值就是「开关显示」**

&emsp;&emsp;**`fmt.Print(true)`** → 打印出**true**（就像开关显示“开”）

&emsp;&emsp;**`fmt.Print(false)`** → 打印出**false**（就像开关显示“关”）。
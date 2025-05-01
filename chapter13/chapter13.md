# chapter 13

## 01. 声明新类型
### 正经版：

&emsp;**作用**：让代码更清晰，防止类型误用。
```go
type Celsius float64   // 新类型，表示摄氏度
type Fahrenheit float64 // 新类型，表示华氏度
```

&emsp;虽然 **`Celsius`** 和 **`Fahrenheit`** 底层都是 **`float64`** ，但它们是不同的类型。

&emsp;直接混合使用会报错：
```go
var c Celsius = 20
var f Fahrenheit = c // 错误！类型不匹配
```

### 大白话版：

&emsp; **声明新类型——给数据“贴标签”**

&emsp;&emsp;**干啥用的？**

&emsp;&emsp;比如你写代码时，有“温度”和“距离”都用数字表示，但直接混用容易搞错。

&emsp;&emsp;这时候可以给它们贴个标签，告诉编译器：“这俩虽然都是数字，但不是一码事！”
```go
type 温度 float64  // 标签：温度
type 距离 float64  // 标签：距离
```

&emsp;&emsp;**实际效果**：
```go
var 今天温度 温度 = 30
var 路长 距离 = 500
今天温度 = 路长  // 报错！编译器：“温度≠距离，别瞎赋值！”
```

## 02. 类型转换
### 正经版：

&emsp;**必须显式转换不同类型**：
```go
// 将Celsius转换为Fahrenheit
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

c := Celsius(25)
f := CToF(c) // 合法转换
```

### 大白话版：

&emsp; **类型转换——“撕标签”再贴**

&emsp;&emsp;**为啥要转？**

&emsp;&emsp;比如你要把“温度”转成“距离”（虽然不合理，但假设需要），必须手动撕掉标签，告诉编译器：“我知道类型不同，但我就要这么用！”
```go
温度值 := 温度(25)
距离值 := 距离(温度值)  // 强行转换（类似“撕掉温度标签，贴距离标签”）
```

## 03. 方法声明
### 正经版：

&emsp;**为类型添加方法：**
```go
// 为Celsius类型添加ToString方法
func (c Celsius) ToString() string {
    return fmt.Sprintf("%.1f°C", c)
}

c := Celsius(25)
fmt.Println(c.ToString()) // 输出：25.0°C
```

&emsp;**· 接收者**：**`(c Celsius)`** 表示这个方法属于 **`Celsius`** 类型。

&emsp;**·** 每个方法只能有一个接收者。

### 大白话版：

&emsp; **方法——给类型“配专属工具”**

&emsp;&emsp;**举个栗子**：

&emsp;&emsp;你定义了一个“钱包”类型，想给它加个“存钱”功能：
```go
type 钱包 struct { 余额 int }

// 给“钱包”配个专属工具：存钱方法
func (w *钱包) 存钱(金额 int) {
  w.余额 += 金额
}

我的钱包 := 钱包{}
我的钱包.存钱(100)  // 调用方法：“我的钱包，执行存钱动作！”
```

&emsp;&emsp;**重点**：

&emsp;&emsp;&emsp;**· `(w *钱包)`** 表示这个方法专属于“钱包”类型。

&emsp;&emsp;&emsp; **·** 就像“电钻是工具箱的工具”，方法是类型的“工具”。

## 04. 方法调用
### 正经版：

&emsp;**通过变量调用方法：**
```go
type Wallet struct {
    balance int
}

// 为Wallet类型添加Deposit方法
func (w *Wallet) Deposit(amount int) {
    w.balance += amount
}

myWallet := Wallet{}
myWallet.Deposit(100) // 正确调用方式
```

### 大白话版：

&emsp; **方法 vs 函数——区别在哪？**

&emsp;&emsp;**函数**：谁都能用，比如“公共扳手”。
```go
func 计算和(a, b int) int { return a + b }
```

&emsp;&emsp;**方法**：专属某个类型，比如“我的钱包的存钱功能”。
```go
func (w 钱包) 查询余额() int { return w.余额 }
```

## 05. 命名冲突规则
### 正经版：

&emsp;**同一个包内不允许重复命名：**
```go
// 错误示例！
func Calculate() {}  // 函数名
type Calculate int   // 类型名（冲突！）
```

### 大白话版：

&emsp; **命名冲突——别重名！**

&emsp;&emsp;**规则**：同一个包里，函数名和类型名不能重复。
```go
func 保存() {}      // ✅ 函数
type 保存 struct {}  // ❌ 报错！和函数名冲突了
```
## 总结

|概念|说明|示例|
|:-----:|:-----:|:-----:|
|类型声明|用 **`type`** 创建新类型|**`type ID string`**|
|类型安全|不同类型需显式转换|**`Fahrenheit(c)`**|
|方法接收者|方法属于某个类型|**`func (u User) Login() {}`**|
|方法调用|通过变量调用|**`user.Login()`**|
|命名冲突|包内函数和类型不能同名|避免同时定义 **`Save`** 函数和类型|

&emsp;通过为特定场景创建类型（如 **`Celsius`** 和 **`ID`** ），代码会更易读且安全。方法让类型具备行为，类似面向对象中的“类方法”。
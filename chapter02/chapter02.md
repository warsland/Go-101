# chapter 02 

## 01. 注释
### 正经版：

&emsp;**单行注释：**// 这是单行注释

&emsp;**多行注释：**
```go
/* 
  这是多行注释
  可以写多行内容
*/
```

### 大白话版：
&emsp;**注释：** 代码里的笔记

&emsp;&emsp;**单行注释：** 写代码时加个 **`//`**，后面写点解释，比如：`// 这里算工资`
&emsp;&emsp;就像在代码里贴个便利贴，给自己或别人看。

&emsp;&emsp;**多行注释：** 用 **`/* */`** 包住一大段说明，比如：
```go
/* 
  这段代码是算税后工资的，
  扣了五险一金和个税 
*/
```
&emsp;&emsp;像在代码里夹一张说明书。

## 02. 输出函数对比
### 正经版：

&emsp;**`Println`**:
&emsp;**特点：** 自动换行，可以输出多个参数（用逗号分隔）。
```go
fmt.Println("Hello", 123, 3.14) // 输出：Hello 123 3.14（自动换行）
```

&emsp;**`Printf`**:
&emsp;**特点：** 需要手动换行（用 **`\n`**），支持格式化占位符  **`%v`**（按顺序替换）。
```go
name := "小明"
age := 20
fmt.Printf("%v今年%v岁\n", name, age) // 输出：小明今年20岁（手动换行）
```

&emsp;**`Print`**:
&emsp;**特点：**不换行，需要手动加 **`\n`**。
```go
fmt.Print("Hello")
fmt.Print("World\n") // 输出：HelloWorld（同一行）
```

### 大白话版：
&emsp;**输出内容到屏幕**

&emsp;&emsp;**`Println`**：直接丢一堆东西，自动换行。
```go
fmt.Println("你好", 100, "块钱")  
// 输出：你好 100 块钱（自动换到下一行）
```
&emsp;&emsp;像发微信，说完一句自动换行。

&emsp;&emsp;**`Printf`**：需要自己挖坑填数据，手动换行。
```go
name := "老王"
fmt.Printf("%v 买了 %v 斤西瓜\n", name, 5)  
// 输出：老王 买了 5 斤西瓜（\n 表示换行）
```
&emsp;&emsp;像填空作文，**`%v`** 是占位符，后面按顺序填。

&emsp;&emsp;**`Print`**：不换行，想换行得自己加 **`\n`**。
```go
fmt.Print("今天")
fmt.Print("天气不错\n")  
// 输出：今天天气不错（连在一起）
```

## 03. 格式化对齐
### 正经版：

&emsp;**`%5v`**：右对齐，左侧补空格（总宽度5）。
&emsp;**`%-5v`**：左对齐，右侧补空格。
```go
fmt.Printf("%5v\n", "Hi")  // 输出：  Hi（左侧3空格）
fmt.Printf("%-5v\n", "Hi") // 输出：Hi   （右侧3空格）
```

### 大白话版
&emsp;**对齐文字：** 排版强迫症

&emsp;&emsp;**`%5v`**：把内容拉到右边，左边补空格（总占5格）。
&emsp;&emsp;`fmt.Printf("%5v", "Hi")` →`   Hi`（左边3空格）

&emsp;&emsp;**`%-5v`**：把内容拉到左边，有边补空格（总占5格）。
&emsp;&emsp;`fmt.Printf("%-5v", "Hi")` →`Hi   `（右边3空格）

## 04. 常量和变量
### 正经版：

&emsp;常量（**`const`**）:
&emsp;&emsp;**不可修改**，声明时必须赋值。
```go
const PI = 3.14
// PI = 3.1415 // 报错：常量不能修改
```

&emsp;变量（**`var`**）:
&emsp;&emsp;**必须声明后使用**，可以修改。
```go
var count int = 10
count = 20 // 正确

// 未声明直接赋值会报错
// age = 20 // 报错：变量未声明
```

&emsp;变量声明的3种方式：
&emsp;&emsp;1.单行声明：
```go
var a int
var b string = "Hello"
```
&emsp;&emsp;2.分组声明：
```go
var (
    x int = 5
    y string
)
```
&emsp;&emsp;3.同行多变量：
```go
var i, j int = 1, 2
```

### 大白话版：
&emsp;**常量和变量**

&emsp;&emsp;常量（**`const`**）：像刻在石碑上的字，不能改。
```go
const 密码 = "123456"
// 密码 = "666"  // 报错！石碑刻好了不能改
```
&emsp;&emsp;变量（**`var `**）：像便利贴，能随时改内容。
```go
var 余额 = 1000
余额 = 800  // 没问题，改就完事了
```
&emsp;&emsp;声明方式：
&emsp;&emsp;&emsp;单行声明：`var 年龄 int = 30`
&emsp;&emsp;&emsp;分组声明（像清单）：
```go
var (
    身高 = 175
    体重 = 70
)
```
&emsp;&emsp;&emsp;一行多个：`var a, b int = 1, 2`

## 05. 不支持前置自增
### 正经版：

&emsp;Go 只支持 **`count++`**，不支持 **`++count`**。
```go
count := 0
count++  // 正确
// ++count // 报错
```

### 大白话版：
&emsp;**自增操作：** 只能后置

&emsp;&emsp;Go 只允许 **`i++`**，不允许 **`++i`**。
```go
i := 0
i++  // 正确：先干活，再加1
// ++i  // 报错！Go 不支持先加再干活
```

## 06. 随机数生成
### 正经版：

&emsp;导入包：**`math/rand`**（导入路径），使用时前缀是 **`rand`**。

&emsp;示例：
```go
import "math/rand"

func main() {
    num := rand.Int() // 生成随机整数
    fmt.Println(num)
}
```

### 大白话版：
&emsp;**生成随机数**

&emsp;&emsp;导入包：**`import "math/rand"`**（路径长，但用的时候直接叫 **`rand`**）

&emsp;&emsp;示例：
```go
num := rand.Int()  // 随机生成一个整数
fmt.Println("幸运数字是：", num)
```
&emsp;&emsp;像抽奖机，每次按按钮出个随机数。

## 07. 总结

| 功能 | 示例/说明 |
|:-------:|:--------:|
| **`Println`** | 自动换行，多参数输出 |
| **`Printf`** | 需占位符 **`%v`**，手动换行 **`\n`** |
| 常量 vs 变量| **`const`** 不可改，**`var`** 可改 |
| 变量声明方式 | 单行、分组、同行多变量 |
| 自增操作 | 只支持 **`count++`**，不支持 **`++count`** |
| 随机数生成 | **`rand.Int()`**（包前缀是 **`rand`**） |

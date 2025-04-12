# chapter 01 

## 01. package关键字
### 正经版：

&emsp;**作用：** 声明当前代码文件所属的包（类似于其他语言的“模块”或“命名空间”）。

&emsp;**规则：** 
&emsp;&emsp; 1.每个Go文件必须以 **`package 包名`** 开头。
&emsp;&emsp; 2.包名通常与文件所在目录名一致（例如目录为 **`utils`**，包名为 *`package utils`*）。
&emsp;&emsp; 3.可执行程序必须属于 **`main`** 包（如 `package main`）。

### 大白话版：

&emsp; **package 是干啥的？**
&emsp;&emsp; 一句话回答：每个Go文件开头必须写这行，相当于说“这个文件属于哪个文件夹”。

&emsp;&emsp; **重点：** 想写一个能直接运行的程序，必须用 package main，否则电脑不知道从哪开始执行！

## 02. import 关键字
### 正经版：

&emsp; **作用：** 导入其他包（例如标准库或第三方包），以便使用其中的函数或变量。

&emsp;**语法：**
```go
import "fmt"      // 导入单个包
import (
    "fmt"
    "math"       // 导入多个包
)
```
&emsp; **调用方式：** 使用 **包名.函数名()**，例如 `fmt.Println("Hello")`。

### 大白话版：

&emsp;&emsp;**import 就是“借东西”** 
&emsp;&emsp;作用：你要用别人写好的功能（比如打印文字），就得先“借”过来。
&emsp;&emsp;**怎么用：**
```go
import "fmt"     // 借一个包
import (
    "fmt"       // 可以一口气借多个
    "math"
)
```
&emsp;&emsp;举个栗子：借了 **fmt** 包，想用它打印文字，就得写 *fmt.Println("你好")*，不能直接写 *Println*，不然电脑不知道你借的是谁家的功能。

## 03. fmt 包
### 正经版：

&emsp;**功能：** 提供格式化输入输出函数 如 *`Println`*, *`Printf`*, *`Scan`* 等）。

&emsp;**示例：**
```go
fmt.Println("Hello, World!")  // 输出并换行
fmt.Printf("Value: %d", 42)   // 格式化输出
```

### 大白话版：

&emsp; **fmt 包——你的打印小助手**
&emsp;&emsp;干啥用的：专门负责把文字输出到屏幕，或者从键盘读输入。

&emsp;&emsp;**常用操作：**
```go
fmt.Println("Hello, World!")  // 打印完自动换行
fmt.Printf("Value: %d", 42) // 像C语言那样格式化输出
```

## 04. func关键字
### 正经版：

&emsp;**作用：** 声明函数。

&emsp;**语法：**
```go
func 函数名(参数) 返回值类型 {
    // 函数体
}
```
&emsp;**示例：**
```go
func add(a int, b int) int {
    return a + b
}
```

### 大白话版：

&emsp; **func——定义函数的标志**
&emsp;&emsp;简单说： 写 **`func`** 就是告诉电脑“我要定义一个函数了！”

&emsp;&emsp;**格式：**
```go
func 函数名(参数) 返回值类型 {
    // 具体干啥活
}
```

&emsp;&emsp;**例子：**
```go
func 加法(a int, b int) int {
    return a + b  // 返回a加b的结果
}
```

## 05. 程序入口：main 包与 main 函数
### 正经版：

&emsp;**规则：** 
&emsp;&emsp; 1.可执行程序必须包含 **`main`** 包（`package main`）。
&emsp;&emsp; 2.程序从 **`main`** 包的 `main 函数开始执行。

&emsp;**示例：**
```go
package main

func main() {
    fmt.Println("程序启动！")
}
```

### 大白话版：

&emsp; **程序的起点：main函数**
&emsp;&emsp;硬性规定：
&emsp;&emsp;&emsp; 1.必须有一个包叫 **`main`**（`package main`）。
&emsp;&emsp;&emsp; 2.必须在这个包里写一个 `main()` 函数，电脑从这里开始运行。

&emsp;&emsp;**正确写法：**
```go
package main

func main() {  // 这里开始执行！
    fmt.Println("程序启动！")
}
```
&emsp;&emsp;**错误情况：** 没写 **`main`** 包或 `main`函数？程序直接罢工不运行！

## 06. 调用其他包的函数
### 正经版：

&emsp;**规则：** 必须通过 `包名.函数名()` 的形式调用。

&emsp;**示例：**
```go
import "math"

func main() {
    fmt.Println(math.Sqrt(16)) // 输出：4
}
```

### 大白话版：

&emsp; **用别人家的函数得带“包名”**
&emsp;&emsp;**为啥：** 假设你同时借了 **`math`** 和 **`utils`** 两个包，它们都有 Add 函数，电脑咋知道用哪个？

&emsp;&emsp;**正确姿势：**
```go
import "math"

func main() {
    fmt.Println(math.Sqrt(16)) // 用math包的Sqrt函数算平方根
}
```
&emsp;&emsp;**错误示范：** 直接写 `Sqrt(16)`，电脑会一脸懵逼：“这谁家的Sqrt啊？”

## 07. 大括号 {} 的格式
### 正经版：

&emsp;**强制规则：** 左大括号 **`{`** 必须与关键字（如 **`func`**,**`if`** , **`for`**）在同一行。

&emsp;**正确示例：**
```go
func main() { // 左大括号与 func 同行
    // 代码
}
```

&emsp;**错误示例：**
```go
func main()
{ // 左大括号换行，编译报错！
    // 代码
}
```

### 大白话版：

&emsp; **大括号 { 必须紧贴关键字**
&emsp;&emsp;**强制规定：** 左大括号 **`{`** 必须和 **`func`**、**`if`**、**`for`** 这些关键字同一行，不能换行！

&emsp;&emsp;**正确示范：**
```go
func main() {  // { 和 func 贴在一起
    fmt.Println("对的！")
}
```

&emsp;&emsp;**错误示范：**
```go
func main() 
{               // { 换行了？直接报错！
    fmt.Println("错的！")
}
```

&emsp;&emsp;**类比：**就像写文章必须句号结尾，Go的格式要求很严格，不按规矩写直接编译失败！

## 08.总结
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

&emsp;**运行结果：** 输出 `Hello, Go!`。
&emsp;**注意：** 所有语法细节（如包名、导入、大括号位置）必须严格遵守，否则程序无法编译。
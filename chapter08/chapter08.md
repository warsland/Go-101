# chapter 08

## 01. 指数形式与类型推断
### 正经版：

&emsp;**问题**：输入极大数（如 **`1000000000`**）时，可以用指数形式简化书写（如 **`1e9`**）。

&emsp;**类型推断规则**：

&emsp;&emsp;如果未显式指定类型，Go会将其推断为 **`float64`**（如 **`x := 1e9`** → **`x`** 是 **`float64`** 类型）。

&emsp;&emsp;若需要整数类型（如 **`int64`**），必须显式转换：
```go
x := int64(1e9)  // 显式转换为int64
```

### 大白话版：

&emsp; **写大数别数0，用“科学计数法”**

&emsp;&emsp;比如要写 **`10亿`**，别傻乎乎打9个零（**`1000000000`**），直接写成 **`1e9`**（读作“1乘以10的9次方”）。

&emsp;&emsp;但有个坑：**Go会默认把它当小数！**比如 **` x := 1e9`**，**`x`** 的类型其实是 **`float64`**（带小数点的数）。

&emsp;&emsp;想当整数用？**必须手动转类型**：
```go
x := int64(1e9)  // 现在x是整数类型啦！
```

## 02. big包的三种类型
### 正经版：

|类型	|用途	|示例   |
|:------:|:------:|:------:|
|**`big.Int`**|任意精度的大整数|处理超过 **`int64`** 范围的整数|
|**`big.Float`**|任意精度的浮点数|高精度科学计算|
|**`big.Rat`**|分数（分子/分母形式）|精确表示分数（如 **`1/3`**）|

&emsp;注意事项：

&emsp;&emsp;一旦使用 **`big.Int`**，运算中的所有部分都需用 **`big`** 类型，不能混用原生类型（如 **`int`**）。

&emsp;&emsp;示例：
```go
a := big.NewInt(100)
b := big.NewInt(200)
result := new(big.Int).Add(a, b) // 正确：用big.Int的方法
```

### 大白话版：

&emsp; **超大数？用“大盒子”装！——big包**

&emsp;&emsp;普通数字类型（如 **`int`**）装不下特别大的数，比如“12345678901234567890”，这时得用Go的 **`big`** 包。

&emsp;&emsp;它有三种“大盒子”：

&emsp;&emsp;&emsp;· **`big.Int`** ：装超级大的整数（比如宇宙中的星星数量）

&emsp;&emsp;&emsp;· **`big.Float`** ：装超高精度的小数（比如算圆周率到1000位）

&emsp;&emsp;&emsp;· **`big.Rat`** ：装分数（比如1/3，精确不丢失）

&emsp;&emsp;重点：一旦用了 **`big.Int`**，所有计算都得用它的方法，不能和普通数字混着用！

&emsp;&emsp;示例：
```go
a := big.NewInt(100)      // 用big.Int装100  
b := big.NewInt(200)  
sum := new(big.Int).Add(a, b)  // 必须用Add方法，不能写a + b  
fmt.Println(sum)          // 输出300  
```

## 03. 创建big.Int的两种方式
### 正经版：

&emsp;**方式1：`NewInt` 函数**

&emsp;适用于数值在 **`int64`** 范围内：
```go
num := big.NewInt(123456789) // 输入int64，返回*big.Int
```

&emsp;**方式2：`SetString`方法**

&emsp;适用于超过 **`int64`** 范围的数（通过字符串传入）：：
```go
bigNum := new(big.Int)
bigNum.SetString("123456789012345678901234567890", 10) // 第二个参数是进制（10进制）
```

### 大白话版：

&emsp; **怎么创建“大盒子”？**

&emsp;&emsp;**小技巧1**：数在普通范围内（比如1万），用 **`NewInt`**：
```go
num := big.NewInt(10000)  // 直接装进去  
```

&emsp;&emsp;**小技巧2**：数太大（比如1后面跟30个0），用字符串传进去：
```go
hugeNum := new(big.Int)  
hugeNum.SetString("1000000000000000000000000000000", 10) // 10表示十进制  
```

## 04. 常量与big.Int的关系
### 正经版：

&emsp;**常量的特点**：

&emsp;&emsp;常量可以是“无类型”（如 **`const x = 12345678901234567890`**）。

&emsp;&emsp;如果变量能容纳常量的值，可以直接赋值：
```go
var y int64 = 1e6 // 合法，因为1e6在int64范围内
```

&emsp;**限制**：

&emsp;&emsp;常量和 **`big.Int`不能直接互换**。即使常量是无类型的，也不能直接赋给 **`big.Int`** 。

&emsp;&emsp;必须通过 **`SetString`** 或 **`NewInt`** 转换：
```go
const huge = 12345678901234567890
bigNum := new(big.Int).SetUint64(huge) // 错误！huge超过uint64范围
bigNum.SetString("12345678901234567890", 10) // 正确
```

### 大白话版：

&emsp; **常量的坑：不能直接塞进“大盒子”**

&emsp;&emsp;常量（**`const`**）可以超级大，但想把它装进 **`big.Int`**？直接赋值会炸！

&emsp;&emsp;比如：
```go
const huge = 12345678901234567890  
bigNum := big.NewInt(huge)  // 报错！因为huge超过int64范围  
```
&emsp;&emsp;正确做法：**把常量转成字符串，再用`SetString`**：
```go
bigNum := new(big.Int)  
bigNum.SetString("12345678901234567890", 10)  // 搞定！  
```

## 05. 性能与使用建议
### 正经版：

&emsp;**缺点**：

&emsp;&emsp;**· `big`** 类型操作比原生类型复杂（需调用方法，如 **`Add()`** 、**`Mul()`**）。

&emsp;&emsp;**·** 性能较低，适用于必须处理极大数或高精度的场景。

&emsp;**建议**：

&emsp;&emsp;**·** 优先使用原生类型（如 **`int`**、 **`float64`**），仅在必要时使用 **`big`** 包。

### 大白话版：

&emsp; **为什么不用“大盒子”？——因为它慢啊！**

&emsp;&emsp;**`big.Int`** 这类大盒子虽然能装，但用起来麻烦（得调用方法，比如 **`.Add()`** 、**`.Mul()`**），而且计算速度比普通类型慢。

&emsp;&emsp;**建议**：不到万不得已（比如数实在太大），优先用普通类型！

## 06. 代码示例

```go
package main

import (
    "fmt"
    "math/big"
)

func main() {
    // 使用NewInt
    a := big.NewInt(100)
    b := big.NewInt(200)
    sum := new(big.Int).Add(a, b)
    fmt.Println("Sum:", sum) // 输出 300

    // 使用SetString处理超大数
    hugeNum := new(big.Int)
    hugeNum.SetString("123456789012345678901234567890", 10)
    fmt.Println("Huge Number:", hugeNum)
}
```

## 总结

&emsp;**·** 指数形式方便书写大数，但注意类型推断为 **`float64`**。

&emsp;**·** **`big`** 包用于超大数或高精度计算，但需统一类型。

&emsp;**·** 常量与 **`big.Int`** 需通过字符串或显式方法转换。

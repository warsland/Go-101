# chapter09

## 01. 字符串的声明和零值
### 正经版：

```go
var s string        // 声明后未赋值，s的值为空字符串""
s = "Hello, 世界"   // 双引号表示普通字符串
```

&emsp;Go用双引号 **`""`** 表示字符串，未赋值的字符串变量默认值为 **`""`**（空字符串）。

### 大白话版：

&emsp; **字符串的默认值：空字符串**

&emsp;&emsp;**场景**：你声明了一个字符串变量但没赋值，Go会默认给它一个空值，就是啥也没有的 **`""`**。

```go
var s string  // 此时 s 的值是空字符串，相当于 ""
s = "你好呀"  // 现在 s 有值了
```

## 02. 转义字符 vs 原始字符串
### 正经版：

&emsp;**·普通字符串**（双引号）：支持转义字符（如 **`\n`** 换行、**`\t`** 制表符）。
```go
s1 := "Hello\tWorld\n" // 输出时会显示制表符和换行
```

&emsp;**·原始字符串**（反引号）：直接显示所有内容（包括换行和缩进），不转义。
```go
s2 := `Hello\tWorld
       This is a raw string!`
// \t会原样显示，换行和缩进也会保留
```

### 大白话版：

&emsp; **双引号 vs 反引号：转不转义的区别**

&emsp;&emsp;**·双引号字符串**：会解析转义字符，比如 **`\n`** 换行、**`\t`** 缩进。
```go
s1 := "Hello\tGo!\n第二行"  // 输出时会显示缩进和换行
```

&emsp;&emsp;**·反引号字符串**：所见即所得，不转义任何字符，还能直接换行。
```go
s2 := `Hello\tGo!  
       这是第二行`  
// 输出时会原样显示 \t，换行和缩进也保留
```

## 03. 类型别名
### 正经版：

&emsp;类型别名允许为一个类型起新名字，二者完全兼容：
```go
type MyString = string
var a string = "hello"
var b MyString = a // 可以直接赋值，无需类型转换
```

### 大白话版：

&emsp; **类型别名：给类型起外号**

&emsp;&emsp;**作用**：比如给 **`string`** 起个别名 **`MyString`** ，它俩完全一样，可以互相替换。
```go
type MyString = string  // 别名
var a string = "hello"
var b MyString = a      // 直接赋值，不用转换
```

## 04. 字符与`rune`类型
### 正经版：
 
&emsp;用单引号 **`''`** 表示字符字面量，默认类型是 **`rune`**（表示Unicode码点）：
```go
c := 'A'       // 类型是rune（等同于int32）
fmt.Printf("%c", c) // 用%c打印字符：输出A
```

&emsp; **`rune`** 用于处理Unicode字符（如中文、表情符号），而 **`byte`** 用于处理ASCII字符。

### 大白话版：

&emsp; **字符和rune类型：处理中文和特殊符号**

&emsp;&emsp;单引号字符：用单引号包裹的字符默认是 **`rune`** 类型，能存中文、emoji等。
```go
c := 'A'     // rune类型，存的是字符的Unicode码
d := '中'    // 也能存中文
fmt.Printf("%c", d) // 输出：中
```

## 05. 字符串的访问与不可变性
### 正经版：
 
&emsp;可以通过索引访问字符，但**不能修改**：
```go
s := "hello"
fmt.Println(s[0])   // 输出104（字符'h'的ASCII码）
// s[0] = 'H'       // 错误！字符串不可修改
```

&emsp;修改字符串需先转换为 **`[]rune`** 或 **`[]byte`** ：
```go
rs := []rune(s)
rs[0] = 'H'
s = string(rs)      // 修改后的字符串为"Hello"
```

### 大白话版：

&emsp; **字符串不能直接修改？**

&emsp;&emsp;**原因**：字符串一旦创建，里面的字符就不能改了，想改得先“拆开”。
```go
s := "hello"
// s[0] = 'H'  // 直接改会报错！
// 正确做法：转成字符数组改
arr := []rune(s)
arr[0] = 'H'
s = string(arr)  // 现在 s 是 "Hello"
```

## 06. UTF-8编码处理
### 正经版：
 
&emsp; **`unicode/utf8`** 包提供处理多字节字符的函数：
```go
s := "世界"
fmt.Println(len(s))             // 输出6（字节数）
fmt.Println(utf8.RuneCountInString(s)) // 输出2（字符数）

// 解码第一个字符
r, size := utf8.DecodeRuneInString(s)
fmt.Printf("字符%c，占用%d字节", r, size) // 输出"字符世，占用3字节"
```

### 大白话版：

&emsp; **中文字符长度问题**

&emsp;&emsp;**坑点**：**`len("中文")`** 返回的是字节数（6字节），不是字符数！
```go
s := "世界"
fmt.Println(len(s))            // 输出6（UTF-8中每个中文占3字节）
fmt.Println(utf8.RuneCountInString(s)) // 输出2（实际字符数）
```

## 总结
| 概念 | 说明 |
|:-------:|:--------:| 
| 字符串零值 | 未赋值的字符串变量默认为 **`""`** |
| 原始字符串 | 	用反引号包围，保留所有格式和转义字符原样 |
| **`rune`** 类型 | 用于处理Unicode字符（如中文），等同于 **`int32`** |
| 字符串不可变 | 只能访问字符，修改需转换为 **`[]rune`** 或 **`[]byte`** |
| UTF-8函数 | **`RuneCountInString`** 统计字符数，  **`DecodeRuneInString`** 解码首个字符 |

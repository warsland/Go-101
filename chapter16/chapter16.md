# chapter 16

## 01. 数组的基本特性
### 正经版：

&emsp;**定长且元素类型相同**：数组的长度在声明时确定，不可更改，所有元素类型必须一致。

&emsp;**声明语法**：**`var 数组名 [长度]类型`**

&emsp;示例：**`var numbers [5]int`** 声明一个包含5个整数的数组，初始值为 **`[0, 0, 0, 0, 0]`**（int的零值是0）。

### 大白话版：

&emsp; **数组是什么？**

&emsp;&emsp;**固定大小的储物柜**：想象一个带编号的储物柜，每个柜子只能放**同一种类型**的东西（比如全放书）。柜子数量在买的时候就定好了，之后不能加也不能拆。

&emsp;&emsp;**声明方式**：
```go
var 柜子 [5]int  // 买了一个有5个格子的柜子，默认全塞0（int的默认值）   
```

## 02. 初始化与默认值
### 正经版：

&emsp;**默认零值**：未显式初始化的元素会被赋为类型的零值（如int为0，string为 **`""`**，bool为 **`false`**）。

&emsp;**复合字面量初始化**：
```go
arr := [3]int{1, 2, 3}        // 显式指定长度
arr2 := [...]int{1, 2, 3}     // 使用...自动推断长度
```

&emsp;**多行初始化**（注意逗号）：
```go
arr := [...]int{
    1,
    2,  // 最后一个元素后的逗号不可省略！
}
```

## 03. 索引与越界
### 正经版：

&emsp;索引范围：合法索引为 **`0`** 到 **`len(arr)-1`**。越界访问会导致编译错误或运行时panic。

&emsp;安全遍历：推荐使用 **`range`** 避免越界：
```go
for index, value := range arr {
    fmt.Println(index, value)
}
```

### 大白话版：

&emsp; **怎么装东西？**

&emsp;&emsp;**编号从0开始**：第一个格子叫0号，第二个叫1号……最后一个叫“长度-1”。
```go
fmt.Println(柜子A[0])  // 输出10（第一个格子）  
```

&emsp;&emsp;**越界会爆炸**：比如柜子只有3个格子，你非要开第4个，程序直接崩溃。

&emsp;&emsp;**自动数格子**：用 **`range`** 循环，不用自己算长度：
```go
for 编号, 内容 := range 柜子A {
    fmt.Println("格子", 编号, "里是", 内容)
}
```

&emsp;&emsp;输出：
```
格子 0 里是 10
格子 1 里是 20
格子 2 里是 30
```

## 04. 数组作为值传递
### 正经版：

&emsp;**赋值和传参时复制**：数组是值类型，赋值或传递给函数时会生成完整副本。

&emsp;示例：
```go
a := [3]int{1, 2, 3}
b := a       // b是a的副本
b[0] = 10    // a[0]仍为1
```

### 大白话版：

&emsp; **复制整个柜子**

&emsp;&emsp;**传值就是复制**：如果你把柜子A赋值给柜子B，或者传给函数，Go会**完整复制一份**，两个柜子互不影响。
```go
柜子B := 柜子A  // 复制一份，B和A完全独立
柜子B[0] = 99  // 改B的第一个格子，A的还是10
```

## 05. 数组长度是类型的一部分
### 正经版：

&emsp;**长度不同则类型不同**：**`[3]int`** 和 **`[5]int`** 是两种不同的类型，不可互相赋值或比较。

&emsp;示例：
```go
var a [3]int
var b [5]int
a = b  // 编译错误！
```

### 大白话版：

&emsp; **柜子类型看长度**

&emsp;&emsp;**长度不同就是两种柜子**：比如 **`[3]int`** 和 **`[5]int`** 是两种完全不同的柜子，不能互相赋值。
```go
var 小柜子 [3]int
var 大柜子 [5]int
小柜子 = 大柜子  // 报错！大小不一样，塞不进去
```

## 06. 性能与使用场景
### 正经版：

&emsp;**大数组的性能问题**：传递大数组会产生复制开销，建议改用**切片**（引用类型）或指针：
```go
func modify(arr *[5]int) {
    arr[0] = 100  // 通过指针修改原数组
}
```

&emsp;**动态集合推荐切片**：需要动态调整长度时，优先使用切片（**`[]int`**）。

## 总结
### 正经版：

&emsp;**数组的适用场景**：固定长度、需严格内存控制的场景（如缓存区）。

&emsp;**核心注意事项**：长度不可变、值传递、长度是类型一部分。

&emsp;**替代方案**：灵活操作时使用切片。

### 大白话版：

&emsp; **什么时候用数组？**

&emsp;&emsp;**固定需求**：比如你要存一周7天的温度（每天一个数），用数组刚好。

&emsp;&emsp;**缺点**：长度不能变！想动态增减？用**切片**（Slice）更灵活。


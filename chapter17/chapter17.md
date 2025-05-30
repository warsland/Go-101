# chapter 17

## 01. 切片的基本概念
### 正经版：

&emsp;**定义**：切片是数组的视图（引用），通过半开区间 **`[start:end]`** 创建，包含起始索引，不包含结束索引。

&emsp;&emsp;示例：**`arr := [5]int{1,2,3,4,5}; slice := arr[1:3]`** → 切片包含 **`arr[1]`** 和 **`arr[2]`**。

&emsp;&emsp;省略索引：

&emsp;&emsp;&emsp;**· `arr[:3]`**：从头到索引2。

&emsp;&emsp;&emsp;**· `arr[1:]`**：从索引1到末尾。

&emsp;&emsp;&emsp;**· `arr[:]`**：整个数组。

### 大白话版：

&emsp; **切片就是“动态窗口”**

&emsp;&emsp;**数组**：想象你有一个固定长度的储物柜，比如5个格子，里面放了数字1到5。

&emsp;&emsp;**切片**：就像你拿了一个“可调节的望远镜”对准储物柜的一部分。比如对准第2到第4个格子（索引从0开始），望远镜看到的范围是 **`[1:3]`**，实际拿到的是第2和第3个格子（因为“含头不含尾”）。

```go
arr := [5]int{1,2,3,4,5}
slice := arr[1:3]  // 拿到的是2和3 
```

## 02. 切片与数组的关系
### 正经版：

&emsp;**视图特性**：切片直接引用底层数组。修改切片元素会同步修改原数组，其他引用同一数组的切片也会受影响。

```go
arr := [3]int{1, 2, 3}
slice1 := arr[:]
slice2 := arr[1:]
slice1[0] = 100  // arr[0] 变为 100，slice2[0] 变为 2 → 原数组已变
```

&emsp;**索引限制**：不支持负数索引（如 **`slice[-1]`**），需手动计算：**`slice[len(slice)-1]`**。

### 大白话版：

&emsp; **切片是“实时监控”**

&emsp;&emsp;**共享底层数据**：如果你用切片修改了某个值，原储物柜（数组）和其他对准这个储物柜的切片都会跟着变！
```go
arr := [3]int{1,2,3}
slice1 := arr[:]   // 看到整个储物柜
slice2 := arr[1:]  // 看到第2、3个格子
slice1[0] = 100    // 储物柜变成[100,2,3]，slice2也会变成[2,3]
```

&emsp;&emsp;**不要用负数**：比如 **`slice[-1]`** 会报错，想拿最后一个元素得写 **`slice[len(slice)-1]`**。


## 03. 字符串切片
### 正经版：

&emsp;**不可变性**：字符串是只读的，切片字符串会生成新字符串，与原字符串独立。
```go
s := "hello"
sub := s[1:3]  // "el"
s = "hallo"    // sub 仍为 "el"
```

### 大白话版：

&emsp; **字符串切片是“拍照”**

&emsp;&emsp;切分字符串会生成一个新字符串，和原字符串互不影响，比如：
```go
s := "hello"
sub := s[1:3] // "el"
s = "haha"    // sub还是"el"，不会变
```

## 04. 切片的声明与初始化
### 正经版：

&emsp;**声明**：
```go
var slice []int        // 初始为 nil（未分配底层数组）
slice := make([]int, 5) // 创建长度为5的切片（底层数组初始化为0）
slice := []int{1,2,3}  // 直接初始化
```

&emsp;**动态扩容**：使用 **`append`** 可能触发底层数组重新分配。
```go
arr := [3]int{1,2,3}
slice := arr[:]        // 引用原数组
slice = append(slice, 4) // 超出容量，底层数组更换 → 与原数组解绑
```

### 大白话版：

&emsp; **切片可以“自由伸缩”**

&emsp;&emsp;**动态长度**：切片不需要一开始就固定长度。比如：
```go
// 直接创建一个切片，像个小袋子，随便装东西
slice := []int{1,2,3}  // 装1、2、3
slice = append(slice, 4) // 再塞个4进去，袋子自动变大
```

&emsp;&emsp;**扩容机制**：如果袋子装满了，系统会悄悄换个更大的袋子（新数组），旧袋子（原数组）就和你没关系了。

## 05. 切片的优势与常规字符串操作
### 正经版：

&emsp;**函数传参**：切片是轻量级的（包含指针、长度、容量），传递时避免复制整个数组。

&emsp;**动态长度**：切片的长度不是类型的一部分，可灵活处理不同长度的数据。

&emsp;常用字符串操作：

&emsp;&emsp;**·去除空格**：**`strings.TrimSpace(" hello ")`** → **`"hello"`**。

&emsp;&emsp;**·连接切片**：**`strings.Join([]string{"a", "b"}, "-")`** → **`"a-b"`**。

### 大白话版：

&emsp; **切片的“方便之处”**

&emsp;&emsp;**传参轻便**：传给函数时，切片像是一张“纸条”，只写储物柜的位置、能看到多长，不用把整个储物柜搬过去。

&emsp;&emsp;**操作灵活**：比如去掉字符串两边的空格：
```go
s := "  你好啊  "
trimmed := strings.TrimSpace(s) // 变成"你好啊"
```

&emsp;&emsp;或者把多个字符串用符号连起来：
```go
words := []string{"Go", "真", "香"}
result := strings.Join(words, "-") // "Go-真-香"
```

## 07. 自定义类型与方法

&emsp; 可以为切片或数组定义新类型并绑定方法：
```go
type MySlice []int
func (ms MySlice) Print() {
   fmt.Println(ms)
}
```

## 关键总结

&emsp;**· 共享底层数组**：多个切片可能共享同一数组，修改需谨慎。

&emsp;**· 动态扩容机制**：**`append`** 可能导致底层数组更换。

&emsp;**· 高效传参**：优先使用切片而非数组。
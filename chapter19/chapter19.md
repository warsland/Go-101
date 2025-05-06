# chapter 19

## 01. 映射（`Map`）基础
### 正经版：

&emsp;**作用**：存储键值对，通过键快速查找值。

&emsp;**声明**：**`map[键类型]值类型`**

&emsp;示例：**`var ages map[string]int`**（键是字符串，值是整数）

### 大白话版：

&emsp; **什么是map？**

&emsp;&emsp;**用途**：像电话本，名字（键）找号码（值）。比如存人名和年龄：**`map[string]int`**。

&emsp;&emsp;&emsp;**声明**：必须说清楚键和值的类型。

&emsp;&emsp;&emsp;**`var phoneBook map[string]string`**（键是名字，值是号码）

## 02. 初始化映射
### 正经版：

&emsp;**必须初始化才能使用**，两种方式：

&emsp;&emsp;**1.make函数**：

&emsp;&emsp;&emsp;**`ages := make(map[string]int)`**

&emsp;&emsp;&emsp;可以预分配空间：**`make(map[string]int, 10)`**（初始长度为0，容量10）

&emsp;&emsp;**2.字面量**：

&emsp;&emsp;&emsp;**`ages := map[string]int{"Alice": 30, "Bob": 25}`**



### 大白话版：

&emsp; **初始化才能用！**

&emsp;&emsp;**两种方法**：

&emsp;&emsp;&emsp;**1.直接写值**（像填表格）：

&emsp;&emsp;&emsp;&emsp;**`phoneBook := map[string]string{"张三": "1380000", "李四": "1390000"}`**

&emsp;&emsp;&emsp;**2.用make创建空map**（先拿个空本子）：

&emsp;&emsp;&emsp;&emsp;**`phoneBook := make(map[string]string)`**

## 03. 操作映射
### 正经版：

&emsp;**· 增/改**：**`ages["Charlie"] = 28`**（键不存在则添加，存在则修改）

&emsp;**· 查**：**`age := ages["Alice"]`**

&emsp;**· 删**：**`delete(ages, "Bob")`**

&emsp;**· 判断键是否存在**：
```go
if value, ok := ages["Alice"]; ok {
    fmt.Println("存在，值是", value)
} else {
    fmt.Println("不存在")
}
```

&emsp;&emsp;**`ok`** 为 **`true`** 表示键存在。


### 大白话版：

&emsp; **增删改查，保姆级操作**

&emsp;&emsp;**· 加/改**：**`phoneBook["王五"] = "1350000"`**（有王五就改，没有就加）

&emsp;&emsp;**· 查**：**`number := phoneBook["张三"]`**

&emsp;&emsp;**· 删**：**`delete(phoneBook, "李四")`**（直接划掉）

&emsp;&emsp;**· 判断有没有这个人**：
```go
if number, ok := phoneBook["赵六"]; ok {
    fmt.Println("赵六的号码是", number)
} else {
    fmt.Println("查无此人！")
}
```

&emsp;&emsp;&emsp;**`ok`** 是 **`true`** 说明找到了，**`false`** 就是没这人。

## 04. 重要特性
### 正经版：

&emsp;**零值问题**：访问不存在的键会返回值类型的零值（如int返回0）。需要用 **`ok`** 区分“键不存在”和“值本身就是零”。

&emsp;**引用类型**：赋值或传参时共享底层数据，修改会影响所有副本。
```go
m1 := map[string]int{"a": 1}
m2 := m1
m2["a"] = 100  // m1["a"]也会变成100
```

### 大白话版：

&emsp; **坑点注意！**

&emsp;&emsp;**查不到不报错**：比如查一个不存在的键，会返回值的“零值”（数字返回0，字符串返回空）。

&emsp;&emsp;&emsp;· 所以必须用 **`ok`** 判断，否则可能误以为“空值”是真实数据。

&emsp;&emsp;**共享数据**：如果复制一个map（比如 **`map2 = map1`** ），改 **`map2`** 的内容，**`map1`** 也会跟着变！（像共用一本通讯录）

## 05. 遍历映射
### 正经版：

&emsp;用 **`range`** 遍历，但顺序是随机的（Go的map无序）：
```go
for key, value := range ages {
    fmt.Println(key, value)
}
```

### 大白话版：

&emsp; **遍历map**

&emsp;&emsp;用 **`range`** 循环，但顺序是乱的（Go的map本来就不保证顺序）：
```go
for name, number := range phoneBook {
    fmt.Println(name, "的号码是", number)
}
```

## 06. 实现集合（Set）
### 正经版：

&emsp;Go没有原生集合，但可以用 **`map`** 模拟：：
```go
set := make(map[string]bool)
set["apple"] = true  // 添加元素
if set["apple"] {    // 检查元素是否存在
    fmt.Println("apple存在")
}
```
&emsp;&emsp;键表示元素，值用 **`bool`** 占位，确保唯一性。

### 大白话版：

&emsp; **集合（Set）怎么玩？**

&emsp;&emsp;用Go没有专门的集合，但用map可以假装：：
```go
// 集合：存不重复的名字
nameSet := make(map[string]bool)
nameSet["小明"] = true  // 添加小明
if nameSet["小明"] {    // 检查小明在不在
    fmt.Println("小明在集合里")
}
```

&emsp;&emsp;&emsp;值用 **`bool`** 占位，只关心键是否存在。

## 总结示例
```go
package main

import "fmt"

func main() {
    // 初始化
    scores := map[string]int{"Alice": 90, "Bob": 85}

    // 添加/修改
    scores["Charlie"] = 95

    // 查（带ok检测）
    if score, ok := scores["Bob"]; ok {
        fmt.Println("Bob的分数是", score)
    }

    // 删
    delete(scores, "Alice")

    // 遍历
    for name, score := range scores {
        fmt.Println(name, "=>", score)
    }
}
```
# chapter 21

## 01. 结构体的基本概念
### 正经版：

&emsp;结构体就像一个小盒子，可以把不同类型的数据打包在一起。比如描述一个人：
```go
type Person struct {
    Name string
    Age  int
}
```

&emsp;**`Name`** 和 **`Age`** 是字段（盒子里的物品）

&emsp;通过点操作符访问：**`person.Name = "Alice"`**

### 大白话版：

&emsp; **结构体是啥？**

&emsp;&emsp;想象你要描述一个“人”，需要记录TA的名字、年龄、身高。

&emsp;&emsp;这时候就可以用结构体——它就像一张表格，把相关的信息打包在一起：
```go
type 人 struct {
    名字 string
    年龄 int
    身高 float64
}
```

&emsp;&emsp;&emsp;**字段**：表格里的每一列（比如“名字”）

&emsp;&emsp;&emsp;**点操作符**：想改年龄？直接 **`小明.年龄 = 18`**

## 02. 初始化结构体的两种方式
### 正经版：

&emsp;**(1) 指定字段名（推荐）**
```go
p := Person{
    Name: "Alice",
    Age:  30,
}
```

&emsp;&emsp;顺序随意，未指定的字段自动赋零值（如int的零值是0）

&emsp;**(2) 按字段顺序初始化**
```go
p := Person{"Alice", 30}
```

&emsp;&emsp;必须严格按照结构体定义的顺序

&emsp;&emsp;适合字段少且结构稳定的情况

### 大白话版：

&emsp; **初始化结构体的两种姿势**

&emsp;&emsp;**姿势一：按字段名填**（推荐）

&emsp;&emsp;就像填表格，想填哪列就填哪列：
```go
小明 := 人{
    名字: "小明",
    年龄: 18,
    身高: 175.5,
}
```

&emsp;&emsp;&emsp;顺序随便，没填的自动填默认值（比如数字默认0，字符串默认空）

&emsp;&emsp;**姿势二：按顺序填**（适合简单情况）

&emsp;&emsp;必须严格按照表格的顺序填，少一个都不行：
```go
小明 := 人{"小明", 18, 175.5} // 名字、年龄、身高
```

## 03. 结构体是值类型
### 正经版：

&emsp;**复制结构体**会生成独立副本，修改互不影响

&emsp;**函数传参**时传递的是副本，函数内修改不影响原结构体（若需要修改原结构体，需使用指针 **`*Person`**）

### 大白话版：

&emsp; **结构体的“分身术”**

&emsp;&emsp;**复制结构体**会生成一个独立副本，修改互不影响：
```go
小明副本 := 小明
小明副本.年龄 = 20 // 原版小明还是18岁
```

&emsp;&emsp;**函数传参**时，传的是副本，函数内改不了原版：
```go
func 过生日(p 人) {
    p.年龄 += 1 // 这里改的是副本，原版小明不变
}
```

&emsp;&emsp;→ 想改原版？传指针：**`func 过生日(p *人)`**

## 04. 结构体切片
### 正经版：

```go
var people []Person
people = append(people, Person{"Alice", 30})
```
&emsp;每个元素都是完整的结构体，避免数据错位问题（比同时维护多个切片更安全）

### 大白话版：

&emsp; **结构体切片：一群人的表格**

&emsp;&emsp;假设要管理一群人的信息：
```go
var 班级 []人
班级 = append(班级, 人{"小明", 18, 175})
班级 = append(班级, 人{"小红", 17, 165})
```

&emsp;&emsp;&emsp;这样比同时维护“名字切片、年龄切片、身高切片”更安全，不容易出错。

## 05. JSON编码与结构体标签
### 正经版：

&emsp;**(1) 基本编码**
```go
type Person struct {
    Name string
    Age  int
}
data, _ := json.Marshal(Person{"Alice", 30})
// 输出：{"Name":"Alice","Age":30}
```

&emsp;&emsp;**字段必须大写开头**（可导出），否则不会被编码

&emsp;**(2) 自定义JSON键名**
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
data, _ := json.Marshal(Person{"Alice", 30})
// 输出：{"name":"Alice","age":30}
```

&emsp;&emsp;使用结构体标签（反引号包裹 **`json:"键名"`**）

&emsp;**(3) 格式化JSON**
```go
data, _ := json.MarshalIndent(Person{"Alice", 30}, "", "  ")
/*
输出：
{
  "name": "Alice",
  "age": 30
}
*/
```

### 大白话版：

&emsp; **结构体转JSON**

&emsp;&emsp;**规则一：字段名必须大写开头**

&emsp;&emsp;Go觉得小写是“私有的”，不让你转成JSON：
```go
type 人 struct {
    Name string // 大写才能转
    Age  int
}
data, _ := json.Marshal(人{"小明", 18})
// 输出：{"Name":"小明","Age":18}
```

&emsp;&emsp;**规则二：自定义JSON键名**（比如用下划线）

&emsp;&emsp;加个“标签”，就像给字段贴个快递单：
```go
type 人 struct {
    Name string `json:"name"` // JSON里显示为name
    Age  int    `json:"age"`
}
data, _ := json.Marshal(人{"小明", 18})
// 输出：{"name":"小明","age":18}
```

&emsp;&emsp;**例子：用户信息转JSON**
```go
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    VIP      bool   `json:"is_vip"` // 标签支持蛇形命名
}

用户 := User{"大锤", "dachui@example.com", true}
data, _ := json.MarshalIndent(用户, "", "  ") // 生成带缩进的JSON
fmt.Println(string(data))
```

&emsp;&emsp;输出：
```go
{
  "username": "大锤",
  "email": "dachui@example.com",
  "is_vip": true
}
```

## 常见问题

&emsp; **为什么JSON字段消失了？** 

&emsp;&emsp;→ 检查字段是否大写开头，或标签是否正确

&emsp;**如何让JSON用下划线命名？**

&emsp;&emsp;→ 使用标签：**`json:"first_name"`**

&emsp;**结构体复制后修改无效？**

&emsp;&emsp;→ 确认是否使用了指针 **`*Struct`**

## 总结

&emsp;结构体：打包数据的容器。

&emsp;初始化：推荐指定字段名的方式。

&emsp;值复制：传递副本，独立修改。想要传原版，需要传指针。

&emsp;JSON处理：注意字段导出和标签。

&emsp;切片管理：用结构体切片代替多个独立切片。
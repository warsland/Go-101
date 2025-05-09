# chapter 22

## 01. 结构体代替类
### 正经版：

&emsp;Go使用**结构体（struct）** 组织数据，类似于其他语言中的“类”。例如：
```go
type Person struct {
    Name string
    Age  int
}
```

### 大白话版：

&emsp; **结构体：Go里的“类”**

&emsp;&emsp;Go没有“类”这个概念，但有个东西叫**结构体（struct）**，你可以把它当成一个“盒子”，里面能装各种数据。比如：
```go
type 人 struct {
    名字 string
    年龄 int
}
```

&emsp;&emsp;这个“人”的盒子里，可以装名字和年龄两个数据。这就是Go版的“类”。

## 02. 方法的绑定
### 正经版：

&emsp;通过定义**接收者（receiver）**，将方法绑定到结构体：
```go
// 值接收者（操作副本）
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s\n", p.Name)
}

// 指针接收者（操作原实例，可修改字段）
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

### 大白话版：

&emsp; **方法：给结构体加功能**

&emsp;&emsp;想给这个“人”添加行为（比如说话），就给他绑个方法。方法分两种：

&emsp;&emsp;&emsp;**值接收者**：操作的是“复制品”，改不了原数据。
```go
func (p 人) 打招呼() {
    fmt.Println("你好，我叫", p.名字)
}
```

&emsp;&emsp;&emsp;**指针接收者**：操作的是“本尊”，能改原数据。
```go
func (p *人) 改年龄(新年龄 int) {
    p.年龄 = 新年龄 // 直接改原数据的年龄
}
```

&emsp;&emsp;**啥时候用指针？**

&emsp;&emsp;&emsp;要改数据时（比如改年龄）。

&emsp;&emsp;&emsp;结构体很大时（用指针省内存）。

## 03. 构造函数
### 正经版：

&emsp;Go没有内置构造函数，但约定用 **`NewType`** 或 **`New`** 函数创建实例：
```go
// 返回指针类型（推荐）
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// 带错误处理的构造函数
func NewPersonWithValidation(name string, age int) (*Person, error) {
    if age < 0 {
        return nil, errors.New("age cannot be negative")
    }
    return &Person{Name: name, Age: age}, nil
}
```

&emsp;命名规范：首字母大写（如 **`NewPerson`**）表示可导出，供其他包使用；小写（如 **`newPerson`**）表示私有。

### 大白话版：

&emsp; **构造函数：造“人”的工厂**

&emsp;&emsp;Go没有专门的构造函数，但大家约定用 **`NewXXX`** 这种名字的函数来造对象。比如：
```go
// 造一个“人”
func 造人(名字 string, 年龄 int) *人 {
    return &人{
        名字: 名字,
        年龄: 年龄,
    }
}
```

&emsp;&emsp;&emsp;**为啥返回指针？** 直接返回结构体本身也行，但用指针更高效（尤其是结构体很大时）。

&emsp;&emsp;&emsp;**带错误检查的构造函数：**
```go
func 严格造人(名字 string, 年龄 int) (*人, error) {
    if 年龄 < 0 {
        return nil, errors.New("年龄不能为负")
    }
    return &人{名字, 年龄}, nil
}
```

## 04. 组合替代继承
### 正经版：

&emsp;通过**结构体嵌入**实现组合，复用其他结构体的字段和方法：
```go
type Employee struct {
    Person   // 嵌入 Person，相当于继承字段和方法
    Company string
}

// 使用示例
emp := Employee{
    Person:  Person{Name: "Alice", Age: 30},
    Company: "Google",
}
emp.SayHello() // 直接调用 Person 的方法
```

### 大白话版：

&emsp; **组合代替继承：拼积木**

&emsp;&emsp;Go没有继承，但可以通过**把结构体当积木拼起来**实现类似功能。比如：
```go
type 员工 struct {
    人      // 把“人”这个结构体嵌进来，相当于继承了名字和年龄
    公司 string
}

// 用的时候：
张三 := 员工{
    人:    人{名字: "张三", 年龄: 25},
    公司: "摸鱼科技",
}
张三.打招呼() // 直接调用“人”的方法，输出“你好，我叫张三”
```

&emsp;&emsp;这样，“员工”就自动拥有了“人”的所有字段和方法，不用重新写一遍。

## 05. 包与构造函数的调用
### 正经版：

&emsp;如果构造函数在包 **`example`** 中命名为 **`NewPerson`**，其他包调用时需要前缀：
```go
import "example"

p := example.NewPerson("Bob", 25)
```
&emsp;&emsp;某些包会将核心构造函数简化为 **`New()`**，例如 **`http.NewRequest()`**。

### 大白话版：

&emsp; **包和导出的规则：大写字母开头才能用**

&emsp;&emsp;如果构造函数叫 **`New人`**，其他包想用的话，必须首字母大写，比如 **`NewPerson`**。

&emsp;&emsp;比如你在包 **`公司`** 里写了个构造函数：
```go
func NewPerson(name string, age int) *人 { ... }
```

&emsp;&emsp;&emsp;其他包调用时：
```go
import "公司"

老王 := 公司.NewPerson("老王", 40)
```

## 关键区别

### 值接收者 vs 指针接收者

|**类型**|	**行为**|	**使用场景**|
|:-----:|:-----:|:-----:|
|**值接收者**|	操作结构体的副本，不影响原实例	| 无需修改原数据时 |
|**指针接收者**|操作原实例，可直接修改字段|需要修改数据或结构体较大时|

##  总结

&emsp;**· 结构体 + 方法** = Go的“类”实现。

&emsp;**· 构造函数**通过约定命名（如 **`NewType`**）创建实例。

&emsp;**· 组合**通过结构体嵌入实现代码复用，而非继承。

&emsp;**· 包级可见性**由函数/结构体名称的首字母大小写决定。
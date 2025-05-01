# chapter 12

## 01. 函数声明
### 正经版：

&emsp;**语法**：**`func 函数名(参数列表) 返回值类型 { ... }`**

&emsp;**参数声明**：变量名在前，类型在后。
```go
func add(x int, y int) int { // 两个int参数
    return x + y
}
```

&emsp;若多个参数类型相同，可简写：
```go
func add(x, y int) int { // x和y都是int
    return x + y
}
```



### 大白话版：

&emsp; **函数就是“做事的工具”**

&emsp;&emsp;**怎么造工具**？ 用 **`func`** 关键字，给工具起个名字，告诉它需要什么材料（参数），最后能做出什么东西（返回值）。
```go
// 加法函数，接收两个int，返回int
func add(a int, b int) int {
    return a + b
}
```

&emsp;&emsp;**简化写法**：如果材料类型一样，可以省着写：
```go
func add(a, b int) int { ... } // a和b都是int
```

## 02. 导出函数/变量
### 正经版：

&emsp;**大写字母开头**的标识符会被导出（其他包可访问）。
```go
// 返回两个int值
func swap(a, b string) (string, string) {
    return b, a
}

// 命名返回值（直接在return时返回）
func split(sum int) (x, y int) {
    x = sum * 4
    y = sum / 2
    return // 自动返回x和y
}
```

### 大白话版：

&emsp; **“大写开头”的才是共享工具**

&emsp;&emsp;比如你在家做了个蛋糕（函数），如果名字是 **` BakeCake()`**，邻居（其他包）就能来借。
如果是 **`bakeCake()`**（小写开头），只有自家人（当前包）能用。
```go
// 共享工具：其他包可以调用
func PublicFunc() { ... }

// 私藏工具：只能自己用
func privateFunc() { ... }
```

## 03. 多返回值
### 正经版：

&emsp;用括号包裹返回值类型，可为返回值命名。
```go
// 返回两个int值
func swap(a, b string) (string, string) {
    return b, a
}

// 命名返回值（直接在return时返回）
func split(sum int) (x, y int) {
    x = sum * 4
    y = sum / 2
    return // 自动返回x和y
}
```

### 大白话版：

&emsp; **函数能“一次返回多个结果”**

&emsp;&emsp;比如你买菜，既能返回总价，又能返回找零：
```go
func buy(price int, money int) (total int, change int) {
    total = price
    change = money - price
    return // 自动返回total和change
}

// 调用：总价20，给了50，得到总价20，找零30
total, change := buy(20, 50)
```

## 04. 可变参数函数
### 正经版：

&emsp;使用 **`...`** 表示可变参数（必须是最后一个参数）。
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// 调用：sum(1, 2, 3) → 6
```

### 大白话版：

&emsp; **“可变参数”**：想塞多少塞多少

&emsp;&emsp;用 **`...`** 表示“随便你传几个参数”，比如算一堆数的总和：
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums { // 挨个加
        total += num
    }
    return total
}

// 用起来超随意：
sum(1, 2, 3)    // 6
sum(10, 20, 30) // 60
```

## 05. 函数隔离与复用
### 正经版：

&emsp;**作用域隔离**：函数内的变量不会影响外部。
```go
func test() {
    x := 10 // 仅在test函数内有效
}
```

&emsp;**无副作用函数**：不修改外部状态，仅依赖输入参数。
```go
// 纯函数：无副作用，输出仅依赖输入
func multiply(a, b int) int {
    return a * b
}
```

### 大白话版：

&emsp; **函数像“独立小房间”**

&emsp;&emsp;**变量隔离**：函数里的变量就像房间里的东西，外面拿不到。
```go
func test() {
    x := 10 // 这个x只在test里有效
}
// 外面打印x会报错！
```

&emsp;&emsp;**无副作用**：理想函数就像计算器，你输入3+5，它只返回8，不会偷偷改其他东西。

## 06. 传值机制
### 正经版：

&emsp;Go默认是**值传递**（函数内修改形参不影响实参）。
```go
func change(num int) {
    num = 10 // 不影响外部的x
}

x := 5
change(x)   // x仍然是5
```

&emsp;若需修改外部变量，需用指针：
```go
func changeByPointer(num *int) {
    *num = 10
}

x := 5
changeByPointer(&x) // x变为10
```

### 大白话版：

&emsp; **传值 vs 传指针**

&emsp;&emsp;**默认是“传复印件”**：函数内部改参数，不影响原件。
```go
func change(num int) {
    num = 10 // 改的是复印件
}

x := 5
change(x)    // x还是5，没变！
```

&emsp;&emsp;**想改原件？传地址！**
```go
func changeReal(num *int) {
    *num = 10 // 通过地址找到原件，改它！
}

x := 5
changeReal(&x) // x变成10了！
```

## 07. 同一包内的函数调用
### 正经版：

&emsp;同一包内的函数可直接调用，无需包名前缀。
```go
// 文件1：math.go（包名为math）
func Add(a, b int) int { ... }

// 文件2：utils.go（同一包math）
func Subtract(a, b int) int {
    return Add(a, -b) // 直接调用Add
}
```

### 大白话版：

&emsp; **同一个小区的工具随便用**

&emsp;&emsp;同一个包（比如都住在“厨房”包）里的函数，互相调用不用加包名：
```go
// 厨房包里的工具1：切菜
func cut() { ... }

// 工具2：炒菜，直接调用切菜
func cook() {
    cut() // 直接写名字就行！
}
```

## 08. 代码拆分与维护
### 正经版：

&emsp;将大型程序拆分为小函数，提高可读性和复用性。

&emsp;例如，处理用户登录的逻辑：
```go
func login(username, password string) bool {
    if validateInput(username, password) {
        return checkCredentials(username, password)
    }
    return false
}

func validateInput(username, password string) bool { ... }
func checkCredentials(username, password string) bool { ... }
```

### 大白话版：

&emsp; **代码要像乐高积木，拆开拼才方便**

&emsp;&emsp;比如“做一顿饭”拆成多个小步骤：
```go
func cookMeal() {
    wash()   // 洗菜
    cut()    // 切菜
    fry()    // 炒菜
}

func wash() { ... }
func cut()  { ... }
func fry()  { ... }
```

&emsp;&emsp;每个小函数只做一件事，修起来也简单！
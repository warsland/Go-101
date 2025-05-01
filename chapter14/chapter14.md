# chapter 14

## 01. 函数作为变量
### 正经版：

&emsp;**声明函数变量**：

&emsp;使用 **`func()`** 声明函数类型，例如：
```go
var myFunc func() int  // 声明一个无参数、返回int类型的函数变量
```

&emsp;赋值时直接使用函数名（不加括号）：
```go
func foo() int { return 42 }
myFunc = foo  // 正确：将函数赋值给变量
```

&emsp;**匿名函数赋值**：

&emsp;可以定义匿名函数直接赋值：
```go
myFunc = func() int { return 100 }
```

### 大白话版：

&emsp; **函数可以当变量用**

&emsp;&emsp;**怎么玩？**

&emsp;&emsp;你可以把函数存到一个变量里，就像存数字或字符串一样。比如：
```go
// 声明一个变量，类型是“无参数、返回int的函数”
var myFunc func() int

// 定义一个普通函数
func sayHello() int {
    return 42
}

// 把函数赋值给变量（注意：不用括号！）
myFunc = sayHello

// 调用变量里的函数
fmt.Println(myFunc()) // 输出：42
```

&emsp;&emsp;**直接塞个匿名函数**

&emsp;&emsp;也可以当场写个没名字的函数塞给变量：
```go
myFunc = func() int {
    return 100
}
fmt.Println(myFunc()) // 输出：100
```

## 02. 函数作为参数和返回值
### 正经版：

&emsp;**函数作为参数**：

&emsp;定义一个接受函数作为参数的函数：
```go
func execute(f func() int) {
    result := f()
    fmt.Println(result)
}
```

&emsp;调用时直接传入函数变量：
```go
execute(myFunc)  // 输出：100
```

&emsp;**函数作为返回值**：

&emsp;函数可以返回另一个函数：
```go
func createFunc() func() int {
    return func() int { return 200 }
}
```

### 大白话版：

&emsp; **函数当参数传给其他函数和返回值**

&emsp;&emsp; **函数当参数传给其他函数**

&emsp;&emsp;&emsp;比如：写一个“万能计算器”

&emsp;&emsp;&emsp;你有一个函数，它接受另一个函数作为参数：
```go
// 定义一个函数，参数是一个“无参数、返回int的函数”
func calculate(f func() int) {
    result := f() // 调用传进来的函数
    fmt.Println("结果是：", result)
}

// 调用时传一个函数进去
calculate(func() int { return 3 + 4 }) // 输出：7
calculate(sayHello)                    // 输出：42
```

&emsp;&emsp; **函数还能当返回值**

&emsp;&emsp;&emsp;比如：工厂模式生成函数

&emsp;&emsp;&emsp;函数可以返回另一个函数，用来动态创建逻辑：
```go
// 返回一个“每次调用都加1”的函数
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 使用
counter := createCounter()
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2（闭包记住了count的值）
```

## 03. 使用 type 定义函数类型
### 正经版：

&emsp;**自定义函数类型**：

&emsp;通过 **`type`** 简化函数类型声明：
```go
type MathFunc func() int  // 定义新类型
var add MathFunc = func() int { return 3 + 4 }
```

&emsp;使用自定义类型作为参数：
```go
func calculate(f MathFunc) {
    fmt.Println(f())
}
```

### 大白话版：

&emsp; **用`type`定义函数类型（装逼用）**

&emsp;&emsp; **为什么用？**

&emsp;&emsp; 让代码更清晰。比如，一堆函数都是“无参数、返回int”，可以给它们起个类型名：
```go
type MathFunc func() int

// 声明变量时更简洁
var add MathFunc = func() int { return 3 + 4 }
var multiply MathFunc = func() int { return 3 * 4 }

// 作为参数也更明确
func calculate(f MathFunc) {
    fmt.Println(f())
}
```
## 04. 闭包与匿名函数
### 正经版：

&emsp;**闭包的定义**：

&emsp;匿名函数可以捕获外部作用域的变量，形成闭包：
```go
func outer() func() {
    x := 0
    return func() {
        x++
        fmt.Println(x)
    }
}

f := outer()
f()  // 输出：1
f()  // 输出：2（闭包修改了外部变量x）
```

&emsp;**闭包的变量引用**：

&emsp;闭包捕获的是变量的引用，而非副本。例如在循环中：
```go
var fs []func()
for i := 0; i < 3; i++ {
    j := i  // 创建局部变量j，避免闭包共享i
    fs = append(fs, func() { fmt.Println(j) })
}
for _, f := range fs {
    f()  // 输出：0、1、2
}
```

### 大白话版：

&emsp; **闭包：函数“记住”外面的变量**

&emsp;&emsp; **闭包是啥？**

&emsp;&emsp;匿名函数（没名字的函数）可以“记住”它外面的变量。比如：
```go
func main() {
    x := 10
    // 匿名函数捕获了外部的x
    addX := func(y int) int {
        return x + y
    }
    fmt.Println(addX(5)) // 输出：15（x=10）
    x = 20               // 修改外部变量
    fmt.Println(addX(5)) // 输出：25（x变成20了！）
}
```

&emsp;&emsp;**坑：循环中的闭包**

&emsp;&emsp;在循环里用闭包要小心！直接这样写会出问题：
```go
var funcs []func()
for i := 0; i < 3; i++ {
    // ❌ 错误写法：所有闭包都引用同一个i
    funcs = append(funcs, func() { fmt.Println(i) })
}
// 调用时，i已经变成3了！
for _, f := range funcs {
    f() // 输出：3、3、3（不是预期的0、1、2！）
}
```

&emsp;&emsp;**正确做法**：在循环内创建一个局部变量：
```go
for i := 0; i < 3; i++ {
    j := i // 每次循环都创建一个新的j
    funcs = append(funcs, func() { fmt.Println(j) })
}
// 现在输出0、1、2
```

## 总结

### 正经版：

&emsp;**函数是一等公民**：可赋值、传参、返回。

&emsp;**闭包**：匿名函数捕获外部变量（引用），适用于动态逻辑。

### 大白话版：

&emsp;**函数当变量用**：把函数存起来，想调就调。

&emsp;**函数当参数/返回值**：写更灵活的代码（比如回调、插件）。

&emsp;**闭包**：函数“偷看”外面的变量，还能改它。

&emsp;**`type`别名**：让代码看起来更专业（其实就是为了少打字）。


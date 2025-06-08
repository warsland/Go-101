# chapter 28

## 01. 错误作为最后一个返回值
### 正经版：

&emsp;在Go中，函数若可能出错，**最后一个返回值**通常是 **`error`** 类型。调用者需立即检查错误，例如：
```go
result, err := someFunction()
if err != nil {
    // 处理错误
}
```

&emsp;**为什么是最后一个返回值？**

&emsp;这是Go的约定俗成，使得代码风格统一，便于阅读和错误检查。

### 大白话版：

&emsp;**错误是函数的“最后一句话”**

&emsp;&emsp;怎么用？

&emsp;&emsp;写函数时，如果可能出错，最后一个返回值专门用来报错。比如：
```go
// 假设这个函数可能会出错
func 除法(a, b int) (结果 int, 错误 error) {
    if b == 0 {
        return 0, errors.New("除数不能是零！") // 报错
    }
    return a / b, nil // 没错就返回nil
}
```

&emsp;&emsp;**调用时马上检查错误**，就像拆快递先看有没有破损：
```go
结果, 错误 := 除法(10, 0)
if 错误 != nil {
    fmt.Println("出错啦：", 错误) // 处理错误
    return
}
fmt.Println("结果是：", 结果)
```

## 02. defer关键字：延迟执行
### 正经版：

&emsp;**`defer`** 用于延迟执行函数或方法，确保在函数返回前执行，常用于资源清理（如关闭文件、释放锁）。
```go
func writeFile() error {
    file, err := os.Create("test.txt")
    if err != nil {
        return err
    }
    defer file.Close() // 确保函数返回前关闭文件
    
    // 写入文件...
    return nil
}
```

&emsp;**特点**：

&emsp;&emsp;多个 **`defer`** 按后进先出顺序执行。

&emsp;&emsp;即使函数中发生 **`panic`** ，**`defer`** 仍会执行。

### 大白话版：

&emsp;**`defer`：最后必须做的事**

&emsp;&emsp;干啥用？

&emsp;&emsp;比如你打开文件写东西，**不管写没写完，最后一定要关文件！** 用 **`defer`** 就能保证：
```go
func 写文件() {
    文件, _ := os.Create("日记.txt")
    defer 文件.Close() // 写在最前面，最后执行
  
    // 写文件内容...
}
```

&emsp;&emsp;**特点**：

&emsp;&emsp;&emsp;多个 **`defer`** 像叠盘子，最后写的先执行（后进先出）。

&emsp;&emsp;&emsp;就算中间程序崩了（比如 **`panic`** ），**`defer`** 也会执行，确保“收尾”。

## 03. 错误类型（`error`）
### 正经版：

&emsp;**`error`** 是一个内置接口：
```go
type error interface {
    Error() string
}
```

&emsp;创建错误：

&emsp;&emsp;使用标准库 **`errors.New()`**：
```go
err := errors.New("文件未找到")
```

&emsp;自定义错误类型（实现 **`Error() string`**）：
```go
type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}

func foo() error {
    return &MyError{"自定义错误"}
}
```

### 大白话版：

&emsp;**错误（error）就是个“能说话的东西”**

&emsp;&emsp;**内置错误**：用 **`errors.New("错误信息")`** 直接生成。

&emsp;&emsp;**自定义错误**：只要你的类型能“说出哪里错了”（实现 **`Error() string方法`** ），就能当错误用。
```go
type 我的错误 struct {
    原因 string
}

func (e *我的错误) Error() string { // 实现Error方法
    return "我的错：" + e.原因
}

func 测试() error {
    return &我的错误{"手滑了"} // 返回自定义错误
}
```

## 04. 类型断言
### 正经版：

&emsp;用于将接口转换为具体类型，判断错误类型：
```go
err := someFunction()
if myErr, ok := err.(*MyError); ok {
    // 处理MyError类型的错误
}
```

&emsp;**应用场景**：针对不同类型的错误进行差异化处理。

## 05. panic与recover
### 正经版：

&emsp;**panic**：触发程序崩溃（类似异常），但**不推荐**常规使用，仅用于不可恢复错误（如程序启动失败）。

&emsp;**recover**：捕获 **`panic`** ，必须在 **`defer`** 中调用：
```go
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    // 可能触发panic的代码
    panic("发生严重错误！")
}
```

&emsp;设计哲学：Go鼓励显式处理错误（通过 **`error`** ），而非依赖 **`panic`** 。

### 大白话版：

&emsp;**`panic`和`recover`：大事不妙！**

&emsp;&emsp;**panic**：程序崩溃（比如代码写错了，除零了）。**尽量别用**，除非真没办法了（比如服务器启动失败）。

&emsp;&emsp;**recover**：救命稻草！在崩溃前做点清理（比如关数据库连接）。**必须和`defer`一起用**。
```go
func 不要慌() {
    defer func() { // defer里调用recover
        if 错误 := recover(); 错误 != nil {
            fmt.Println("救回来了！错误是：", 错误)
        }
    }()
  
    panic("完蛋了！") // 触发崩溃
}
```

## 06. 错误处理最佳实践
### 正经版：

&emsp;**尽早检查错误**：调用函数后立即检查 **`err`**。

&emsp;**隔离可能出错的代码**：将可能出错的逻辑与非出错逻辑分离。

&emsp;**简化错误处理**：通过 **`defer`** 和辅助函数减少重复代码。

&emsp;**批量检查错误**：例如使用 **`errgroup`** 包一次性收集多个错误。


### 大白话版：

&emsp;**错误处理小技巧**

&emsp;&emsp;**提前检查参数**：

&emsp;&emsp;比如函数一开头检查参数是否合法，不合法直接报错，后面代码就不用提心吊胆了。
```go
func 登录(用户名 string) error {
    if 用户名 == "" {
        return errors.New("用户名不能为空！")
    }
    // 后面放心用用户名...
}
```

&emsp;&emsp;**批量检查错误**：

&emsp;&emsp;比如注册时要检查用户名、密码、邮箱，可以一次性收集所有错误，不用一个一个报。

## 对比其他语言（Python、Java）

&emsp;Go没有 **`try-catch`** ，错误需显式检查。

&emsp;**`panic`** 类似异常，但更轻量且不鼓励滥用。

## 完整错误处理流程案例
```go
package main

import (
    "errors"
    "fmt"
)

// 自定义错误
type ValidationError struct {
    Field string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("字段 %s 验证失败", e.Field)
}

func processInput(input string) error {
    if input == "" {
        return errors.New("输入不能为空")
    }
    if len(input) < 5 {
        return &ValidationError{Field: "用户名"}
    }
    return nil
}

func main() {
    err := processInput("abc")
    if err != nil {
        switch e := err.(type) {
        case *ValidationError:
            fmt.Println("验证错误:", e)
        default:
            fmt.Println("通用错误:", e)
        }
    }
}
```
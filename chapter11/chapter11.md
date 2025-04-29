# chapter 11

## 01. `strings.Replace` 函数
### 正经版：

&emsp;**正确函数名：`strings.Replace`**（注意是 **`strings`** 包，非 **`srting`**）

&emsp;**功能**：替换字符串中的子串。

&emsp;**参数**：

&emsp;&emsp;**·`s`**：原始字符串。

&emsp;&emsp;**·`old`**：需要被替换的子串。

&emsp;&emsp;**·`new`**：替换后的新子串。

&emsp;&emsp;**·`n`**：替换次数。若为 **`-1`**，则替换所有匹配项；若为正数（如 **`2`**），则替换前 **`n`** 个匹配项；若为 **`0`**，不替换。

&emsp;**示例：**
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "apple banana apple cherry apple"
	old := "apple"
	new := "orange"

	// 替换第一个匹配项
	result1 := strings.Replace(s, old, new, 1)
	fmt.Println(result1) // 输出: orange banana apple cherry apple

	// 替换所有匹配项
	result2 := strings.Replace(s, old, new, -1)
	fmt.Println(result2) // 输出: orange banana orange cherry orange

	// n=0，不替换
	result3 := strings.Replace(s, old, new, 0)
	fmt.Println(result3) // 输出: apple banana apple cherry apple
}
```

### 大白话版：

&emsp; **`strings.Replace`：批量改错字**

&emsp;&emsp;**功能**：把字符串里的某个词，换成另一个词。

&emsp;&emsp;参数：

&emsp;&emsp;&emsp;原始句子（比如："我有一只狗，狗很可爱"）

&emsp;&emsp;&emsp;要改的词（比如："狗"）

&emsp;&emsp;&emsp;改成什么词（比如："猫"）

&emsp;&emsp;&emsp;改几次（比如：改1次、全改、或者不改）

&emsp;&emsp;举个栗子：

&emsp;&emsp;&emsp;原始句子：**`"苹果 香蕉 苹果 樱桃 苹果"`**

&emsp;&emsp;&emsp;把“苹果”改成“橘子”，改第一个：→ **`"橘子 香蕉 苹果 樱桃 苹果"`**（只改第一个苹果）

&emsp;&emsp;&emsp;全部改成橘子：→ **`"橘子 香蕉 橘子 樱桃 橘子"`**（所有苹果都换）

&emsp;&emsp;&emsp;不改：→ 原样输出。

&emsp;&emsp;小贴士：

&emsp;&emsp;&emsp;如果要全部替换，直接写 **`n=-1`**，别数数了。

&emsp;&emsp;&emsp;如果写 **`n=0`**，就是摆烂，啥也不干。

## 02. `strings.ToUpper` 函数
### 正经版：

&emsp;**正确函数名**：**`strings.ToUpper`**（属于 strings 包）

&emsp;**功能**：将字符串中所有小写字母转换为大写字母。

&emsp;**参数**：**`s`** 为原始字符串。

&emsp;**返回值**：全大写的字符串。

&emsp;**示例**：
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, 世界! Go is fun."
	upper := strings.ToUpper(s)
	fmt.Println(upper) // 输出: HELLO, 世界! GO IS FUN.
}
```
&emsp;注意：非字母字符（如中文、数字、符号）不受影响。

### 大白话版：

&emsp; **`strings.ToUpper`：一键变吼叫模式**

&emsp;&emsp;功能：把字符串里所有字母变成大写，像在吼人一样。

&emsp;&emsp;举个栗子：

&emsp;&emsp;&emsp;原句：**`"Hello, 世界！go is cool."`**

&emsp;&emsp;&emsp;按大写键：**`"HELLO, 世界！GO IS COOL."`**

&emsp;&emsp;&emsp;（中文、符号、数字不受影响，只有字母被吼）

## 总结

&emsp;**包名正确性**：Go语言的字符串操作函数大多位于 **`strings`** 包，需确保导入：
```go
import "strings"
```

&emsp;**`参数逻辑`**：

&emsp;&emsp;**`strings.Replace`** 的 **`n`** 参数控制替换次数：

&emsp;&emsp;&emsp;**`n = -1`**：替换所有。

&emsp;&emsp;&emsp;**`n > 0`**：替换前 **`n`** 个。

&emsp;&emsp;&emsp;**`n = 0`**：不替换。

&emsp;&emsp;**`strings.ToUpper`** 只处理字母字符，其他字符保留原样。

&emsp;**实践建议**：通过编写小型测试代码验证函数行为，加深理解。


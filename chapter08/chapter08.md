在输入巨大的数字时，与其输入多个0，不如使用Go的指数形式写出数值。

如果用户没有显式地为包含指数的数值指定变量类型，则Go将推断其类型为float64.

> big包提供了三种类型：
> 存储大数的big.Int，
> 存储任意精度浮点数的big.Float，
> 存储分数的big.Rat。

一旦决定使用big.Int,就需要在等式的每一部分都使用这种类型，即使已经存在的常量亦是如此。

使用big.Int类型的最基本方法时使用NewInt函数，该函数接受一个int64类型的输入，返回一个big.Int类型的输出。

如果大数超出int64的取值上限，可以通过给定一个string来创建一个对应big.Int类型的值（SetString函数）

虽然大类型可以精确表示任意大小的数值，但是使用起来相较于原生类型要麻烦，且运行速度也相对较慢。

在声明常量时，如果没有显式的指定常量的类型，那么Go语言会直接将其标识为无类型。

如果变量的大小可以容纳常量，则常量可以作为变量的值。

P.S. 常量与big.Int的值无法互换。


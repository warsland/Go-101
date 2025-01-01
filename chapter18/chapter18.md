程序员可以通过切片和内置的append函数，实现按需增长的可变数组长度。

append是一个可变参数函数，可以一次向切片添加多个元素。

切片中可见元素的数量决定了切片的长度。如果切片底层的数组比切片的，那么切片还有容积可供增长。

当切片的底层数组没有足够的空间执行追加操作时，append函数会把切边中包含的元素复制到新分配的数组中，新数组的容积是原数组的两倍,其中额外的分配量可以为append函数提供操作空间。

三索引切分操作可以限制新建切片的容量，同时三索引切片操作不会覆盖底层数组的元素。

当且怕的容积不足执行append操作时，Go必须常见新数组并复制旧数组中的内容。

make函数可以指定切片的长度和容积，实现切片的预分配，从而在切片被填满之前，避免额外的内存分配和数组复制操作。

声明可变参数函数：在函数最后一个形参的名字后面添加省略号。

在调用函数时，可以通过省略号展开切片中的多个元素，并将其作为多个实参传递给函数。
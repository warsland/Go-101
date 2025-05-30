在Go语言中通过结构实现组合，并通过嵌入的语言特性实现方法的转发。

可以通过结构和组合对关联的字段进行分组，使得结构更加清晰。

在使用较小的类型构造需要的组合后，可以为每种类型绑定方法进一步组织代码。

通过转发方法可以使得方法变得容易使用。GO语言可以通过结构嵌入实现自动转发方法。

将类型嵌入结构后，嵌入类型的所有方法都将自动为被嵌入类型所用，且将类型嵌入结构不需要指定字段名，结构会自动为被嵌入的类型生成同名的字段。同时，嵌入还能够使得外部结构可以直接访问内部结构的字段。

除去结构，Go语言支持将其他任意类型嵌入结构。

如果结构中的多种类型都有同名的方法，且并没有使用被嵌入的类型调用任何一个方法，那么Go会指出命名冲突，程序仍可以运行。但是当被嵌入类型调用相关方法时，Go编译器会因为歧义报错。

解决歧义可以直接为类型实现一个同名的方法，类型的方法有衔接高于嵌入类型的其他同名方法。并且可以通过手动转发新的同名方法至指定的嵌入类型。

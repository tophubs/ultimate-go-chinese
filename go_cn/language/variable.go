package main

import "fmt"

func main() {
	// --------------
	// 内置类型
	// --------------

	// 类型提供了完整性和可读性.
	// - 我们分配了多少内存？
	// - 这些内存代表了什么？

	// 类型可以是特定的，比如int32或int64。
	// 例子，
	// - uint8包含一个以10为基数的数字，使用一个字节的内存。
	// - int32包含一个使用4字节内存的以10为基数的数。

	// 当我们声明一个类型而不是非常特定时，比如uint或int，
	// 它会根据构建代码所依据的体系结构进行映射。
	// 在64位操作系统上，int会映射到int64。类似地，在32位操作系统中，它变成了int32。

	// 字母的大小是字母的字节数，它匹配我们的地址大小。
	// 例如，在64位体系结构中，字母大小是64位(8字节)，地址大小是64位，那么整数应该是64位。

	// ------------------
	// 零值概念
	// ------------------

	// 我们创建的每个值都必须初始化。 如果我们初始化，对应的值会被设成零值。
	// 对于整个内存分配，我们将那个位重置为0，下面是具体类型的默认零值。
	// - Boolean false
	// - Integer 0
	// - Floating Point 0
	// - Complex 0i
	// - String "" (空字符串)
	// - Pointer nil

	// 字符串是一系列的uint8类型。
	// 一个字符串是包含两个字段的结构体: 第一个字段是一个指向底层字符数组的指针, 第二个字段是字符串的长度
	// 如果它被设成了零值，它的第一个字段值为nil，第二个为0.

	// ----------------------
	// 声明和初始化
	// ----------------------

	// var是为类型初始化零值的唯一语法。
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// 使用短变量声明操作符，我们可以同时定义和初始化。
	aa := 10
	bb := "hello" // 1st word points to a array of characters, 2nd word is 5 bytes
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// ---------------------
	// 转换和强制转换
	// ---------------------

	// Go没有强制转换，但是有转换。
	// 我们必须分配更多的内存，而不是让编译器假装拥有更多字节。

	// 具体类型转换，如aaa int转换为int32。
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

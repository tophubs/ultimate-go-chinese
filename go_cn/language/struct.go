package main

import (
	"fmt"
	"unsafe"
)

// 不同类型字段表示的struct示例
type example struct {
	flag    bool
	counter int16
	pi      float32
}

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	e byte
	c int8
	a bool
	b int32
	d int64
}

type Part3 struct {
	d int64
	b int32
	a bool
	c int8
	e byte
}

func main() {
	// ----------------------
	// 声明并初始化
	// ----------------------

	//
	// 声明一个示例类型的变量，并将其设置为零值。
	// 例如，我们分配多少内存?
	// 一个bool是1字节，int16是2字节，float32是4字节
	// 放在一起，我们有7个字节。然而，实际的答案是8。
	// 这使我们产生了填充和对齐的新概念。
	// 填充字节位于bool和int16之间。原因在于对齐。

	// 对齐的思想：这种硬件在对齐边界上读取内存效率更高。
	// 具体内存对齐分析参考：https://zhuanlan.zhihu.com/p/53413177

	// 规则 1:
	// 根据特定值的大小，Go决定了我们需要的对齐方式。每两个字节的值必须遵循两个字节的边界。
	// 由于bool值只有1字节并且从地址0开始，所以下一个int16必须从地址2开始。被跳过的地址的字节变成1字节填充。
	// 类似地，如果它是一个4字节的值，那么我们将有一个3字节的填充值。
	var e1 example

	// 展示值。
	fmt.Printf("%+v\n", e1)

	// 规则 2:
	// 最大的字段代表整个结构的填充。我们需要尽可能减少填充量。
	// 始终按字节大小从高到底的顺序排列字段,或者从低到高的顺序排列字段，而不是乱序。
	// 这会减小结构体的内存占用大小，例子如下。

	part1 := Part1{}
	part2 := Part2{}
	part3 := Part3{}

	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
	fmt.Printf("part3 size: %d, align: %d\n", unsafe.Sizeof(part3), unsafe.Alignof(part3))

	// 在本例中，整个结构大小必须遵循一个8字节的值，因为int64是8字节。
	// type example struct {
	//     counter int64
	//     pi      float32
	//     float   bool
	// }

	// 使用example类型结构体初始化并声明一个变量
	// 每一行都以逗号结尾。
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// 展示字段值。
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	// 使用匿名结构体声明并初始化一个变量。
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)

	// ---------------------------
	// 命名类型 vs 匿名类型
	// ---------------------------

	// 如果我们有两个名称类型相同的结构，我们不能赋值一个给另一个。
	// 例如，example1和example2是相同的结构，var ex1 example1, var ex2 example2。
	// ex1 = ex2是不被允许的。我们必须通过执行转换显式地表示ex1 = example1(ex2)。
	// 但是，如果ex是相同匿名结构类型的值(如上面的e3)，那么可以指定ex1 = ex
	var e4 example
	e4 = e3

	fmt.Printf("%+v\n", e4)
}

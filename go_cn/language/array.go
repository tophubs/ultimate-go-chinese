// ---------
// CPU 缓存
// ---------

// Cores DO NOT access main memory directly but their local caches.
// CPU 不会直接访问计算机的主存，而是访问自带的高速缓存。
// 高速缓存中存储的是数据和指令。

// 高速缓存从最快到最慢：L1-> L2-> L3->主内存。

// How do we write code that can be sympathetic with the caching system to make sure that
//我们如何编写对缓存系统表示同情的代码，以确保
// we don't have a cache miss or at least, we minimalize cache misses to our fullest potential?
//我们没有缓存未命中，或者至少没有将缓存未命中最小化到最大程度？

// Processor has a Prefetcher. It predicts what data is needed ahead of time.
//处理器有一个预取器。它可以提前预测需要哪些数据。
// There are different granularity depending on where we are on the machine.
//根据我们在计算机上的位置，可以使用不同的粒度。
// Our programming model uses a byte. We can read and write to a byte at a time. However, from the
//我们的编程模型使用一个字节。我们可以一次读写一个字节。但是，从
// caching system POV, our granularity is not 1 byte. It is 64 bytes, called a cache line. All
//缓存系统POV，我们的粒度不是1个字节。它是64个字节，称为高速缓存行。我们所有的
// memory us junked up in this 64 bytes cache line.
//内存都在这64字节的缓存行中阻塞了。

// Since the caching mechanism is complex, Prefetcher tries to hide all the latency from us.
//由于缓存机制很复杂，因此Prefetcher会尝试向我们隐藏所有延迟。
// It has to be able to pick up on predictable access pattern to data.
//它必须能够采用可预测的数据访问模式。
// -> We need to write code that create predictable access pattern to data
//->我们需要编写代码来创建可预测的数据访问模式

// One easy way is to create a contiguous allocation of memory and to iterate over them.
//一种简单的方法是创建连续的内存分配并对其进行迭代。
// The array data structure gives us ability to do so.
//数组数据结构使我们能够执行此操作。
// From the hardware perspective, array is the most important data structure.
//从硬件角度来看，数组是最重要的数据结构。
// From Go perspective, slice is. Array is the backing data structure for slice (like Vector in C++).
//从Go角度来看，slice是。数组是切片的支持数据结构（如C ++中的Vector）。
// Once we allocate an array, whatever it size, every element is equal distant from other element.
//一旦分配了数组，无论数组大小如何，每个元素与其他元素的距离都相等。
// As we iterate over that array, we begin to walk cache line by cache line. As the Prefetcher see
//当我们遍历该数组时，我们开始逐行遍历缓存行。当Prefetcher看到
// that access pattern, it can pick it up and hide all the latency from us.
//该访问模式时，它可以选择它并向我们隐藏所有延迟。

// For example, we have a big nxn matrix. We do LinkedList Traverse, Column Traverse, and Row Traverse
//例如，我们有一个很大的nxn矩阵。我们进行LinkedList遍历，列遍历和行遍历
// and benchmark against them.
//并以此为基准。
// Unsurprisingly, Row Traverse has the best performance. It walk through the matrix cache line
//毫不奇怪，行遍历具有最佳性能。它通过缓存行
// by cache line and create a predictable access pattern.
//遍历矩阵缓存行并创建可预测的访问模式。
// Column Traverse does not walk through the matrix cache line by cache line. It looks like random
//列遍历不会逐行通过矩阵高速缓存行。看起来像是随机
// access memory pattern. That is why is the slowest among those.
//访问内存模式。这就是为什么其中最慢的。
// However, that doesn't explain why the LinkedList Traverse's performance is in the middle. We
//但是，这并不能解释为什么LinkedList Traverse的性能处于中间位置。
// just think that it might perform as poorly as the Column Traverse.
//我们只是认为它的性能可能与“列遍历”一样差。
// -> This leads us to another cache: TLB - Translation lookaside buffer. Its job is to maintain
//->这将导致我们进入另一个缓存：TLB-转换后备缓冲区。它的工作是维护
// operating system page and offset to where physical memory is.
//操作系统页面并偏移到物理内存所在的位置。

// ----------------------------
// Translation lookaside buffer
// ----------------------------

// Back to the different granularity, the caching system moves data in and out the hardware at 64
// bytes at a time. However, the operating system manages memory by paging its 4K (traditional page
// size for an operating system).
// TLB: For every page that we are managing, let's take our virtual memory addresses because that
// that we use (softwares run virtual addresses, its sandbox, that is how we use/share physical
// memory) and map it to the right page and offset for that physical memory.

// A miss on the TLB can be worse than just the cache miss alone.
// The LinkedList is somewhere in between is because the chance of multiple nodes being on the same
// page is probably pretty good. Even though we can get cache misses because cache lines aren't
// necessary in the distance that is predictable, we probably not have so many TLB cache misses.
// In the Column Traverse, not only we have cache misses, we probably have a TLB cache miss on
// every access as well.

// Data-oriented design matters.
// It is not enough to write the most efficient algorithm, how we access our data can have much
// more lasting effect on the performance than the algorithm itself.
//回到不同的粒度，缓存系统一次以64个字节将数据移入和移出硬件。但是，操作系统通过分页4K（操作系统的传统页大小）来管理内存。
//TLB：对于我们正在管理的每个页面，让我们获取虚拟内存地址，因为
//我们所使用的地址（软件运行虚拟地址，其沙箱，即我们使用/共享物理内存的方式）并将其映射到该物理内存的正确页面和偏移量。
//TLB上的未命中可能比仅缓存未命中更严重。
//LinkedList介于两者之间，是因为多个节点位于同一页上的可能性非常好。即使我们可以得到高速缓存未命中，因为在可预测的距离内
//不需要高速缓存行，但我们可能没有那么多TLB高速缓存未命中。
//在“列遍历”中，不仅我们有缓存未命中，而且每次访问也可能有一个TLB缓存未命中。
//面向数据的设计很重要。
//编写最高效的算法是不够的，
//我们如何访问数据可以比算法本身对性能产生更持久的影响。

package main

import "fmt"

func main() {
	// -----------------------
	// Declare and initialize
	// -----------------------

	// Declare an array of five strings that is initialized to its zero value.
	// Recap: a string is a 2 word data structure: a pointer and a length
	// Since this array is set to its zero value, every string in this array is also set to its
	// zero value, which means that each string has the first word pointed to nil and
	// second word is 0.
	//声明一个由五个字符串组成的数组，该数组初始化为零值。
	//概述：一个字符串是一个2字数据结构：一个指针和一个长度
	//由于此数组设置为零值，因此该数组中的每个字符串也均设置为零值，这意味着每个字符串第一个单词指向nil，第二个单词为0。
	//  -----------------------------
	// | nil | nil | nil | nil | nil |
	//  -----------------------------
	// |  0  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	var strings [5]string

	// At index 0, a string now has a pointer to a backing array of bytes (characters in string)
	// and its length is 5.
	//在索引0处，字符串现在有一个指向字节的后备数组的指针（字符串中的字符）
	//长度为5。

	// -----------------
	// What is the cost?
	// -----------------

	// The cost of this assignment is the cost of copying 2 bytes.
	// We have two string values that have pointers to the same backing array of bytes.
	// Therefore, the cost of this assignment is just 2 words.
	//此分配的成本是复制2个字节的成本。
	//我们有两个字符串值，它们的指针指向相同的字节后备数组。
	//因此，此作业的成本仅为2个字。

	//  -----         -------------------
	// |  *  |  ---> | A | p | p | l | e | (1)
	//  -----         -------------------
	// |  5  |                  A
	//  -----                   |
	//                          |
	//                          |
	//     ---------------------
	//    |
	//  -----------------------------
	// |  *  | nil | nil | nil | nil |
	//  -----------------------------
	// |  5  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	// ---------------------------------
	// Iterate over the array of strings
	// ---------------------------------

	// Using range, not only we can get the index but also a copy of the value in the array.
	// fruit is now a string value; its scope is within the for statement.
	// In the first iteration, we have the word "Apple". It is a string that has the first word
	// also points to (1) and the second word is 5.
	// So we now have 3 different string value all sharing the same backing array.
	//使用范围，不仅可以获取索引，还可以获取数组中值的副本。
	//水果现在是一个字符串值；它的范围在for语句内。
	//在第一次迭代中，我们有单词“ Apple”。它是一个字符串，它的第一个单词
	//也指向（1），第二个单词也指向5。
	//因此，我们现在有3个不同的字符串值，它们共享相同的后备数组。

	// What are we passing to the Println function?
	// We are using value semantic here. We are not sharing our string value. Println is getting
	// its own copy, its own string value. It means when we get to the Println call, there are now
	// 4 string values all sharing the same backing array.
	//我们传递给Println函数的是什么？
	//我们在这里使用值语义。我们不共享我们的字符串值。 Println获取
	//自己的副本，即自己的字符串值。这意味着当我们进入Println调用时，现在有
	//4个字符串值都共享相同的后备数组。

	// We don't want to take an address of a string.
	// We know the size of a string ahead of time.
	// -> it has the ability to be on the stack
	// -> not creating allocation
	// -> not causing pressure on the GC
	// -> the string has been designed to leverage value mechanic, to stay on the stack, out of the
	// way of creating garbage.
	// -> the only thing that has to be on the heap, if anything is the backing array, which is the
	// one thing that being shared
	//我们不想接受字符串的地址。
	//我们提前知道了字符串的大小。
	//->它具有进入堆栈的能力
	//->不创建分配
	//->不会对GC造成压力//
	//>字符串被设计为利用值机制，以留在堆栈上，
	//创建垃圾的方式。
	//->唯一必须在堆上的东西，如果有什么是后备数组，这是//被共享的东西
	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}
	// Declare an array of 4 integers that is initialized with some values using literal syntax.
	//声明一个由4个整数组成的数组，该数组使用文字语法用一些值初始化。
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers using traditional style.
	//使用传统样式遍历数字数组。
	fmt.Printf("\n=> Iterate over array using traditional style\n")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// ---------------------
	// Different type arrays
	// ---------------------

	// Declare an array of 5 integers that is initialized to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initialized with some values.
	four := [4]int{10, 20, 30, 40}

	fmt.Printf("\n=> Different type arrays\n")
	fmt.Println(five)
	fmt.Println(four)
	// When we try to assign four to five like so five = four, the compiler says that
	// "cannot use four (type [4]int) as type [5]int in assignment"
	// This cannot happen because they have different types (size and representation).
	// The size of an array makes up its type name: [4]int vs [5]int. Just like what we've seen
	// with pointer. The * in *int is not an operator but part of the type name.
	//当我们尝试将4分配给5时，例如5 = 4，编译器会说
	////不能在分配中使用4（[4] int）作为[5] int类型”
	//这不会发生，因为它们具有不同的类型（大小和表示形式）。
	//数组的大小组成其类型名称：[4] int与[5] int。就像我们用指针看到的
	//。 * int中的*不是运算符，而是类型名称的一部分。

	// Unsurprisingly, all array has known size at compiled time.

	// -----------------------------
	// Contiguous memory allocations
	// -----------------------------

	// Declare an array of 6 strings initialized with values.
	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}

	// Iterate over the array displaying the value and address of each element.
	// By looking at the output of this Printf function, we can see that this array is truly a
	// contiguous block of memory. We know a string is 2 word and depending on computer
	// architecture, it will have x byte. The distance between two consecutive IndexAddr is exactly
	// x byte.
	// v is its own variable on the stack and it has the same address every single time.
	//遍历显示每个元素的值和地址的数组。
	//通过查看此Printf函数的输出，我们可以看到此数组确实是
	//连续的内存块。我们知道一个字符串是2个字，根据计算机
	//的体系结构，它将有x个字节。两个连续的IndexAddr之间的距离正好是
	//x字节。
	//v是其在堆栈上的变量，并且每次都具有相同的地址。
	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}
}

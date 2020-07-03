# Ultimate Go study guide

- [Ultimate Go study guide中文翻译版](https://github.com/tophubs/ultimate-go-chinese/blob/master/README_CHINESE_DOC.md)
- [Ultimate Go study guide英文版](https://github.com/tophubs/ultimate-go-chinese/blob/master/README.md)

[![Go Report Card](https://goreportcard.com/badge/github.com/hoanhan101/ultimate-go)
](https://goreportcard.com/report/github.com/hoanhan101/ultimate-go)
![](https://img.shields.io/github/stars/hoanhan101/ultimate-go)
![](https://img.shields.io/github/forks/hoanhan101/ultimate-go)
[![hackernews](https://img.shields.io/badge/hackernews-450%2B-orange)](https://news.ycombinator.com/item?id=20701671)
[![r/golang](https://img.shields.io/badge/r/golang-255%2B-orange)](https://www.reddit.com/r/golang/comments/cqqi9h/ultimate_go_study_guides_with_heavily_documented/)
[![r/compsci](https://img.shields.io/badge/r/compsci-60%2B-orange)](https://www.reddit.com/r/compsci/comments/cr3jzh/ultimate_go_study_guides_with_heavily_documented/)
[![r/programming](https://img.shields.io/badge/r/programming-40%2B-orange)](https://www.reddit.com/r/programming/comments/cr3gqu/ultimate_go_study_guides_with_heavily_documented/)

> [101多个编码面试问题，包括详细的解决方案，测试用例和程序分析 →](https://github.com/hoanhan101/algo)

> [添加你的邮件订阅最新最新消息 →](https://tinyletter.com/hoanhan)

<p align="center">
  <img src="gopher.png" alt="gopher" width="300"/>
</p>

## Motivation

这个仓库包含了我学习Go和计算机系统的笔记。不同的人有不同的学习风格。
对我而言，我通过做事和逐步学习实例来学习最好。
因此，我尝试仔细记录笔记并直接在源代码上进行注释，而不是编写Markdown文件。
这样一来，我可以在阅读时理解每一行代码，并牢记幕后的理论。

在整片文档中，我还添加部分我觉得有用的链接在当中

如果您有兴趣获得最近的更新，请添加邮箱快速订阅 [邮箱添加地址 →](https://tinyletter.com/hoanhan)

## 目录

- **设计哲学**:
  [指导方案](https://github.com/ardanlabs/gotraining/blob/master/topics/go/README.md)
- **语言机制**
  - **语法**
    - 变量: [Built-in types | Zero value concept | Initialization | Conversion vs Casting](go_cn/language/variable.go)
    - 结构体: [Initialization | Name type vs Anonymous type](go_cn/language/struct.go)
    - 指针: 
      - [Passing by value | Escape analysis | Stack space | Garbage Collection](go_cn/language/pointer.go)
      - [Golang's Code Review Receiver Type](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)
    - 常量: [Initialization | iota](go_cn/language/constant.go)
    - 函数: [Initialization](go_cn/language/function.go)
  - **数据结构**
    - Array: [CPU Cache | TLB | Initialization | Iteration | Type array | Contiguous memory allocation](go_cn/language/array.go)
    - Slice: [Initialization | Length vs Capacity | Reference Type | Appending | Slice of Slice | Copy of Slice | UTF-8](go_cn/language/slice.go)
    - Map: [Initialization | Iteration | Deleting | Finding | Restriction ](go_cn/language/map.go)
  - **解耦**
    - 方法: 
      - [Value and Pointer Receiver Call](go_cn/language/method_1.go)
      - [Value and Pointer Semantics](go_cn/language/method_2.go)
      - [Methods are just functions | Function variable](go_cn/language/method_3.go)
    - Interface: 
      - [Valueless type | Concrete type vs Interface type | Relationship | Polymorphic function](go_cn/language/interface_1.go)
      - [Interface via Pointer Receiver | Method set | Slice of Interface](go_cn/language/interface_2.go)
    - Embedding: 
      - [Declaring fields, NOT Embedding](go_cn/language/embedding_1.go)
      - [Embedding type | Inner type promotion](go_cn/language/embedding_2.go)
      - [Embedded type and Interface](go_cn/language/embedding_3.go)
      - [Outer and inner type implementing the same Interface](go_cn/language/embedding_4.go)
    - Exporting:
      - [Guideline](go_cn/language/exporting/README.md)
      - [Exported identifier](go_cn/language/exporting/exporting_1)
      - [Accessing a value of an unexported identifier](go_cn/language/exporting/exporting_2)
      - [Unexported fields from an exported struct](go_cn/language/exporting/exporting_3)
      - [Exported types with embedded unexported types](go_cn/language/exporting/exporting_4)
- **Software Design**
  - Composition:
    [Guideline](https://github.com/ardanlabs/gotraining/tree/master/topics/go#interface-and-composition-design)
    - Grouping types: 
      - [Grouping By State](go_cn/design/grouping_types_1.go)
      - [Grouping By Behavior](go_cn/design/grouping_types_2.go)
    - Decoupling: 
      - [Struct Composition](go_cn/design/decoupling_1.go)
      - [Decoupling With Interface](go_cn/design/decoupling_2.go)
      - [Interface Composition](go_cn/design/decoupling_3.go)
      - [Decoupling With Interface Composition](go_cn/design/decoupling_4.go)
    - Conversion: 
      - [Interface Conversions | Type Assertion](go_cn/design/conversion_1.go)
      - [Runtime Type Assertion](go_cn/design/conversion_2.go)
    - Interface Pollution: 
      - [Interface Pollution](go_cn/design/pollution_1.go)
      - [Remove Interface Pollution](go_cn/design/pollution_2.go)
    - Mocking: 
      - [Package To Mock](go_cn/design/mocking_1.go)
      - [Sample Client](go_cn/design/mocking_2.go)
  - Error Handling: 
    - [Default error values](go_cn/design/error_1.go)
    - [Error variables](go_cn/design/error_2.go)
    - [Type as context](go_cn/design/error_3.go)
    - [Behavior as context](go_cn/design/error_4.go)
    - [Finding the bug/pitfall of nil value of error interface](go_cn/design/error_5.go)
    - [Wrapping Errors](go_cn/design/error_6.go)
  - Packaging: [Guideline](https://github.com/ardanlabs/gotraining/blob/master/topics/go_cn/design/packaging/README.md)
  - Dependency management: [Go Modules](https://blog.golang.org/using-go-modules)
- **Concurrency**
  - **Mechanics**
    - Goroutine: 
      - [Go Scheduler Internals](go_cn/concurrency/goroutine_1.go)
      - [Language Mechanics](go_cn/concurrency/goroutine_2.go)
      - [Goroutine time slicing](go_cn/concurrency/goroutine_3.go)
      - [Goroutines and parallelism](go_cn/concurrency/goroutine_4.go)
    - Data race: 
      - [Race Detection](go_cn/concurrency/data_race_1.go)
      - [Atomic Functions](go_cn/concurrency/data_race_2.go)
      - [Mutexes](go_cn/concurrency/data_race_3.go)
      - [Read/Write Mutex](go_cn/concurrency/data_race_4.go)
    - Channel: 
      - [Guideline](https://github.com/ardanlabs/gotraining/tree/master/topics/go#concurrent-software-design)
      - [Language Mechanics | Unbuffered channel: Signaling with(out) data](go_cn/concurrency/channel_1.go)
      - [Unbuffered channel: Double signal | Buffered channel: Close and range | Unbuffered channel: select and receive | Unbuffered channel: select and send | Buffered channel: Select and drop](go_cn/concurrency/channel_2.go)
      - [Unbuffered channel (Tennis match)](go_cn/concurrency/channel_3.go)
      - [Unbuffered channel (Replay race)](go_cn/concurrency/channel_4.go)
      - [Buffered channel: Fan Out](go_cn/concurrency/channel_5.go)
      - [Select](go_cn/concurrency/channel_6.go)
  - **Patterns**
    - Context: 
      - [Store and retrieve values from a context](go_cn/concurrency/context_1.go)
      - [WithCancel](go_cn/concurrency/context_2.go)
      - [WithDeadline](go_cn/concurrency/context_3.go)
      - [WithTimeout](go_cn/concurrency/context_4.go)
      - [Request/Response](go_cn/concurrency/context_5.go)
    - Pattern
      - Task
      - Logger
- **测试用例和分析**
  - 测试: 
    - [基本单元测试](go_cn/testing/basic_test.go)
    - [表格测试](go_cn/testing/table_test.go)
    - [Sub Test](go_cn/testing/sub_test.go)
    - [Web Server](go_cn/testing/web_server)
    - [Mock Server](go_cn/testing/web_test.go)
    - [Test Coverage](go_cn/testing/README.md)
  - 性能基准
    - [Basic Benchmark](go_cn/benchmark/basic_test.go)
    - [Sub Benchmark](go_cn/benchmark/sub_test.go)
  - 模糊测试
    - [Guideline](https://github.com/ardanlabs/gotraining/blob/master/topics/go_cn/testing/fuzzing/README.md)
  - 分析
    - 堆栈跟踪: [Review](go_cn/profiling/stack_trace_1.go) | [Packing](go_cn/profiling/stack_trace_2.go)
    - GODEBUG: [Memory Tracing](go_cn/profiling/memory_tracing.go)

## For more resources:

- [Ultimate Go Programming](https://www.safaribooksonline.com/library/view/ultimate-go-programming/9780134757476/)
- [ardanlabs/gotraining/topics/courses/go](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md)
- [Computer Systems: A Programmer's Perspective](https://www.amazon.com/Computer-Systems-Programmers-Perspective-3rd/dp/013409266X)
- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack)
- [Thoughts on Go performance optimization](https://github.com/dgryski/go-perfbook)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/hoanhan101/ultimate-go.svg)](https://starchart.cc/hoanhan101/ultimate-go)

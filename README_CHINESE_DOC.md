# Ultimate Go study guide

- [Ultimate Go study guide中文翻译版](/README_CHINESE_DOC.md)
- [Ultimate Go study guide英文版](/README.md)

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

这个仓库包含了我学习Go和计算机系统知识的笔记，不同的人有不同的学习风格。
对我而言，我通过直接写代码来学得更好。
因此，我尝试仔细记录笔记并直接在源代码上进行注释，而不是编写Markdown文件。
这样一来，我可以在阅读时理解每一行代码，并牢记幕后的理论。

在整篇文档中，我还添加部分我觉得有用的链接在当中

如果您有兴趣获得最近的更新，请添加邮箱快速订阅 [邮箱添加地址 →](https://tinyletter.com/hoanhan)

## 目录

- **设计哲学**:
  [指导方案](https://github.com/ardanlabs/gotraining/blob/master/topics/go/README.md)
- **语言机制**
  - **语法**
    - 变量: [内置类型 | 零值概念 | 初始化 | 转换 vs 强制转换](go_cn/language/variable.go)
    - 结构体: [初始化 | 命名类型 vs 匿名类型](go_cn/language/struct.go)
    - 指针:
      - [值传递 | 逃逸分析 | 栈空间 | 垃圾回收](go_cn/language/pointer.go)
      - [Golang代码评审之接收者类型](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)
    - 常量: [初始化 | iota](go_cn/language/constant.go)
    - 函数: [初始化](go_cn/language/function.go)
  - **数据结构**
    - Array: [CPU缓存 | TLB(页表缓存) | 初始化 | 迭代 | 数组类型 | 连续内存分配](go_cn/language/array.go)
    - Slice: [初始化 | 长度 vs 容量 | 引用类型 | 切片扩容 | 切片的切片 | 切片拷贝 | UTF-8](go_cn/language/slice.go)
    - Map: [初始化 | 迭代 | 删除 | 查找 | 限制 ](go_cn/language/map.go)
  - **解耦**
    - 方法: 
      - [值接收者和指针接收者调用](go_cn/language/method_1.go)
      - [值和指针语义](go_cn/language/method_2.go)
      - [方法就是函数 | 函数变量](go_cn/language/method_3.go)
    - 接口: 
      - [值类型 | 具体类型 vs 接口类型 | 关系 | 多态函数](go_cn/language/interface_1.go)
      - [指针接收者接口 | 方法设置 | 切片接口](go_cn/language/interface_2.go)
    - 嵌入: 
      - [声明字段而不是嵌入](go_cn/language/embedding_1.go)
      - [嵌入类型 | 内部类型提升](go_cn/language/embedding_2.go)
      - [嵌入类型和接口](go_cn/language/embedding_3.go)
      - [实现相同接口的外部类型和内部类型](go_cn/language/embedding_4.go)
    - 导出:
      - [指导方案](go_cn/language/exporting/README.md)
      - [导出标识符](go_cn/language/exporting/exporting_1)
      - [访问未导出标识符的值](go_cn/language/exporting/exporting_2)
      - [导出结构体的未导出字段](go_cn/language/exporting/exporting_3)
      - [具有嵌入式未导出类型的导出类型](go_cn/language/exporting/exporting_4)
- **软件设计**
  - 组成:
    [指导方案](https://github.com/ardanlabs/gotraining/tree/master/topics/go#interface-and-composition-design)
    - 分组类型: 
      - [按状态分组](go_cn/design/grouping_types_1.go)
      - [按行为分组](go_cn/design/grouping_types_2.go)
    - 解耦: 
      - [结构体组合](go_cn/design/decoupling_1.go)
      - [用接口解耦](go_cn/design/decoupling_2.go)
      - [接口组合](go_cn/design/decoupling_3.go)
      - [接口组合解耦](go_cn/design/decoupling_4.go)
    - 转换: 
      - [接口转换 | 类型断言](go_cn/design/conversion_1.go)
      - [运行时类型断言](go_cn/design/conversion_2.go)
    - 接口污染: 
      - [接口污染](go_cn/design/pollution_1.go)
      - [移除接口污染](go_cn/design/pollution_2.go)
    - 模拟: 
      - [模拟的包](go_cn/design/mocking_1.go)
      - [客户端样本](go_cn/design/mocking_2.go)
  - 错误处理: 
    - [错误默认值](go_cn/design/error_1.go)
    - [错误变量](go_cn/design/error_2.go)
    - [类型为上下文](go_cn/design/error_3.go)
    - [行为作为上下文](go_cn/design/error_4.go)
    - [查找错误接口的空值的bug/陷阱](go_cn/design/error_5.go)
    - [错误包装](go_cn/design/error_6.go)
  - 包: [指导方案](https://github.com/ardanlabs/gotraining/blob/master/topics/go_cn/design/packaging/README.md)
  - 依赖管理: [Go Modules](https://blog.golang.org/using-go-modules)
- **并发**
  - **结构**
    - Goroutine: 
      - [Go内部调度器](go_cn/concurrency/goroutine_1.go)
      - [语言结构](go_cn/concurrency/goroutine_2.go)
      - [Goroutine时间片](go_cn/concurrency/goroutine_3.go)
      - [Goroutines和并行](go_cn/concurrency/goroutine_4.go)
    - 数据竞争: 
      - [竞态检测](go_cn/concurrency/data_race_1.go)
      - [原子函数](go_cn/concurrency/data_race_2.go)
      - [锁](go_cn/concurrency/data_race_3.go)
      - [读写锁](go_cn/concurrency/data_race_4.go)
    - Channel: 
      - [指导方案](https://github.com/ardanlabs/gotraining/tree/master/topics/go#concurrent-software-design)
      - [语言结构 | 未缓冲的channel: Signaling with(out) data](go_cn/concurrency/channel_1.go)
      - [未缓冲的channel: 双重信号 | 带缓冲的channel: 关闭和遍历 | 未缓冲的channel: 选择和接收 | 未缓冲的channel: 选择和发送 | 带缓冲的channel: Select and drop](go_cn/concurrency/channel_2.go)
      - [未缓冲的channel (Tennis match)](go_cn/concurrency/channel_3.go)
      - [未缓冲的channel (Replay race)](go_cn/concurrency/channel_4.go)
      - [带缓冲的channel: Fan Out](go_cn/concurrency/channel_5.go)
      - [Select](go_cn/concurrency/channel_6.go)
  - **模式**
    - Context: 
      - [在上下文存储和检索数据](go_cn/concurrency/context_1.go)
      - [WithCancel](go_cn/concurrency/context_2.go)
      - [WithDeadline](go_cn/concurrency/context_3.go)
      - [WithTimeout](go_cn/concurrency/context_4.go)
      - [请求/响应](go_cn/concurrency/context_5.go)
    - Pattern
      - Task
      - Logger
- **测试和分析**
  - 测试: 
    - [基础单元测试](go_cn/testing/basic_test.go)
    - [表格测试](go_cn/testing/table_test.go)
    - [子测试](go_cn/testing/sub_test.go)
    - [web服务器](go_cn/testing/web_server)
    - [模拟服务器](go_cn/testing/web_test.go)
    - [测试覆盖率](go_cn/testing/README.md)
  - 基准测试
    - [基础基准测试](go_cn/benchmark/basic_test.go)
    - [子基准测试](go_cn/benchmark/sub_test.go)
  - Fuzzing
    - [指导方案](https://github.com/ardanlabs/gotraining/blob/master/topics/go_cn/testing/fuzzing/README.md)
  - 分析
    - 堆栈跟踪: [Review](go_cn/profiling/stack_trace_1.go) | [Packing](go_cn/profiling/stack_trace_2.go)
    - GODEBUG: [内存跟踪](go_cn/profiling/memory_tracing.go)

## 更多资源:

- [Ultimate Go Programming](https://www.safaribooksonline.com/library/view/ultimate-go-programming/9780134757476/)
- [ardanlabs/gotraining/topics/courses/go](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md)
- [Computer Systems: A Programmer's Perspective](https://www.amazon.com/Computer-Systems-Programmers-Perspective-3rd/dp/013409266X)
- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack)
- [Thoughts on Go performance optimization](https://github.com/dgryski/go-perfbook)


## Stargazers over time

[![Stargazers over time](https://starchart.cc/hoanhan101/ultimate-go.svg)](https://starchart.cc/hoanhan101/ultimate-go)

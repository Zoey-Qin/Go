# Go 学习资料

- https://tip.golang.org/ref/spec - go语言规范，所有的教程最终都归结于此；
- [The Go Programming Language](http://files.ictrek.local/resource/study/books/The Go Programming Language.pdf) - 比较经典的书，但是缺少一些新的特性，比如go mod和泛型；
- [Go Documents](https://go.dev/doc/) - go官方文档首页；
- https://go.p2hp.com/doc/ - go 中文手册
- [Go Modules Reference](https://go.dev/ref/mod) - go mod参考手册；重点阅读：
    - Modules, packages, and versions
    - Finding a repository for a module path
- go泛型文档：
    - [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics) - go泛型入门教程；
    - [An Introduction To Generics](https://go.dev/blog/intro-generics) - 泛型的介绍；
    - [Interface types](https://tip.golang.org/ref/spec#Interface_types) - interface定义

TIPS：

- 编译型语言，默认全静态编译（不依赖libc）；
- 使用go mod来新建项目，而不使用老的GOPATH模式；
- import语句的包名是字符串，字符串的解析规则在go mod参考手册中；
- go泛型比较新，但很重要；
- go默认执行**值拷贝**，但是没有拷贝构造函数；
- 关键概念：
    - goroutine（有栈协程）；
    - channel；
    - 类型的零值；
    - go按值拷贝，但是某些类型的值本身是引用（interface、slice、map、channel等）；
    - 基于GC的内存管理；





---

方向对了就不怕路远
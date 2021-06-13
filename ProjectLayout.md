# Project Layout

## 定义

*项目* 是产出为可执行文件的代码集合

*包* 是可以被项目引用的代码集合

*服务* 是基于 Go-Micro 框架开发的微服务

*系统* 是用于达成某种商业用途的多个服务的集合

*仓库* 是 git repository

## 实施

1. 这是一个 Monorepo 型的最终项目，而不是可以被引用的包
2. 一个服务的所有资料只在一个目录下，常用的资料有
    1. 服务的 **Interface** ，以 protobuf 描述
    2. 服务的实现（可以有多个）
    3. generate.go 生成代码和接口文档
    4. 服务的用途说明，操作手册，架构设计，注意事项
3. 如果是可以被依赖的服务
    1. 直接目录下是服务的 **Interface**
    2. 子目录，一般是 `impl`，是具体实现
4. 如果是最终服务，可以不需要子目录，所有内容平铺
5. 系统里的关键服务，直接在仓库下
6. 系统里的可选服务，在仓库的 `servers` 目录下
7. 关键服务不允许依赖可选服务
8. 服务不允许循环依赖
9. 仓库的 `main.go` 能组合必要的服务，产出便于部署和测试的可执行文件
10. 选取合适的服务实现，再进行组合，达成 Feature Toggle
11. 尽量只用 Go 工具链，go generate; go build 等等

## 参考

* [如何不 Review 每一行代码，同时保持代码不被写乱](https://mp.weixin.qq.com/s/UtBkJYpQHIvRQ_AQnzxxMw)
* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

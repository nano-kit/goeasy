# Project Layout 项目目录结构

## Terminology 定义

For clarity, this guide defines the following terms:

* *Project* : a file set which is built as an executable.
* *Package* : a file set which can be imported by projects.
* *Service* : a micro-service developed with Go-Micro framework.
* *System* : is composed by multiple services to serve a business purpose.
* *Repo* : git repository.

为了清晰起见，在实施条例中定义了以下这些专有名词，

* *项目* ：产出为可执行文件的代码集合
* *包* ：可以被项目引用的代码集合
* *服务* ：基于 Go-Micro 框架开发的微服务
* *系统* ：用于达成某种商业用途的多个服务的集合
* *仓库* ：git repository

## Guideline 实施

1. This project is a Mono-Repo, it is not a package.
2. A folder contains all the elements of a service, including
    1. service interface, described by protobuf;
    2. service implementations;
    3. generate.go for code/doc generating;
    4. doc for usage, manual, design and caveat.
3. If a service can be used by other service
    1. put service interface codes in the folder;
    2. put service implementation in sub-folder such as `impl`.
4. If a service can not be used by other service, put all the codes in the folder directly.
5. Put critical services folder in the repo root.
6. Put optional services folder in the folder `servers`.
7. Critical services should not depend on optional services.
8. Cyclic dependencies is not allowed between services.
9. Repo `main.go` composes any services to produce an all-in-one executable.
10. Feature Toggle is done by select service implementation and composition.
11. Go toolchain is preferred, such as go generate, go build etc.

---

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

## Reference 参考

* [如何不 Review 每一行代码，同时保持代码不被写乱](https://mp.weixin.qq.com/s/UtBkJYpQHIvRQ_AQnzxxMw)
* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

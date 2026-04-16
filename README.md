# Plugin System

这是一个基于 Go 原生 plugin 包实现的动态插件执行引擎。主要为了实现业务逻辑与主程序的解耦，并提供基础的容错治理能力。

## 核心能力

- **动态加载**：在运行时动态加载 .so 格式的插件，插件的修改与编译不影响主程序。
- **标准化契约**：通过 core.Plugin 接口统一输入输出标准 (map[string]interface{})。
- **Panic 隔离**：Manager 层接管插件内部发生的致命崩溃，防止宿主进程意外退出。
- **超时控制**：支持基于 context.Context 的执行时间限制，防止劣质插件阻塞流水线。

## 目录说明

- core/ : 核心接口定义
- manager/ : 插件管理器（负责加载、调度、异常恢复）
- plugins/filter/ : 一个模拟敏感词过滤、异常和耗时操作的业务插件
- main.go : 调度测试入口

## 环境依赖

- Go 1.20+
- Linux / macOS / FreeBSD
> 注意：由于 Go 官方 plugin 库的底层限制，本项目原生不支持 Windows。Windows 环境请使用 WSL 进行编译和测试。

## 快速运行

1. 克隆项目
git clone https://github.com/vertindWD/plugin-system-demo.git
cd plugin-system-demo

2. 编译业务插件为动态链接库 (.so)
go build -buildmode=plugin -o ./plugins/filter.so ./plugins/filter/main.go

3. 运行测试
go run main.go

运行后，主程序会依次测试并打印三种场景：正常数据流转、Panic 拦截恢复、以及 Context 超时强杀。

# Go

## 安装

```bash
tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz
# ~/.bash_profile 或者 /etc/profile
# go工程目录
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
source ~/.bash_profile
which go
go version
```

## 运行

```bash
go run hello.go      # 直接运行
go build hello.go    # 编译为二进制
./hello              # 执行编译产物
```

特殊规则：
1. 自动插入分号`;`
2. 函数左大括号 `{` 必须跟函数声明同行
3. main 函数必须在 main 包中，无参数，无返回值
4. 控制语句条件不需要括号
5. for 是唯一循环关键字,无while/do while
6. ++ 和 -- 不是表达式,只能作为语句
7. 不允许未使用变量
8. 不允许未使用 import
9. 大小写决定访问权限，无private/public等访问权限控制关键字
10. 类型必须显式转换
11. 变量声明方式有限制 `:=` 仅限在函数内使用
12. 多返回值是语言特性
13. 没有构造函数
14. 没有类和继承,使用struct、interface、composition
15. 接口隐式实现，无implements, extends
16. defer 延迟执行
17. panic/recover 替代异常,无try/catch
18. nil规则特殊，可为nil的类型pointer、slice、map、channel、func、interface
19. switch 不需要 break, 默认不fallthrough
20. package 是文件级组织单位
21. 文件初始化顺序 1. import包初始化 2. init() 3. main()
22. 空标识符 `_`用于忽略

| 规则             | 原因               |
|------------------|--------------------|
| 分号可省略       | 自动插入分号       |
| main必须main包   | 程序入口约定       |
| `{`不能换行      | 避免分号歧义       |
| if无括号         | 简化语法           |
| 只有for循环      | 统一设计           |
| ++不是表达式     | 减少复杂语义       |
| 变量必须使用     | 避免无效代码       |
| import必须使用   | 避免无效依赖       |
| 大小写控制权限   | 替代public/private |
| 类型不能隐式转换 | 避免错误           |
| 接口隐式实现     | 降低耦合           |
| switch无需break  | 减少冗余           |
| 没有class继承    | 组合优先           |
| defer后进先出    | 资源释放方便       |


## Go 语言总览

| 维度         | 说明                                                       |
|--------------|------------------------------------------------------------|
| **定位**     | 面向系统编程的静态强类型编译型语言                         |
| **设计者**   | Robert Griesemer, Rob Pike, Ken Thompson（Unix 和 C 的老兵） |
| **诞生**     | 2009 年公开，2012 年 1.0 发布                               |
| **核心理念** | 简洁、高效、可维护、并发原生                                  |

---

## 一、语言设计哲学

| 特点           | 说明                               |
|----------------|------------------------------------|
| **静态强类型** | 编译期类型检查，类型安全            |
| **编译型**     | 编译为单一二进制，无外部依赖        |
| **并发原生**   | goroutine + channel 是语言一级公民 |
| **简洁**       | 只有 25 个关键字，语法极简          |
| **快速编译**   | 编译速度远快于 C++/Rust            |
| **垃圾回收**   | 内置 GC，免去手动内存管理           |
| **内存安全**   | 指针存在但禁止指针运算             |
| **隐式接口**   | Duck Typing，无需显式 implements    |
| **组合优先**   | 组合优于继承（没有 class 和继承）    |

25 个关键字一览：

```
break    default      func    interface  select
case     defer        go      map        struct
chan     else         goto    package    switch
const    fallthrough  if      range      type
continue for          import  return     var
```

---

## 二、核心语法

### 变量

```go
var name string = "Go"       // 显式声明
name := "Go"                 // 短声明（类型推断）
var (
    a int
    b string = "hello"
)
```

`:=` 只能用于函数内部。

### 基本类型

```
bool, string
int / int8 / int16 / int32 / int64
uint / uint8 / uint16 / uint32 / uint64
float32 / float64
complex64 / complex128
byte      // uint8 别名
rune      // int32 别名，表示 Unicode 码点
```

### 控制流 — 无需括号，无三元表达式

```go
if x > 0 { ... }

for i := 0; i < n; i++ { ... }   // 唯一循环关键字
for condition { ... }              // 相当于 while
for { ... }                       // 无限循环

switch x {
case 1: ...
default: ...
}
```

### 函数 — 支持多返回值（Go 的标志性特性）

```go
func add(a, b int) int { return a + b }
func div(a, b int) (int, error) { ... }  // 多返回值，常用作错误处理
```

### defer — 延迟执行（栈式 LIFO）

```go
defer file.Close()   // 函数返回前执行，常用于资源清理
```

### 结构体（无 class）

```go
type User struct {
    Name string
    Age  int
}
```

### 方法（绑定到类型，非类）

```go
func (u User) Greet() string { return "Hello " + u.Name }
func (u *User) SetName(name string) { u.Name = name }  // 指针接收者
```

### 接口 — 隐式实现（Duck Typing）

```go
type Greeter interface {
    Greet() string
}
// User 无需显式声明 implements，只要实现了 Greet() 即自动满足 Greeter 接口
```

### 错误处理

```go
if err != nil { return err }   // Go 风格：显式处理错误

// 自定义错误
type MyError struct { Msg string }
func (e MyError) Error() string { return e.Msg }
```

### 切片（动态数组）

```go
s := []int{1, 2, 3}        // 字面量
s = append(s, 4)           // 追加
s2 := make([]int, 0, 10)   // make(type, len, cap)
sub := s[1:3]              // 切片操作 [1, 3)
```

### map

```go
m := make(map[string]int)
m["key"] = 42
val, ok := m["key"]        // ok 表示键是否存在
delete(m, "key")
```

### 指针（无指针运算）

```go
var x int = 42
var p *int = &x
*p = 43
```

### 泛型（Go 1.18+）

```go
func Map[T any](s []T, f func(T) T) []T {
    result := make([]T, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}
```

---

## 三、并发模型（Go 最核心的卖点）

### goroutine — 轻量级线程

```go
go func() {
    fmt.Println("Hello from goroutine")
}()
```

- 栈初始仅 ~4KB（对比线程 ~1MB），可扩展
- 由 Go 运行时调度（M:N 调度模型），非 OS 线程

### channel — 通信管道

```go
ch := make(chan int)       // 无缓冲 channel（同步）
ch := make(chan int, 10)   // 有缓冲 channel（异步）

ch <- 42          // 发送
val := <-ch       // 接收
close(ch)         // 关闭
```

### select — 多路复用

```go
select {
case v := <-ch1:
    fmt.Println(v)
case ch2 <- x:
    fmt.Println("sent", x)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no channel ready")
}
```

### 并发模式

```go
// 工作池
jobs := make(chan int, 100)
results := make(chan int, 100)
for w := 1; w <= 3; w++ {
    go worker(jobs, results)
}

// Fan-in / Fan-out
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            for v := range ch { out <- v }
            wg.Done()
        }(c)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}
```

**核心哲学**：
> "Do not communicate by sharing memory; instead, share memory by communicating."
> （不要通过共享内存来通信，而要通过通信来共享内存。）

---

## 四、标准库亮点

Go 的标准库被公认为现代语言中最完善的标准库之一：

| 包                  | 用途               | 亮点                                       |
|---------------------|--------------------|--------------------------------------------|
| `net/http`          | HTTP 客户端/服务端 | 无需框架即可建 Web 服务                    |
| `encoding/json`     | JSON 序列化        | 结构体 tag 映射                            |
| `io` / `os` / `fmt` | I/O 操作           | 统一的 Reader/Writer 接口                  |
| `context`           | 超时/取消/链路追踪 | 并发取消的标准方案                         |
| `sync`              | 同步原语           | Mutex, RWMutex, WaitGroup, Once, Pool, Map |
| `testing`           | 测试框架           | 内置 benchmark 和覆盖率                    |
| `time`              | 时间处理           | 严格的时间格式化模板                       |
| `flag`              | 命令行参数         | 内置 CLI 参数解析                          |
| `reflect`           | 反射               | 运行时类型检查和操作                       |
| `regexp`            | 正则               | RE2 语法                                   |
| `sort`              | 排序               | 自定义排序只需实现 Interface               |
| `crypto`            | 加密               | 完整的密码学套件                           |

一个完整的 Web 服务器只需要标准库：

```go
package main
import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Go!"))
    })
    http.ListenAndServe(":8080", nil)
}
```

---

## 五、应用场景

| 场景                  | 代表项目                                        | 为什么适合                               |
|-----------------------|-------------------------------------------------|------------------------------------------|
| **云原生基础设施**    | Docker, Kubernetes, Prometheus, Terraform, Etcd | 编译快、单二进制部署、并发模型适合网络服务 |
| **微服务 / API 服务** | 各类后端服务（字节跳动、美团大量使用）             | 标准库内置 HTTP，高并发，运维简单          |
| **CLI 工具**          | Hugo, Cobra, gh CLI, fzf（Go 版）                 | 跨平台单二进制，无运行时依赖              |
| **网络代理/中间件**   | Traefik, Caddy, Envoy（部分）                     | goroutine + channel 天然适合网络并发     |
| **数据库/存储**       | TiDB, CockroachDB, InfluxDB, Milvus             | 高性能、内存安全、并发控制好               |
| **DevOps 工具**       | Drone CI, Velero, Telegraf, OpenFaaS            | 编译快、易于分发部署                      |
| **边缘计算/IoT**      | 边缘网关、数据采集器                             | 资源占用低，交叉编译容易                  |

**国内大厂使用情况**：
- **字节跳动**：微服务基础设施全面 Go 化（Kitex 框架）
- **腾讯**：容器平台 TKE、CDN 调度
- **阿里巴巴**：云原生基础设施、Dubbo-go
- **七牛云**：早期 Go 布道者，存储系统全面用 Go

**不适合的场景**：
- 需要复杂继承体系的大型 GUI 应用
- 对延迟极度敏感的系统（GC 有 STW 暂停，虽然已有极大改善）
- 底层系统编程（内核、驱动 → C/Rust）
- 科学计算/数据科学（Python 生态更强）

---

## 六、Go vs 其他语言

| 对比维度 | Go                | Java         | Rust         | Python   | Node.js        |
|----------|-------------------|--------------|--------------|----------|----------------|
| 类型系统 | 静态强类型        | 静态强类型   | 静态强类型   | 动态     | 动态           |
| 编译速度 | ⚡ 极快            | 🐢 较慢      | 🐢 较慢      | N/A      | N/A            |
| 并发模型 | goroutine+channel | 线程+锁+异步 | async/await  | GIL 限制 | 事件循环       |
| 内存管理 | GC（并发）          | GC（分代）     | 所有权系统   | GC       | GC             |
| 部署方式 | 单二进制          | 需 JRE       | 单二进制     | 需解释器 | 需运行时       |
| 学习曲线 | 📗 低             | 📘 中        | 📕 高        | 📗 低    | 📗 低          |
| 包管理   | go mod（官方）      | Maven/Gradle | Cargo        | pip      | npm            |
| 泛型支持 | 1.18+ (2022)      | 有           | 有           | 无       | 无             |
| 错误处理 | 多返回值          | 异常         | Result 类型  | 异常     | 异常           |
| 空安全   | 指针可为 nil      | 可选类型     | Option       | None     | null/undefined |
| 代码风格 | gofmt 强制统一    | 规范灵活     | rustfmt 强制 | PEP8     | Prettier       |

---

## 七、工程实践

### 项目结构

```
project/
├── cmd/             # 可执行程序入口
│   └── server/
│       └── main.go
├── internal/        # 内部包（外部不可导入）
│   ├── handler/
│   └── service/
├── pkg/             # 可导出的公共库
├── api/             # 协议定义（proto, swagger）
├── configs/         # 配置文件
├── scripts/         # 构建脚本
├── test/            # 外部测试
└── go.mod           # 模块定义
```

### go mod

```bash
go mod init example.com/mypackage   # 初始化模块
go get example.com/lib@v1.2.3       # 添加依赖
go mod tidy                         # 清理依赖
go mod vendor                       # 依赖 vendor
```

### 构建

```bash
# 交叉编译（Go 的一大优势）
GOOS=linux GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build   # Apple Silicon
GOOS=windows GOARCH=amd64 go build

# 构建优化
go build -ldflags="-s -w" .         # 去除符号表减小体积
go build -race .                    # 竞态检测
```

### 测试

```go
// 文件名 *_test.go
func TestAdd(t *testing.T) {
    result := add(1, 2)
    if result != 3 {
        t.Errorf("expected 3, got %d", result)
    }
}

// 表格驱动测试（Go 标准风格）
func TestDivide(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 6, 3, 2},
        {"zero_divisor", 6, 0, 0},  // 预期兜底逻辑
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := divide(tt.a, tt.b); got != tt.want {
                t.Errorf("divide() = %v, want %v", got, tt.want)
            }
        })
    }
}

// Benchmark
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(1, 2)
    }
}
```

```bash
go test                        # 运行测试
go test -v                     # 详细输出
go test -bench=.               # 运行 benchmark
go test -cover                 # 覆盖率
go test -race                  # 竞态检测
```

---

## 八、常用框架和工具

### Web 框架

| 框架  | 特点                          |
|-------|-------------------------------|
| Gin   | 最流行，高性能，中间件丰富      |
| Echo  | 简约，性能好，内置可接入        |
| Fiber | 类 Express 风格的高性能框架   |
| Beego | 全栈框架（类似 Django）         |
| Chi   | 轻量，兼容 net/http，适合微服务 |

### 其他生态

| 类别       | 推荐库                       |
|------------|------------------------------|
| 数据库 ORM | GORM, Ent, SQLx              |
| 配置       | Viper, envconfig             |
| 日志       | Zap（高性能）, Zerolog, Logrus |
| CLI 框架   | Cobra（K8s 在用）, Urfave/cli  |
| 任务队列   | Asynq, Machinery             |
| RPC        | gRPC-Go, Kitex（字节跳动）     |
| 消息队列   | Sarama（Kafka）, NSQ           |
| 测试       | Testify, GoMock, Ginkgo      |
| 代码检查   | golangci-lint, staticcheck   |
| 热重载     | Air, Fresh, CompileDaemon    |

---

## 九、学习路线建议

```
基础语法
  ├── 变量、类型、控制流
  ├── 函数、包、模块
  └── 结构体、接口、方法
       ↓
深入核心
  ├── goroutine + channel（并发编程）
  ├── 错误处理 + defer
  └── 反射 + 泛型
       ↓
标准库实战
  ├── net/http 构建 Web 服务
  ├── encoding/json 数据序列化
  └── testing 测试覆盖
       ↓
项目实战（从易到难）
  ├── CLI 工具（文件处理/日志分析）
  ├── HTTP API 服务（CRUD + 数据库）
  ├── 并发爬虫/数据采集器
  └── 容器化微服务（Docker + K8s）
```

### 推荐资源

| 资源                                                                            | 说明                                     |
|---------------------------------------------------------------------------------|------------------------------------------|
| [Go Tour](https://go.dev/tour/)                                                 | 官方交互式入门教程                       |
| 《The Go Programming Language》                                                   | 权威书籍（Alan Donovan & Brian Kernighan） |
| [Effective Go](https://go.dev/doc/effective_go)                                 | 官方最佳实践                             |
| [Go by Example](https://gobyexample.com/)                                       | 示例驱动的学习站                         |
| [Go 语言设计与实现](https://draveness.me/golang/)                               | 中文深入源码分析                         |
| [Golang Developer Roadmap](https://github.com/Alikhll/golang-developer-roadmap) | 学习路线图                               |

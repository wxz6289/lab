# Go 中的 init 和 defer

## defer — 延迟调用

defer 用于注册一个函数调用，在**外层函数返回之前**执行。常用于资源清理、解锁、关闭文件等。

### defer 核心特点

1. **后进先出（LIFO）** — 多个 defer 按声明顺序倒序执行
2. **参数在声明时求值** — 不是执行时求值
3. **在外层函数 return 之后、返回值返回之前执行**
4. **可以修改命名返回值**

### defer 基本用法

```go
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close() // 函数返回前关闭文件
    // 业务逻辑...
    return nil
}
```

### 多个 defer 的 LIFO 顺序

```go
func example() {
    defer fmt.Println("first deferred")  // 最后执行
    defer fmt.Println("second deferred") // 其次执行
    defer fmt.Println("third deferred")  // 最先执行
    fmt.Println("hello")
}
// 输出:
// hello
// third deferred
// second deferred
// first deferred
```

### 参数即时求值

```go
func count() {
    i := 0
    defer fmt.Println(i) // 输出 0（声明时 i=0）
    i++
    defer fmt.Println(i) // 输出 1（声明时 i=1）
    i++
}
// 输出:
// 1
// 0
```

### 修改命名返回值

```go
func double(n int) (result int) {
    defer func() {
        result *= 2 // 可以访问和修改命名返回值
    }()
    return n + 1 // result = n+1, 然后 defer 修改为 (n+1)*2
}
// double(3) => 8
```

### 常见的 defer 模式

| 场景 | 典型写法 |
| --- | --- |
| 关闭文件 | `defer f.Close()` |
| 解锁互斥锁 | `defer mu.Unlock()` |
| 记录函数耗时 | `defer timeTrack(time.Now(), "funcName")` |
| 恢复 panic | `defer func() { recover() }()` |
| 关闭 channel | `defer close(ch)` |

### 常见陷阱

- **不要在 defer 中层调用 `recover()`** — `recover()` 必须在 `defer` 的函数**内部**调用，`defer recover()` 不会生效

  ```go
  // ❌ 错误
  defer recover()

  // ✅ 正确
  defer func() { recover() }()
  ```

- **defer 放在错误检查之前会导致问题** — 延迟注册应该在拿到资源后马上做，而不是有错误后才做

  ```go
  // ❌ 错误 — 如果 f 是 nil，defer f.Close() 会 panic
  f, _ := os.Open(path)
  defer f.Close()
  ```

---

## init — 包初始化函数

`init` 是 Go 中特殊的函数，在 `main` 函数执行前自动运行，用于包级别的初始化。

### init 核心特点

1. **自动执行** — 无需显式调用，程序启动时自动运行
2. **执行顺序按包导入依赖链，一个包内按文件名字母序、文件内按声明顺序**
3. **每个文件可以有多个 `init` 函数**
4. **无参数、无返回值**
5. **执行完成后才执行 `main`**

### init 基本用法

```go
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "dsn")
    if err != nil {
        log.Fatal(err)
    }
}
```

### 初始化执行顺序

```text
import → 常量(const) → 全局变量(var) → init() → main()
```

详细地说：

1. 初始化该包导入的依赖包（递归）
2. 计算包级常量的值
3. 执行包级变量声明和赋值（依赖分析驱动）
4. 执行 `init` 函数（按出现顺序）

### 注册模式

一种常见用法是在 `init` 中向全局注册表注册自己（如数据库驱动）：

```go
// mysql/driver.go
func init() {
    sql.Register("mysql", &MySQLDriver{})
}

// main.go — 只需导入，无需显式引用
import _ "github.com/go-sql-driver/mysql"
```

空导入 `_` 就是为了触发 `init` 的副作用。

### 多个 init 的执行顺序

```go
// 一个文件内多个 init
func init() { fmt.Println("init 1") }
func init() { fmt.Println("init 2") }

// 如果 a.go 和 b.go 在同一包
// a.go: func init() { fmt.Println("a") }
// b.go: func init() { fmt.Println("b") }
// 执行顺序按文件名字母序: a → b
```

### 注意事项

- **不要依赖同一个包内不同文件 init 的特定顺序** — 避免耦合
- **不要依赖不同包之间 init 的隐式顺序** — 通过导入结构控制，`go vet` 会检测初始化循环
- **init 不能显式调用** — 不能写 `init()`
- **init 中的 panic 会导致程序崩溃** — 建议做健壮的错误处理

---

## defer + init 对比表

| 维度 | defer | init |
| --- | --- | --- |
| 执行时机 | 外层函数返回前 | main 函数启动前 |
| 执行次数 | 每次函数调用时 | 程序生命周期一次 |
| 参数/返回值 | 无 | 无参数、无返回值 |
| 用途 | 资源清理、锁释放 | 包级别初始化、驱动注册 |
| 显式调用 | 不可显式调用 | 不可显式调用 |
| 数量 | 任意多个（LIFO） | 每个文件任意多个 |
| 错误处理 | 配合命名返回值 | 通常 log.Fatal |

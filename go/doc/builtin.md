# Go 内置函数完整总结

Go 语言有 15 个内置函数，定义在 `builtin` 包中，无需导入即可使用。

---

## 1. 切片操作

### append

```go
func append(slice []T, elems ...T) []T
```

向切片追加元素。容量不足时自动扩容，返回新切片；可追加另一个切片（用 `...` 解包）。

### copy

```go
func copy(dst, src []T) int
```

将 src 元素复制到 dst，复制 min(len(dst), len(src)) 个元素，返回复制的元素个数。

```go
s := []int{1, 2}
s = append(s, 3, 4)                    // [1, 2, 3, 4]
s = append(s, []int{5, 6}...)          // [1, 2, 3, 4, 5, 6]

dst := make([]int, 2)
n := copy(dst, s)                      // n = 2, dst = [1, 2]
```

---

## 2. 内存分配

### make

```go
func make(T, size ...IntegerType) T
```

创建并初始化 slice/map/channel，返回已初始化的值（非指针），适用于引用类型。

### new

```go
func new(T) *T
```

分配零值内存，返回指针，适用于所有类型。

```go
s := make([]int, 5, 10)   // len=5, cap=10 的切片
m := make(map[string]int) // 空的 map
c := make(chan int, 1)    // 带缓冲的 channel

p := new(int)             // *int，值为 0
*p = 42
```

**核心区别：** `make` 只能用于 slice/map/channel 且返回初始化后的值；`new` 可用于任何类型但只分配零值内存返回指针。

---

## 3. 长度与容量

### len

返回长度，适用于 string, slice, array, map, channel。

### cap

返回容量，适用于 slice（底层数组容量）、channel（缓冲区大小）、array。

```go
s := []int{1, 2, 3}
len(s) // 3
cap(s) // 3（如果是 make 创建可能不同）

s = append(s, 4)
len(s) // 4
cap(s) // 6（扩容了）

ch := make(chan int, 5)
len(ch) // 0（队列中元素数）
cap(ch) // 5（缓冲区容量）

arr := [3]int{1, 2, 3}
len(arr) // 3
cap(arr) // 3
```

---

## 4. 通道操作

### close

```go
func close(c chan<- T)
```

关闭 channel，通知接收方不再有数据。

```go
ch := make(chan int)
go func() {
    for i := 0; i < 3; i++ {
        ch <- i
    }
    close(ch) // 发送完毕，关闭
}()
for v := range ch {
    fmt.Println(v) // 0, 1, 2，然后退出循环
}
```

注意：只能关闭发送端；关闭已关闭的 channel 会 panic；向关闭的 channel 发送数据会 panic。

---

## 5. 复数运算

### complex

```go
func complex(r, i FloatType) ComplexType
```

构造复数。

### real

```go
func real(c ComplexType) FloatType
```

获取实部。

### imag

```go
func imag(c ComplexType) FloatType
```

获取虚部。

```go
c := complex(3, 4)     // 3+4i
r := real(c)           // 3.0
i := imag(c)           // 4.0
```

---

## 6. Map 操作

### delete

```go
func delete(m map[K]V, key K)
```

从 map 中删除键值对。

```go
m := map[string]int{"a": 1, "b": 2}
delete(m, "a")           // m 现在是 {"b": 2}
delete(m, "c")           // 安全：删除不存在的 key 是空操作
```

---

## 7. 异常处理

### panic

```go
func panic(v any)
```

触发 panic，逐层执行 defer 后退出。

### recover

```go
func recover() any
```

捕获 panic，恢复执行（仅在 defer 中有用）。

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
panic("something went wrong") // 被 defer 中的 recover 捕获，不会崩溃
```

关键区别：`panic` 类似抛异常，`recover` 类似 try-catch。但 Go 的 `recover` 只能在 defer 函数中直接调用才有效，嵌套函数中调用则无效。

---

## 8. 调试输出

### print / println

```go
print("hello")    // 输出到 stderr
println("world")  // 输出到 stderr + 换行
```

仅用于调试，不保证保留在语言中。生产代码应使用 `fmt.Print`/`fmt.Println`。

---

## 9. 快速对比表

| 函数 | 返回值 | 适用类型 | 常见误区 |
| --- | --- | --- | --- |
| make | 初始化后的值 | slice, map, channel | 不能用于 int/struct 等值类型 |
| new | 指针（零值） | 所有类型 | 不要对 slice/map 用 new，得到的是 nil |
| append | 新切片 | slice | 必须接收返回值，`s = append(s, x)` |
| copy | 复制个数 | slice | 目标切片要有足够长度 |
| len | int | string, slice, map, array, channel | - |
| cap | int | slice, channel, array | map 没有 cap，string 的 cap = len |
| close | 无 | channel | 不能在接收端关闭 |
| delete | 无 | map | 删不存在的 key 是安全的 |
| panic | 无 | - | - |
| recover | any | - | 只在 defer 直接调用有效 |
| complex/real/imag | 对应类型 | complex64, complex128 | - |

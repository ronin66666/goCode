## 使用方法

### 切片结构

可在包`src/runtime/slice.go`中看到，切片在运行时由指向底层数组的指针、长度、容量组成

```go
type slice struct {
   array unsafe.Pointer //指向对应的底层数组元素地址
   len   int
   cap   int
}
```

### 切片初始化


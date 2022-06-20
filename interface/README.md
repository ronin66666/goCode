## Go语言中的接口

在Go语言中可以为任何自定义的类型添加方法，而不仅仅是类。

接口是一种特殊的类型，是其他类型可以实现的方法签名的集合，方法签名包含：**方法名** 、**输入参数**和**返回值**，如`ReadCloser`接口

```go
type ReadCloser interface {
	Read(p []byte) (n int, err error)
	Close() error		
}
```

## Go接口使用方法

### 接口的声明与定义

- 带方法签名的接口

  ```go
  type interfaceName interface {
  	funcNameA()
  	funcNameB(a int, b string) error
  }
  ```

  

- 空接口

  指不带方法签名的接口 `interface{}`， 空接口可以储存结构体、字符串、整数等任何类型。如`fmt.println()`的参数就是一个空接口

  Go语言中提供了一种获取空接口动态类型的方法`i.(type)`，`i`代表接口变量、`type`是固定的关键字，同时此语法仅在`switch`语句中有效。

  ```go
  switch f := arg.(type) {
  	case bool:
  		p.fmtBool(f, verb)
    case float32:
    	p.fmtFloat(float64(f),32,verb)
    ...
  }
  ```

  

如果只是对接口进行声明，则当前接口变量为`nil`

```go
var i interfaceName
//i = nil	
```

### 接口实现

Go语言中接口的实现是隐式的。不用明确地指出某一个类型实现了某一个接口，只要在某一类型的方法中**实现了接口中的全部方法签名**，意味着此类型实现了这一接口

### 接口动态类型

一个接口类型的变量能够接收**任何实现了此接口**的用户**自定义类型**。

当接口中储存了具体的动态类型时，可以调用接口中所有方法

### 多接口

一个类型可以同时实现多个接口。自定义接口可以是其他接口的组合，如io库中的`ReadWriter`接口。Go语言设计指出，就鼓励开发者使用组合而不是继承的方式来编写程序。

### 接口类型断言

使用语法`i.(Type)`在运行时获取储存在接口中的类型。

在编译时会保证类型`Type`是一定实现了接口`i`的类型，否则编译不通过

### 接口的比较性

两个接口之间可以通过==或!=进行比较

```go
var a, b interface{}
fmt.println(a == b)
```

比较规则：

- 动态值为nil的接口变量总是相等的

- 如果只有1个接口为nil，那么比较结果false

- 如果两个接口部位nil且接口变量具有相同的动态类型和动态类型值，那么两个接口是相同的

- 如果接口储存的动态类型值是不可比较的，那么在运行时会报错

## 接口组成

带方法签名的接口在运行时具体结构由`iface`构成，空接口的实现方式有所不同

类型在`runtime/runtime2.go`文件中定义

```go
type iface struct {
	tab *itab
	data unsafe.Pointer
}
```

data：储存接口中动态类型的函数指针

tab: 储存了接口的类型、接口中的动态数据类型、动态数据类型的函数指针等。itab是接口的核心

```go
// layout of Itab known to compilers
// allocated in non-garbage-collected memory
// Needs to be in sync with
// ../cmd/compile/internal/reflectdata/reflect.go:/^func.WriteTabs.
type itab struct {
	inter *interfacetype  //代表接口本身
	_type *_type					//储存动态类型
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

其中_type字段代表接口储存的动态类型。Go语言的各种数据类型都是在`_type`字段的基础上通过增加额外字段来管理的，如切片

```go
type slicetype struct {
	type _type
	elem *_type
}
```

`_type`包含了类型的大小、哈希、标志、偏移量等元数据

```go
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
```




















标准库`flag`

主要功能是实现命令行参数的解析，

## Flag基本使用

```go
func main() {
	var name string
	/**
	p 指针变量 储存命令行输入的值 如 -name=hello 则 *p = hello
	name: 命令行标识 -name
	value: 默认值，
	usage: 帮助信息
	 */
	flag.StringVar(&name, "name", "flag命令行参数", "帮助信息")
	flag.Parse() //定义命令行后，进行解析
	log.Println("name : ", name)
}
```

执行命令：

```go
go run main.go -name=hello
name :  hello
```

注意：` -flag` 仅支持bool类型
	      `-flag x` 仅支持非布尔类型
	      `-flag=x `都支持

### 子命令

最常见的功能是子命令的使用，一个工具可能包含了大量相关联的功能命令，以此形成工具集

```go
func test2()  {
   var name string
   flag.Parse() //解析命令行参数，参数会储存到flag.Args中
   args := flag.Args() //获取输入命令参数
   if len(args) <= 0 {
      return
   }
   switch args[0] {
   case "go":
      goCmd := flag.NewFlagSet("go", flag.ExitOnError)
      goCmd.StringVar(&name, "name", "go 语言", "帮助信息")
      _ = goCmd.Parse(args[1:])
   case "php":
      phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
      phpCmd.StringVar(&name, "n", "php语言", "帮助信息")
      _ = phpCmd.Parse(args[1:])
   }
   log.Printf("name : %s", name)
}
```

执行命令

```bash
go run main.go go -name go语言好
name : go语言好
```

### 分析

`flag.Parse()`:总是在所有命令行参数注册的最后调用，其功能主要是解析并绑定命令行参数

内部实现

```go

//实例化一个空的命令集， 默认错误模式为ExitOnError
//错误模式分为三种：
/**
const (
	ContinueOnError ErrorHandling = iota // 返回错误继续执行
	ExitOnError                          // 退出程序
	PanicOnError                         // Call panic with a descriptive error.
)
*/
var CommandLine = NewFlagSet(os.Args[0], ExitOnError) 

func Parse() {
	// Ignore errors; CommandLine is set for ExitOnError.
	CommandLine.Parse(os.Args[1:])
}

func (f *FlagSet) Parse(arguments []string) error {
	f.parsed = true
	f.args = arguments //保存参数到args中， 因此上面示例代码中使用flag.Args()可获取到输入的命令
	for {
		seen, err := f.parseOne() //parseOne 核心处理逻辑
		if seen { //判断是否重复处理
			continue
		}
		if err == nil { 
			break
		}
		switch f.errorHandling { //判断错误类型
		case ContinueOnError:
			return err
		case ExitOnError:
			if err == ErrHelp {
				os.Exit(0)
			}
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
	return nil
}
```

`FlagSet.parseOne()`

命令行解析的核心方法，所有的命令最后都会流转到`parseOne`中进行处理，该函数主要对值类型进行判断，最后通过`flag`提供的`Value.Set`方法将参数值设置到对应的`flag`中

```go
func (f *FlagSet) parseOne() (bool, error) {
  //参数绑定规则校验 ----start----
	if len(f.args) == 0 {
		return false, nil
	}
	s := f.args[0]
	if len(s) < 2 || s[0] != '-' {
		return false, nil
	}
	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
    if len(s) == 2 { // "--" terminates the flags(中断处理--)
			f.args = f.args[1:]
			return false, nil
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, f.failf("bad flag syntax: %s", s)
	}
  // ------ end ------
	// it's a flag. does it have an argument?
	f.args = f.args[1:]
	hasValue := false
	value := ""
  //for循环获取相关值
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:] 
			hasValue = true
			name = name[0:i]
			break
		}
	}
	m := f.formal
	flag, alreadythere := m[name] // BUG
	if !alreadythere {
    //帮助信息
		if name == "help" || name == "h" { // special case for nice help message.
			f.usage()
			return false, ErrHelp
		}
		return false, f.failf("flag provided but not defined: -%s", name)
	}
	//判断value值类型
	if fv, ok := flag.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
		if hasValue {//有值：如 -v true
			if err := fv.Set(value); err != nil {
				return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
			}
		} else { //没有值，设置为true 如 -v 
			if err := fv.Set("true"); err != nil {
				return false, f.failf("invalid boolean flag %s: %v", name, err)
			}
		}
	} else {//不是bool类型
		// It must have a value, which might be the next argument.
    if !hasValue && len(f.args) > 0 {
			// value is the next arg
			hasValue = true
			value, f.args = f.args[0], f.args[1:] 
		}
		if !hasValue {
			return false, f.failf("flag needs an argument: -%s", name)
		}
		if err := flag.Value.Set(value); err != nil {
			return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
		}
	}
	if f.actual == nil {
		f.actual = make(map[string]*Flag)
	}
	f.actual[name] = flag
	return true, nil
}
```



### 定义参数类型

`flag`的命令参数类型是可以自定义的，在`Value.Set`方法中，只需要实现其对应的`Value`相关的两个接口就可以了

```go
type Value interface {
	String() string
	Set(string) error
}
```

自定义类型

```go
type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}

func (n *Name) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name("hello: " + value)
	return nil
}

func main()  {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}
```

执行命令

```bash
go run main.go  -name=Go语言
name: hello: Go语言
```


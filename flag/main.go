package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {
	//test1()
	//test2()
	test3()
}

//命令行
func test1() {
	var name string
	/**
	定义命令行参数
	p 指针变量 储存命令行输入的值 如 -name=hello 则 *p = hello
	name: 命令行标识 -name
	value: 默认值，
	usage: 帮助信息
	 */
	flag.StringVar(&name, "name", "flag命令行参数", "帮助信息")
	flag.Parse() //定义命令行后，进行解析
	log.Println("name : ", name)
	//定义完后执行
	// go run main.go -name=hello
	// name :  hello

	/**
	注意： -flag 仅支持bool类型
	      -flag x 仅支持非布尔类型
	      -flag=x 都支持
	 */
}

/**
子命令
最常见的功能是子命令的使用，一个工具可能包含了大量相关联的功能命令，以此形成工具集
 */

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
	/**
	执行命令
	go run main.go go
	name : go 语言
	执行命令
	go run main.go go -name go语言好
	name : go语言好
	*/
}

// 自定义定义参数类型
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

func test3()  {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}





package cmd

import (
	"flag/cmd_utils/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

/**
便捷单词转换工具
子命令：word
-str | -s 指定输入的字符串
-model | -m 指定转换模式
执行命令
go run main.go word -m 1 -s HelloWord
输出结果：HELLOWORD
*/

var model uint8 //模式
var str string	//输入内容

func init() {
	//命令行参数设置和初始化
	//五个参数分别为：储存输入内容的值，命令标识，命令短标识，默认值，提示信息
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Uint8VarP(&model, "model", "m", 0, "请输入单词转换的模式")
}

const (
	ModelUpper			= iota + 1		//全部转为大写
	ModelLower							//全部转为小写
	ModelUnderscoreToUpperCamelCase		//下划线转大写驼峰
	ModelUnderscoreToLowerCamelCase		//下划线转小写驼峰
	ModelCamelCaseToUnderscore			//驼峰转下划线
)

// -cmd_utils 子命令描述
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下",
	"1: 全部转为大写",
	"2: 全部转为小写",
	"3: 下划线转大写驼峰",
	"4: 下划线转小写驼峰",
	"5: 驼峰转下划线",
}, "\n")

/**
定义子命令
Use: 子命令的命令表示 go run main.go word //word为子命令
Short：简短说明，在help命令输出的帮助信息中展示
Long: 完整说明 ，在help命令输出的帮助信息中展示: go run main.go word help
 */
var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词个数转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch model  {
		case ModelUpper:
			content = word.ToUpper(str)
		case ModelLower:
			content = word.ToLower(str)
		case ModelUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModelUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModelCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help cmd_utils 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}


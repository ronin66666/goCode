package word

import (
	"strings"
	"unicode"
)
//转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)//单词首字母转为大写
	return strings.Replace(s, " ", "", -1)
}

//下划线单词转为小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	//unicode.ToLower(rune(s[0])) 将第一个字符转为小写
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰单词转为下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {//第一个字符转为小写
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {//大写转为前加 _
			output = append(output, '_')
		}
		//转为小写
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}


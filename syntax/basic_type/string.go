package main

import (
	"strings"
	"unicode/utf8"
)

func main() {
	// "可以转译\" `不可以转译
	println("aaa")
	println(`a
	b
   可以换行`)

	println("hello" + "go")
	//5
	println(len("hello"))
	//6
	println(len("你好"))
	//2 中文字符统计
	println(utf8.RuneCountInString("你好"))

	//字符函数

	strings.Split("asdwe", "s")
}

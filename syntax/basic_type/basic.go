package main

var Global = "首字母大写，全局变量"
var internal = "首字母小写，包内变量，私有变量"

var (
	First  = "a"
	second = 1
)

func main() {

	var a int = 100
	//自动推断
	var b = 200

	//自动推断，只能用于局部变量，方法内部
	c := 300
	println(a, b, c)
}

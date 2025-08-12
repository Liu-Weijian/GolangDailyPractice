package main

func main() {

	//不需要的返回值可以用_接收
	name1, _ := Func7("aa")
	println(name1)
	//使用:= 的前提，左边必须有一个新变量
}

// 无返回值
func Func1(a int) {

}
func Func2(a int, b int) {}

func Func3(a, b string) {

}

//没有重载
//func Func3(a,b string,c string)  {}

// 有返回值
func Func4(a string) string {
	return "aa"
}

// 多返回值
func Func5(a string) (string, int) {
	return "aa", 1
}

func Func6(a string) (name string, age int) {
	return "aa", 1
}

func Func7(a string) (name string, age int) {
	name = "aa"
	age = 1
	return
}

func Func8(a string) (name string, age int) {
	//不赋值 返回默认值 "",0
	return
}

//错误，返回值必须同时命名或同时不命名
//func Func8(a string) (name string,int){
//	//不赋值 返回默认值 "",0
//	return
//}

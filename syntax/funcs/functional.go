package main

func Functional4() string {
	println("hello functional4")
	return "aa"
}

func Functional5(a int) string {
	println("hello functional4")
	return "aa"
}

// 方法可以赋值给变量
func UseFunctional() {
	myFunc4 := Functional4
	myFunc4()

	myFunc5 := Functional5
	myFunc5(15)

}

func Functional6() {
	//新定义一个方法，赋值给fn
	fn := func(a int) string {
		return "aa"
	}

	fn(1)
}

func Functional7() func() string {
	//返回一个string的无参数方法
	return func() string {
		return "hello"
	}
}

package main

// 递归
// 错误使用会出现 stack overflow
func recursive(n int) {

	//必须有退出循环机制
	if n > 10 {
		return
	}
	recursive(n + 1)
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	A()
}

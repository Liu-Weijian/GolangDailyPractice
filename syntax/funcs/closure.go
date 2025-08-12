package main

// 匿名函数引用了外部作用域中的变量时就成了闭包函数
func Closure(age int) func() int {
	age = 0
	return func() int {
		age++
		return age
	}
}

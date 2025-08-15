package main

import "fmt"

// 返回的前一刻调用,栈队列形式，LOFI
func Defer() {

	defer println("defer1")
	defer println("defer2")
}

func DeferClosure() {
	i := 0

	//i=1
	defer func() {
		println(i)
	}()

	i = 1
}

func DeferClosureV2() {
	i := 0

	//i=0 值传递
	defer func(val int) {
		println(val)
	}(i)

	i = 1
}

// defer不生效
func DeferReturn1() int {
	a := 0
	defer func() {
		a = 1
	}()

	return a
}

// defer生效，同个a地址
func DeferReturn2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("--%p--%d\n", &i, i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Print(val)
			fmt.Printf("--%p--", &val)
			fmt.Println()
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Print(j)
			fmt.Printf("--%p--", &j)
			fmt.Println()
		}()
	}
}

func main() {
	//Defer()
	DeferClosureLoopV1()
	fmt.Println("-----")
	//DeferClosureLoopV2()
	//fmt.Println("-----")
	//DeferClosureLoopV3()
}

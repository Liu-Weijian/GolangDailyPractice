package main

import "fmt"

// T 类型参数，名字叫做T，约束是any,等于没有约束
type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UserList() {
	var l List[int]
	l.Append(1)

	//错误
	//l.Append(2.3)
}

func Sum[T Number](vals ...T) T {
	var sum T
	for _, val := range vals {
		sum = sum + val
	}
	return sum
}

// ~各类型的衍生类型
type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

type Integer int

func main() {
	fmt.Printf("%d \n", Sum[int](1, 3, 4, 5, 6))
	fmt.Printf("%d \n", Sum[Integer](1, 3, 4, 5, 6))
}

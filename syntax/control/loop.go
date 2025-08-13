package control

import "fmt"

func loop() {

	for i := 0; i < 10; {
		defer fmt.Println(i)
		i++
	}

}

func loop2() {
	i := 0
	for i < 10 {
		defer fmt.Println(i)
		i++
	}

	//死循环
	//for true {
	//	fmt.Println(i)
	//}

}

// ForArr 数组
func ForArr() {
	arr := [3]int{1, 2, 3}
	//range 返回值 index下表 ,val值
	for index, val := range arr {
		println("下标", index, "值", val)

	}
}

// ForSlice 切片
func ForSlice() {
	arr := []int{1, 2, 3}
	//range 返回值 index下表 ,val值
	for index, val := range arr {
		println("下标", index, "值", val)

	}
}

// ForMap 数组
func ForMap() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	//遍历顺序随机
	//遍历时value的地址是相同的
	//range  key ,val值
	for k, v := range m {
		println("k", k, "v", v)
	}
	//range  key
	for k := range m {
		println("k", k, "v", m[k])
	}
}

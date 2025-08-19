package main

import "fmt"

func Map() {
	m := map[string]int{
		"key1": 123,
	}
	m["hello"] = 345
	fmt.Println(m)

	//val为key：value的value值，ok为bool值，判断是否有当前key值
	val, ok := m["key1"]
	if ok {
		fmt.Println(val)
	}
	fmt.Print(ok)

	//容量 len
	m2 := make(map[string]int, 12)
	m2["key1"] = 123

	//delete,无返回值
	delete(m2, "key1")
}

func GoMap() {
	m1 := make(map[int]int, 16)
	for i := 0; i < 10; i++ {
		go func() {
			m1[i] = i
		}()
	}
}

func main() {
	//Map()
	//GoMap()
}

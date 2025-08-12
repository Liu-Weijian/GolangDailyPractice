package main

func main() {
	a := 456
	b := 123
	println(a + b)
	println(a - b)
	println(a * b)
	if b != 0 {
		println(a / b)
		println(a % b)
	}

	c := 12.3
	println(float64(a) + c)

	//int和 int* 也需要转换
	var d int32 = 13
	println(int32(a) + d)

}

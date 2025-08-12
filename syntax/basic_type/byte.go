package main

func main() {
	var a byte = 'a'
	var b byte = 'b'
	println(b)
	//ascii码
	println(a)
	//ascii码对应字符
	println("%c", a)

	str := "this is string"
	//值复制，不是指针
	var bs []byte = []byte(str) //类型转换
	bs[0] = 'T'

	println(str)

	//不改变
	println(string(bs))

	//aa := true
	//bb := false
	//!(a && b) equal (a != b)
	//cc := !(aa && bb)

}

package variable

import "fmt"

func main() {
	//默认类型 可以省略，会推断
	var a int = 123

	var b = 323

	var c string = "123"
	fmt.Println(a, b, c)
}

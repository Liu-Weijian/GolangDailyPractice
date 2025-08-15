package main

import "fmt"

func main() {
	u := User{}
	fmt.Printf("%v \n", u)
	fmt.Printf("%+v \n", u)
	fmt.Printf("%#v \n", u)

	//指针
	up := &User{}
	fmt.Printf("%+v \n", *up)
	//也是指针
	up2 := new(User)
	fmt.Printf("%+v \n", *up2)

	u4 := User{Name: "a1", Age: 1}
	fmt.Printf("%+v \n", u4)

	//var up3 *User

	ChangeUser()
}

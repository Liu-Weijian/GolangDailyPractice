package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(name string) {
	u.Name = name
}
func (u *User) ChangeAge(age int) {
	u.Age = age
}

func ChangeUser() {
	u1 := User{Name: "aaa", Age: 10}

	u1.ChangeName("bbb")
	u1.ChangeAge(20)
	fmt.Println(u1)
}

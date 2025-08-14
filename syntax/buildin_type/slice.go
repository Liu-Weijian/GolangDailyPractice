package main

import "fmt"

func Slice() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1: %v, len=%d cap=%d \n", s1, len(s1), cap(s1))
	s2 := make([]int, 3, 5)
	fmt.Printf("s2: %v, len=%d cap=%d \n", s2, len(s2), cap(s2))

	//{0,0,0,0} len=4 cap=4
	s3 := make([]int, 4)
	//make([]int,0,4)
	fmt.Printf("s3: %v, len=%d cap=%d \n", s3, len(s3), cap(s3))

	//len = 1 cap=5
	s4 := make([]int, 0, 5)
	s4 = append(s4, 1)
	fmt.Printf("s4: %v, len=%d cap=%d \n", s4, len(s4), cap(s4))
}

func SubSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}

	//左边是闭区间，右边是开区间(不取端点),len为[l:r]l到r的个数，cap为l到end的个数
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len=%d cap=%d \n", s2, len(s2), cap(s2))
}

func SliceExtend() {
	s1 := []int{1, 2, 3, 4}

	//未扩容
	s2 := s1[1:]
	fmt.Printf("s2: %v, len=%d cap=%d \n", s2, len(s2), cap(s2))
	s2[0] = 99
	fmt.Printf("s2: %v, len=%d cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d cap=%d \n", s1, len(s1), cap(s1))

	//扩容
	s2 = append(s2, 199)
	fmt.Printf("s2: %v, len=%d cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d cap=%d \n", s1, len(s1), cap(s1))
}

func main() {
	SliceExtend()
}

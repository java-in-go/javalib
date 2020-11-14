package main

import (
	"fmt"
)

func main() {
	//list := list.New()
	//list.PushBack(1)
	testSlice()
}

func testSlice() {
	s1 := make([]int, 0, 2)
	_ = append(s1, 1)
	fmt.Println(len(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1))
}

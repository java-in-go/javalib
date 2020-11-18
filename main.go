package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list := list.New()
	//list.PushBack(1)
	//testSlice()
	testSliceClear()
}

func testList() {
	l := list.New()
	//list.Element{
	//
	//}
	//	list2.()
	l.PushBack(1)
	l.InsertAfter(2, 1)

}

func testSliceClear() {
	s1 := make([]interface{}, 0, 10)
	_ = append(s1, 1)
	fmt.Println(len(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1))
	for i := range s1 {
		s1[i] = nil
	}
	fmt.Println(len(s1))
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

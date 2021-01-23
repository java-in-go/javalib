package util

import "fmt"

type Object = interface{}

const Null = "Null"

func main() {
	fmt.Println(test() == Null)
}

func test() interface{} {
	return Null
}

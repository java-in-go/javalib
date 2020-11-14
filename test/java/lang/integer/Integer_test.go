package integer

import (
	"fmt"
	"javalib/java/lang/integer"
	"testing"
)

func TestNew(t *testing.T) {
	i := integer.New(1)
	x := integer.FromStr("2")
	value := x.IntValue()
	fmt.Println(i.ToStr())
	fmt.Println(i)
	fmt.Println(value)
}
